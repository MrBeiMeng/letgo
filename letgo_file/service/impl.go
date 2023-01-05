package service

import (
	"fmt"
	"letgo_repo/code_lists"
	data_access "letgo_repo/letgo_file/data_access"
	"letgo_repo/letgo_file/data_access/models"
	"letgo_repo/letgo_file/service/type_def"
	utils "letgo_repo/letgo_file/utils"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// getQuestions
//
//	@Description: 获取保存的 questions 列表
func getQuestions() (questionsMap map[int]type_def.Question, questionsList []type_def.Question) {

	questionsMap = make(map[int]type_def.Question)

	// 注册题目
	for _, solution := range code_lists.QuestionSolutionsV1 {
		question := data_access.ProblemsMapper.GetByCodeNum(solution.CodeNum)

		var questionValue type_def.Question

		questionValue.Questions = question
		questionValue.Url = "https://leetcode.cn/problems/" + question.TitleSlug
		for _, topTag := range questionValue.TopicTags {
			questionValue.Tags = append(questionValue.Tags, topTag.NameTranslated)
		}

		for _, topCompanyTag := range questionValue.TopCompanyTags {
			questionValue.TopUsedCompanies = append(questionValue.TopUsedCompanies, topCompanyTag.Slug)
		}
		codeNum, _ := strconv.Atoi(questionValue.FrontendQuestionId)
		questionValue.RunFunc = solution.RunFunc
		for _, test := range solution.Tests {
			questionValue.Tests = append(questionValue.Tests, test)
		}
		questionValue.CodeNum = codeNum

		questionsMap[codeNum] = questionValue
		questionsList = append(questionsList, questionValue)
	}

	return
}

type CodeServiceImpl struct {
}

func (c CodeServiceImpl) GetToDos() (resultList []type_def.ToDoQuestion) {
	for _, modelToDoQuestion := range data_access.ProblemsMapper.GetTodos() {
		modelToDoQuestion.CodeNums = strings.TrimSpace(modelToDoQuestion.CodeNums)
		codeNums := strings.Split(modelToDoQuestion.CodeNums, ",")
		allNum := len(codeNums)
		// 查询做完了的数量
		countDone := data_access.ProblemsMapper.CountDone(codeNums)
		// 装配
		toDoQuestion := convModel(modelToDoQuestion, countDone, allNum)

		resultList = append(resultList, toDoQuestion)
	}
	return
}

func convModel(modelToDoQuestion models.ToDoQuestion, countDone int, allNum int) type_def.ToDoQuestion {
	toDoQuestion := type_def.ToDoQuestion{}
	toDoQuestion.Theme = modelToDoQuestion.Theme
	toDoQuestion.CodeNums = modelToDoQuestion.CodeNums
	toDoQuestion.Progress = getProgressStr(countDone, allNum)
	toDoQuestion.Master = modelToDoQuestion.Master
	return toDoQuestion
}

func getProgressStr(a, b int) string {
	fa := float64(a)
	fb := float64(b)

	progress := fa / fb
	progressStr := formatProgress(a, b, progress)

	return progressStr
}

func formatProgress(a int, b int, progress float64) (progressStr string) {
	progressStr = fmt.Sprintf("[%d/%d]\t", a, b)
	if progress >= 1 {
		progressStr = utils.GetColorGreen(progressStr)
		return
	}

	if progress >= 0.7 {
		progressStr = utils.GetColorYellow(progressStr)
		return
	}

	if progress >= 0.5 {
		progressStr = utils.GetColorBlue(progressStr)
		return
	}

	return progressStr
}

func (c CodeServiceImpl) OperateLog(summary, msg, opType string) {
	data_access.ProblemsMapper.OperationLog(summary, msg, opType)
}

func (c CodeServiceImpl) GetByCodeNum(num int) (result type_def.Question) {

	var question models.Questions

	data_access.MysqlDB.Where("frontend_question_id = ?", num).First(&question)

	result.Questions = question
	result.Url = "https://leetcode.cn/problems/" + question.TitleSlug
	codeNum, _ := strconv.Atoi(question.FrontendQuestionId)
	result.CodeNum = codeNum
	return result
}

func (c CodeServiceImpl) InitTodoCode(num int) {
	data_access.ProblemsMapper.InitInsertQuestionStatus(num)
}

func (c CodeServiceImpl) Run(codeNum int, argsStr string, rightAnswer ...string) {
	// 获取对应题目
	solution, ok := getSolutionByCodeNum(codeNum)
	if !ok {
		fmt.Printf("查无此题[%d]", codeNum)
	}

	// 获取参数列表
	argsStrList := make([]string, 0)
	if !strings.EqualFold(argsStr, "") { // 如无参数，则使用默认测试参数
		argsStrList = append(argsStrList, argsStr)
	} else {
		for _, testArgsStr := range solution.Tests {
			argsStrList = append(argsStrList, testArgsStr)
		}
	}

	// log
	fmt.Printf("| 运行时间\t| %s\n", utils.GetColorGreen(time.Now().Format("2006-01-02 15:04:13")))
	t := reflect.TypeOf(solution.RunFunc)
	fmt.Printf("| 函数类型\t| %s\n", t.String())

	for _, tmpStrArgs := range argsStrList { //运行并打印日志
		argsStrSlice := utils.RoughSplit(tmpStrArgs)

		calledList, duration := runWithStrSlice(solution.RunFunc, argsStrSlice)

		durationStr := durationFormat(duration)

		tmpAnswer := sprintCalled(calledList)

		savingFlag := false

		if len(rightAnswer) > 0 {
			if strings.EqualFold(strings.TrimSpace(tmpAnswer), strings.TrimSpace(rightAnswer[0])) {
				fmt.Printf(" %v\b\b\b\b\b", utils.GetColorGreen("●"))
			} else {
				fmt.Printf(" %v\b\b\b\b\b", utils.GetColorRed("▼"))
			}

			go func() {
				err := data_access.ProblemsMapper.SaveAnswer(codeNum, tmpStrArgs, rightAnswer[0])
				if err != nil {
					println(err.Error())
				}
				savingFlag = true
			}()
		}

		fmt.Printf("| %s\t| 参数列表 %s\t| 结果 %s\t| 用时%s\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b", time.Now().Format("15:04:13"), tmpStrArgs, tmpAnswer, durationStr)

		printNum := 0
		for !savingFlag {
			if printNum < 3 {
				fmt.Printf(".")
				printNum++
			} else {
				fmt.Printf("\b")
				printNum--
			}
			time.Sleep(time.Millisecond * 100)
		}
		println("saved")
	}

}

func sprintCalled(calledList []reflect.Value) string {
	for _, cd := range calledList {
		switch cd.Kind() {
		case reflect.Int:
			fallthrough
		case reflect.Bool:
			fallthrough
		case reflect.Slice:
			return fmt.Sprintf(" %v", cd)
		case reflect.Pointer:
			linkedList := cd.Convert(reflect.TypeOf(&code_lists.ListNode{}))
			call := linkedList.MethodByName("Sprint").Call([]reflect.Value{})
			if len(call) < 1 {
				return "函数异常"
			}
			return fmt.Sprintf("%v", call[0])
		default:
			return fmt.Sprintf("return kind | %v", cd.Kind())
		}
	}

	return ""
}

func printCalled(calledList []reflect.Value) {
	for _, cd := range calledList {
		switch cd.Kind() {
		case reflect.Int:
			fallthrough
		case reflect.Bool:
			fallthrough
		case reflect.Slice:
			fmt.Printf(" %v\t", cd)
		case reflect.Pointer:
			linkedList := cd.Convert(reflect.TypeOf(&code_lists.ListNode{}))
			linkedList.MethodByName("Print").Call([]reflect.Value{})
		default:
			fmt.Printf("return kind | %v", cd.Kind())
		}
	}
}

func durationFormat(duration time.Duration) string {
	str := utils.GetColorGreen(fmt.Sprintf("\t[%v]\t", duration))
	if duration > 300 {
		str = utils.GetColorYellow(str)
	}

	if duration > 500 {
		str = utils.GetColorRed(str)
	}
	return str
}

func getSolutionByCodeNum(codeNum int) (solution code_lists.QuestionSolution, ok bool) {
	for _, tmpSolution := range code_lists.QuestionSolutionsV1 {
		if tmpSolution.CodeNum == codeNum {
			solution = tmpSolution
			ok = true
			break
		}
	}
	return
}

// runWithStrSlice
//
//	@Description: 自动解析 argStrSlice 数组,带入至 runFunc 并导出结果和 运行用时
//	@param runFunc
//	@param argsStrSlice
//	@return []reflect.Value 运行结果列表
//	@return time.Duration 运行用时
func runWithStrSlice(runFunc interface{}, argsStrSlice []string) ([]reflect.Value, time.Duration) {

	t := reflect.TypeOf(runFunc)
	v := reflect.ValueOf(runFunc)

	if len(argsStrSlice) != t.NumIn() {
		println("参数数量错误")
		return nil, time.Duration(1)
	}

	var argsSlice []reflect.Value
	for i := 0; i < t.NumIn(); i++ {
		switch t.In(i).Kind() {
		case reflect.Slice:
			argsSlice = sliceHandler(t, i, argsStrSlice, argsSlice)
		case reflect.Int:
			num := code_lists.ArgsHandlerV1.GetInt(argsStrSlice[i])
			argsSlice = append(argsSlice, reflect.ValueOf(num))
		case reflect.Pointer:
			linkedList := code_lists.ArgsHandlerV1.GetLinkedList(argsStrSlice[i])
			argsSlice = append(argsSlice, reflect.ValueOf(linkedList))
		case reflect.String:
			argsSlice = append(argsSlice, reflect.ValueOf(argsStrSlice[i]))
		default:
			fmt.Printf(utils.GetColorRed("other kind [%d]\n"), t.In(i).Kind())
			panic("stopped")
		}
	}

	startedTime := time.Now()
	called := v.Call(argsSlice)
	time.Sleep(time.Microsecond)
	return called, time.Now().Sub(startedTime)
}

func sliceHandler(t reflect.Type, i int, argsStrSlice []string, argsSlice []reflect.Value) []reflect.Value {
	sliceKind := t.In(i).Elem().Kind()
	switch sliceKind {
	case reflect.Int:
		nums := code_lists.ArgsHandlerV1.GetIntArr(argsStrSlice[i])
		argsSlice = append(argsSlice, reflect.ValueOf(nums))
	case reflect.Slice:
		sliceKind2 := t.In(i).Elem().Elem().Kind()
		switch sliceKind2 {
		case reflect.Int:
			nums := code_lists.ArgsHandlerV1.GetIntMatrix(argsStrSlice[i])
			argsSlice = append(argsSlice, reflect.ValueOf(nums))
		}
	default:
		fmt.Printf("slice kind [%d]", sliceKind)
	}
	return argsSlice
}

func (c CodeServiceImpl) Search(queryWrapper type_def.CodeQueryWrapper) (resultList type_def.Questions) {

	return resultFilter(queryWrapper, resultList)
}

func resultFilter(queryWrapper type_def.CodeQueryWrapper, resultList []type_def.Question) []type_def.Question {
	_, questionsList := getQuestions()
	for _, item := range questionsList {

		if !queryWrapper.ShowHidden {
			if item.Visible == false {
				continue
			}
		}

		if queryWrapper.Star != "" && !strings.EqualFold(queryWrapper.Star, item.Star) {
			continue
		}

		if queryWrapper.Level != "" && !strings.EqualFold(queryWrapper.Level, item.Difficulty) {
			continue
		}

		if queryWrapper.Status != "" && !strings.EqualFold(queryWrapper.Status, item.Status) {
			continue
		}

		if queryWrapper.CodeTitle != "" && strings.Index(item.Title, queryWrapper.CodeTitle) == -1 {
			continue
		}

		if queryWrapper.CodeNum != 0 && item.CodeNum != queryWrapper.CodeNum {
			continue
		}

		if len(queryWrapper.Tags) != 0 {
			NotInFlag := true
			qTagsStr := strings.Join(queryWrapper.Tags, "-")
			// 如果 item 的tags中有任何再queryWrapper 的tags中
			for _, tag := range item.Tags {
				if strings.Index(qTagsStr, tag) != -1 {
					NotInFlag = false
				}
			}

			if NotInFlag {
				continue
			}
		}

		resultList = append(resultList, item)
	}

	return resultList
}
