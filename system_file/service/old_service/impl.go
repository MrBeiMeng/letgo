package old_service

import (
	"fmt"
	"letgo_repo/system_file/code_enter"
	"letgo_repo/system_file/data_access/models"
	"letgo_repo/system_file/data_access/problems_mapper"
	"letgo_repo/system_file/service/type_def"
	utils "letgo_repo/system_file/utils"
	"letgo_repo/system_file/utils/logger"
	"reflect"
	"sort"
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
	for _, solution := range code_enter.QuestionSolutionsV1 {
		question := problems_mapper.ProblemsMapper.GetByCodeNum(solution.CodeNum)

		var questionValue type_def.Question

		questionValue.Question = question
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
	for _, modelToDoQuestion := range problems_mapper.ProblemsMapper.GetTodos() {
		modelToDoQuestion.CodeNums = strings.TrimSpace(modelToDoQuestion.CodeNums)
		codeNums := strings.Split(modelToDoQuestion.CodeNums, ",")
		allNum := len(codeNums)
		// 查询做完了的数量
		countDone := problems_mapper.ProblemsMapper.CountDone(codeNums)
		// 装配
		toDoQuestion := convModel(modelToDoQuestion, countDone, allNum)
		toDoQuestion.CodeNums = formatCodeNum(codeNums)

		resultList = append(resultList, toDoQuestion)
	}
	return
}

func formatCodeNum(codeNums []string) string {
	resultStrList := make([]string, 0)
	for _, codeNum := range codeNums {
		question := problems_mapper.ProblemsMapper.GetByCodeNumInDB(codeNum)
		color := utils.GetColorCyan("·")
		switch question.Difficulty {
		case "EASY":
			color = utils.GetColorCyan("·")
		case "MEDIUM":
			color = utils.GetColorYellow("·")
		case "HARD":
			color = utils.GetColorPurple("·")
		}

		resultStrList = append(resultStrList, fmt.Sprintf("%s%s", codeNum, color))
	}

	return strings.Join(resultStrList, ",")
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
	progressStr = fmt.Sprintf("[%d/%d]", a, b)
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

	return utils.GetColorDefault(progressStr)
}

func (c CodeServiceImpl) OperateLog(summary, msg, opType string) {
	problems_mapper.ProblemsMapper.OperationLog(summary, msg, opType)
}

func (c CodeServiceImpl) GetByCodeNums(nums []string) (result type_def.Questions) {
	for _, num := range nums {
		numInt, err := strconv.Atoi(num)
		if err != nil {
			logger.Logger.Warn(err.Error())
		}
		result = append(result, c.GetByCodeNum(numInt))
	}

	return result
}

func (c CodeServiceImpl) GetByCodeNum(num int) (result type_def.Question) {

	var question models.Question

	//data_access.MysqlDB.Where("frontend_question_id = ?", num).First(&question)
	question = problems_mapper.ProblemsMapper.GetByCodeNum(num)

	result.Question = question
	result.Url = "https://leetcode.cn/problems/" + question.TitleSlug
	codeNum, _ := strconv.Atoi(question.FrontendQuestionId)
	result.CodeNum = codeNum

	result.Question = question
	result.Url = "https://leetcode.cn/problems/" + question.TitleSlug
	for _, topTag := range result.TopicTags {
		result.Tags = append(result.Tags, topTag.NameTranslated)
	}

	for _, topCompanyTag := range result.TopCompanyTags {
		result.TopUsedCompanies = append(result.TopUsedCompanies, topCompanyTag.Slug)
	}
	return result
}

func (c CodeServiceImpl) InitTodoCode(num int) {
	problems_mapper.ProblemsMapper.InitInsertQuestionStatus(num)
}

func (c CodeServiceImpl) Run(runWrapper type_def.RunWrapper) {
	codeNum := runWrapper.CodeNum

	// 获取对应题目
	solution, ok := getSolutionByCodeNum(codeNum)
	if !ok {
		fmt.Printf("查无此题[%d]\n", codeNum)
		return
	}

	// 获取参数列表 - 新传入的 或 保存在数据库的
	serviceArgsList := getArgs(codeNum, runWrapper.ArgsStr, runWrapper.RightAnswer, solution)
	if len(serviceArgsList) == 0 {
		fmt.Printf("暂无测试参数\n")
		return
	}

	// log
	fmt.Printf("| 运行时间\t| %s\n", utils.GetColorGreen(time.Now().Format("2006-01-02 15:04:13")))
	t := reflect.TypeOf(solution.RunFunc)
	fmt.Printf("| 函数类型\t| %s\n", t.String())

	for _, tmpServiceArg := range serviceArgsList { //运行并打印日志
		argsStrSlice := utils.RoughSplit(tmpServiceArg.Args)

		calledList, duration := runWithStrSlice(solution.RunFunc, argsStrSlice)

		durationStr := durationFormat(duration)

		tmpAnswer := sprintCalled(calledList)

		if tmpServiceArg.RightAnswer != "" {
			verifyAnswer(tmpAnswer, tmpServiceArg)
		}

		if (testFromUser(tmpServiceArg) && tmpServiceArg.RightAnswer != "") || runWrapper.SaveAll {
			saveTest(runWrapper, tmpAnswer, tmpServiceArg, codeNum)
		}

		if len(tmpServiceArg.Args) >= 38 {
			tmpServiceArg.Args = tmpServiceArg.Args[:38] + "..."
		}

		fmt.Printf("| %s\t| 参数列表 %s\t| 结果 %s\t| 用时 %s", time.Now().Format("15:04:13"), tmpServiceArg.Args, tmpAnswer, durationStr)
		println()
	}

	if runWrapper.Done {
		problems_mapper.ProblemsMapper.QuestionDone(fmt.Sprintf("%d", codeNum))
		fmt.Printf("question status done\n")
	}
}

func saveTest(runWrapper type_def.RunWrapper, tmpAnswer string, tmpServiceArg type_def.QuestionTest, codeNum int) {
	rightAnswer := tmpAnswer
	if tmpServiceArg.RightAnswer != "" { // 如果用户携带了正确结果,否则使用函数运行的结果
		rightAnswer = tmpServiceArg.RightAnswer
	}

	runFunc := func() {
		err := problems_mapper.ProblemsMapper.SaveOrUpdateTest(codeNum, runWrapper.ArgsStr, rightAnswer)
		if err != nil {
			println(err.Error())
		}
		print("saved")
	}

	utils.ThreadUtil.AddThread(runFunc)
}

func testFromUser(tmpServiceArg type_def.QuestionTest) bool {
	return !tmpServiceArg.Saved
}

func goSaveTest(codeNum int, savingFlag *bool, arg, rightAnswer string) bool {

	*savingFlag = true
	go func() {
		err := problems_mapper.ProblemsMapper.SaveOrUpdateTest(codeNum, arg, rightAnswer)
		if err != nil {
			println(err.Error())
		}
		*savingFlag = false
		print("saved")
	}()
	return *savingFlag
}

func waitSaving(savingFlag *bool, printNum *int) {
	print("\t")
	for *savingFlag {
		if *printNum < 3 {
			fmt.Printf(".")
			*printNum++
		} else {
			fmt.Printf("\b")
			*printNum--
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func verifyAnswer(tmpAnswer string, tmpServiceArg type_def.QuestionTest) {
	if strings.Index(tmpServiceArg.RightAnswer, "@") != -1 {
		tmpArr := strings.Split(tmpServiceArg.RightAnswer, "@")
		command := tmpArr[0]
		rightAnswer := tmpArr[1]

		switch command {
		case "equalArr":
			// 往下进行，直接比较

			tmpAnswer := strings.TrimSpace(tmpAnswer)

			//fmt.Printf("%s_%s", tmpAnswer, rightAnswer)
			if strings.EqualFold(tmpAnswer, rightAnswer) {
				fmt.Printf(" %v ", utils.GetColorGreen("equal check ●"))
			} else {
				fmt.Printf(" %v ", utils.GetColorRed("equal check ▼"))
			}

			return

		case "sameArr":

			rightAnswerArr := strings.Split(strings.Trim(rightAnswer, "[] "), " ")
			tmpAnswerArr := strings.Split(strings.Trim(tmpAnswer, "[] "), " ")

			sort.Sort(sort.StringSlice(rightAnswerArr))
			sort.Sort(sort.StringSlice(tmpAnswerArr))

			for i, value := range rightAnswerArr {
				if tmpAnswerArr[i] != value {
					fmt.Printf(" %v ", utils.GetColorRed("same check ▼"))
					return
				}
			}

			fmt.Printf(" %v ", utils.GetColorGreen("same check ●"))
			return
		}

	}

	if strings.EqualFold(strings.TrimSpace(tmpAnswer), strings.TrimSpace(tmpServiceArg.RightAnswer)) {
		fmt.Printf(" %v ", utils.GetColorGreen("●"))
	} else {
		fmt.Printf(" %v ", utils.GetColorRed("▼"))
	}

	return
}

func getArgs(codeNum int, argsStr string, rightAnswer string, solution code_enter.QuestionSolution) []type_def.QuestionTest {
	serviceArgsList := make([]type_def.QuestionTest, 0)
	tests := problems_mapper.ProblemsMapper.GetTests(fmt.Sprintf("%d", codeNum)) // 自数据库获取保存的参数列表
	if !strings.EqualFold(argsStr, "") {                                         // 如无参数，则加入数据库中的所有测试参数
		test := type_def.QuestionTest{Args: argsStr, FrontendQuestionId: fmt.Sprintf("%d", codeNum)}
		if rightAnswer != "" {
			test.RightAnswer = rightAnswer
		}

		serviceArgsList = append(serviceArgsList, test)
	} else { // 可以在 enter.go 中直接加入测试参数，这里对其进行优先加载
		for _, testArgsStr := range solution.Tests {
			serviceArgsList = append(serviceArgsList, type_def.QuestionTest{Args: testArgsStr, FrontendQuestionId: fmt.Sprintf("%d", solution.CodeNum)})
		}
		for _, modelQuestionTest := range tests {
			serviceArgsList = append(serviceArgsList, type_def.QuestionTest{Args: modelQuestionTest.Args, FrontendQuestionId: fmt.Sprintf("%d", solution.CodeNum)})
		}
	}

	modelTestMap := make(map[string]models.QuestionTest)
	for _, testObj := range tests {
		modelTestMap[testObj.Args] = testObj
	}

	for i, serviceArg := range serviceArgsList { // 例如使用 run 123 -a[1,2,3] 时如已保存过正确答案，自动匹配答案
		if serviceArg.RightAnswer != "" {
			continue
		}

		if value, ok := modelTestMap[serviceArg.Args]; ok {
			serviceArgsList[i].Saved = true
			serviceArgsList[i].RightAnswer = value.RightAnswer
		}
	}

	return serviceArgsList
}

func sprintCalled(calledList []reflect.Value) string {
	for _, cd := range calledList {
		switch cd.Kind() {
		case reflect.Int:
			fallthrough
		case reflect.Bool:
			return fmt.Sprintf("%v", cd)
		case reflect.Array: // 数组和 slice 一视同仁
			fallthrough
		case reflect.Slice:

			strAnswers := make([]string, 0)

			for i := 0; i < cd.Len(); i++ {
				strAnswers = append(strAnswers, fmt.Sprintf("%v", cd.Index(i)))
			}

			return "[" + strings.Join(strAnswers, ",") + "]"
		case reflect.Pointer:
			linkedList := cd.Convert(reflect.TypeOf(&code_enter.ListNode{}))
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
			linkedList := cd.Convert(reflect.TypeOf(&code_enter.ListNode{}))
			linkedList.MethodByName("Print").Call([]reflect.Value{})
		default:
			fmt.Printf("return kind | %v", cd.Kind())
		}
	}
}

func durationFormat(duration time.Duration) string {
	str := fmt.Sprintf("[%v]", duration)
	if duration > 300*time.Millisecond {
		return utils.GetColorYellow(str)
	}

	if duration > 500*time.Millisecond {
		return utils.GetColorRed(str)
	}

	return utils.GetColorGreen(str)
}

func getSolutionByCodeNum(codeNum int) (solution code_enter.QuestionSolution, ok bool) {
	for _, tmpSolution := range code_enter.QuestionSolutionsV1 {
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
			num := code_enter.ArgsHandlerV1.GetInt(argsStrSlice[i])
			argsSlice = append(argsSlice, reflect.ValueOf(num))
		case reflect.Pointer:
			linkedList := code_enter.ArgsHandlerV1.GetLinkedList(argsStrSlice[i])
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
		nums := code_enter.ArgsHandlerV1.GetIntArr(argsStrSlice[i])
		argsSlice = append(argsSlice, reflect.ValueOf(nums))
	case reflect.String:
		strList := code_enter.ArgsHandlerV1.GetStringArr(argsStrSlice[i])
		argsSlice = append(argsSlice, reflect.ValueOf(strList))
	case reflect.Slice:
		sliceKind2 := t.In(i).Elem().Elem().Kind()
		switch sliceKind2 {
		case reflect.Int:
			nums := code_enter.ArgsHandlerV1.GetIntMatrix(argsStrSlice[i])
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
