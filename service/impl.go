package service

import (
	"fmt"
	"letgo_repo/code_lists"
	"letgo_repo/data_access"
	"letgo_repo/service/type_def"
	"letgo_repo/utils"
	"reflect"
	"strconv"
	"strings"
)

type CodeServiceImpl struct {
}

func (c CodeServiceImpl) InitTodoCode(num int) {
	data_access.ProblemsMapper.InitInsertQuestionStatus(num)
}

func (c CodeServiceImpl) GetLinkedList(linkedLists string) (result []*code_lists.ListNode) {
	return code_lists.ArgsHandlerV1.GetLinkedLists(linkedLists)
}

func (c CodeServiceImpl) SearchInDBByNo(codeNum int) (result code_lists.CodeInfo) {
	question := data_access.ProblemsMapper.GetByCodeNumInDB(codeNum)

	result.Title = question.TranslatedTitle
	result.CodeNum, _ = strconv.Atoi(question.Id)
	result.Level = question.Level
	result.Description = question.TranslatedContent
	result.Visible = true
	result.Url = "https://leetcode.cn/problems/" + question.TitleSlug
	result.EnglishTitleSlug = question.TitleSlug
	return result
}

func (c CodeServiceImpl) Run(codeNum int, argsStr string) {
	// 获取对应题目
	codeChallenge, ok := code_lists.CodeChallengeList.GetByCodeNum(codeNum)
	if !ok {
		fmt.Printf("查无此题[%d]", codeNum)
		return
	}

	// 获取参数列表
	if strings.EqualFold(argsStr, "") { // 如无参数，则使用默认测试参数
		argsStr = codeChallenge.GetTests()[0]
	}
	argsStrSlice := utils.RoughSplit(argsStr)

	// 运行
	runWithStrArgs(codeChallenge.RunFunc, argsStrSlice)
}

func runWithStrArgs(runFunc interface{}, argsStrSlice []string) {
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

func (c CodeServiceImpl) Search(queryWrapper type_def.CodeQueryWrapper) (resultList code_lists.CodeChallengeListObj) {

	for _, item := range code_lists.CodeChallengeList {

		if !queryWrapper.ShowHidden {
			if item.CodeInfo.Visible == false {
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
