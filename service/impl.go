package service

import (
	"fmt"
	"letgo_repo/code_lists"
	"letgo_repo/data_access"
	"letgo_repo/data_access/models"
	"letgo_repo/service/type_def"
	"letgo_repo/utils"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//
//var QuestionsMap map[int]type_def.Question = make(map[int]type_def.Question)
//var QuestionsList []type_def.Question
//
//func init() {
//
//	// 注册题目
//	for _, solution := range code_lists.QuestionSolutionsV1 {
//		question := data_access.ProblemsMapper.GetByCodeNum(solution.CodeNum)
//
//		var questionValue type_def.Question
//
//		questionValue.Questions = question
//		questionValue.Url = "https://leetcode.cn/problems/" + question.TitleSlug
//		for _, topTag := range questionValue.TopicTags {
//			questionValue.Tags = append(questionValue.Tags, topTag.NameTranslated)
//		}
//
//		for _, topCompanyTag := range questionValue.TopCompanyTags {
//			questionValue.TopUsedCompanies = append(questionValue.TopUsedCompanies, topCompanyTag.Slug)
//		}
//		codeNum, _ := strconv.Atoi(questionValue.FrontendQuestionId)
//		questionValue.RunFunc = solution.RunFunc
//		for _, test := range solution.Tests {
//			questionValue.Tests = append(questionValue.Tests, test)
//		}
//		questionValue.CodeNum = codeNum
//		QuestionsMap[codeNum] = questionValue
//		QuestionsList = append(QuestionsList, questionValue)
//	}
//}

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

func (c CodeServiceImpl) Run(codeNum int, argsStr string) {
	// 获取对应题目
	var solution code_lists.QuestionSolution

	foundFlag := false
	for _, tmpSolution := range code_lists.QuestionSolutionsV1 {
		if tmpSolution.CodeNum == codeNum {
			solution = tmpSolution
			foundFlag = true
			break
		}
	}
	if !foundFlag {
		fmt.Printf("查无此题[%d]", codeNum)
	}

	// 获取参数列表
	if !strings.EqualFold(argsStr, "") { // 如无参数，则使用默认测试参数
		runWithArgsStr(argsStr, solution)
		return
	}

	fmt.Printf("| 运行时间\t| %s\n", utils.GetColorGreen(time.Now().Format("2006-01-02 15:04:13")))

	for _, argsStr2 := range solution.Tests {
		runWithArgsStr(argsStr2, solution)
		fmt.Printf("\n")
	}
}

func runWithArgsStr(argsStr string, solution code_lists.QuestionSolution) {
	argsStrSlice := utils.RoughSplit(argsStr)
	runWithStrSlice(solution.RunFunc, argsStrSlice)
}

func runWithStrSlice(runFunc interface{}, argsStrSlice []string) {
	fmt.Printf("| 参数列表 \t| %v\t|\n", strings.Join(argsStrSlice, ","))

	t := reflect.TypeOf(runFunc)
	v := reflect.ValueOf(runFunc)

	if len(argsStrSlice) != t.NumIn() {
		println("参数数量错误")
		return
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
	time.Sleep(time.Millisecond)

	duration := time.Now().Sub(startedTime) - time.Millisecond

	str := utils.GetColorGreen(fmt.Sprintf("\t[%v]\t", duration))
	if duration > 300 {
		str = utils.GetColorYellow(str)
	}

	if duration > 500 {
		str = utils.GetColorRed(str)
	}

	fmt.Printf("| return\t|")
	for _, cd := range called {
		switch cd.Kind() {
		case reflect.Int:
			fallthrough
		case reflect.Bool:
			fallthrough
		case reflect.Slice:
			fmt.Printf(" %v\t", cd)
		case reflect.Pointer:
			linkedList := cd.Convert(reflect.TypeOf(&code_lists.ListNode{}))
			linkedList.MethodByName("Print").Call([]reflect.Value{}) // todo 未测试
		default:
			fmt.Printf("return kind | %v", cd.Kind())
		}
	}
	fmt.Printf("%s", str)
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
