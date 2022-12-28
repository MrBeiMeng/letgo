package service

import (
	"fmt"
	"letgo_repo/code_lists"
	"letgo_repo/data_access"
	"letgo_repo/service/type_def"
	"letgo_repo/utils"
	"reflect"
	"strings"
)

var QuestionsMap map[int]type_def.Question = make(map[int]type_def.Question)
var QuestionsList []type_def.Question

func init() {

	// 注册题目
	for _, solution := range code_lists.QuestionSolutionsV1 {
		question := data_access.ProblemsMapper.GetByCodeNum(solution.CodeNum)

		var questionValue type_def.Question

		questionValue.RunFunc = solution.RunFunc
		questionValue.Tests = solution.Tests
		questionValue.CodeNum = solution.CodeNum

		questionValue.Tags = strings.Split(question.Tags, ",")
		questionValue.Title = question.Title
		questionValue.Level = question.Level
		questionValue.Url = "https://leetcode.cn/problems/" + question.TitleSlug
		questionValue.EnglishTitleSlug = question.TitleSlug
		questionValue.Status = question.Status
		questionValue.Description = question.Content
		questionValue.Star = question.Star
		questionValue.Visible = question.Visible

		QuestionsMap[questionValue.CodeNum] = questionValue
		QuestionsList = append(QuestionsList, questionValue)
	}
}

type CodeServiceImpl struct {
}

func (c CodeServiceImpl) InitTodoCode(num int) {
	data_access.ProblemsMapper.InitInsertQuestionStatus(num)
}

func (c CodeServiceImpl) Run(codeNum int, argsStr string) {
	// 获取对应题目
	codeChallenge, ok := QuestionsMap[codeNum]
	if !ok {
		fmt.Printf("查无此题[%d]", codeNum)
		return
	}

	// 获取参数列表
	if !strings.EqualFold(argsStr, "") { // 如无参数，则使用默认测试参数
		runWithArgsStr(argsStr, codeChallenge)
		return
	}

	for _, argsStr2 := range codeChallenge.Tests {
		runWithArgsStr(argsStr2, codeChallenge)
	}
}

func runWithArgsStr(argsStr string, codeChallenge type_def.Question) {
	argsStrSlice := utils.RoughSplit(argsStr)
	runWithStrSlice(codeChallenge.RunFunc, argsStrSlice)
}

func runWithStrSlice(runFunc interface{}, argsStrSlice []string) {
	fmt.Printf(" args \t|%v\n", argsStrSlice)

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
			sliceKind := t.In(i).Elem().Kind()
			switch sliceKind {
			case reflect.Int:
				nums := code_lists.ArgsHandlerV1.GetIntArr(argsStrSlice[i])
				argsSlice = append(argsSlice, reflect.ValueOf(nums))
			default:
				fmt.Printf("slice kind [%d]", sliceKind)
			}
		case reflect.Int:
			num := code_lists.ArgsHandlerV1.GetInt(argsStrSlice[i])
			argsSlice = append(argsSlice, reflect.ValueOf(num))
		case reflect.Pointer:
			linkedList := code_lists.ArgsHandlerV1.GetLinkedList(argsStrSlice[i])
			argsSlice = append(argsSlice, reflect.ValueOf(linkedList))
		default:
			fmt.Printf("other kind [%d]", t.In(i).Kind())
		}
	}

	called := v.Call(argsSlice)

	fmt.Printf("return\t|")
	for _, cd := range called {
		switch cd.Kind() {
		case reflect.Bool:
			fallthrough
		case reflect.Slice:
			fmt.Printf("%v", cd)
		case reflect.Pointer:
			linkedList := cd.Convert(reflect.TypeOf(code_lists.ListNode{}))
			linkedList.MethodByName("Print").Call([]reflect.Value{}) // todo 未测试
		default:
			fmt.Printf("return kind | %v", cd.Kind())
		}
	}
}

func (c CodeServiceImpl) Search(queryWrapper type_def.CodeQueryWrapper) (resultList type_def.Questions) {

	return resultFilter(queryWrapper, resultList)
}

func resultFilter(queryWrapper type_def.CodeQueryWrapper, resultList []type_def.Question) []type_def.Question {
	for _, item := range QuestionsList {

		if !queryWrapper.ShowHidden {
			if item.Visible == false {
				continue
			}
		}

		if queryWrapper.Star != "" && !strings.EqualFold(queryWrapper.Star, item.Star) {
			continue
		}

		if queryWrapper.Level != "" && !strings.EqualFold(queryWrapper.Level, item.Level) {
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
