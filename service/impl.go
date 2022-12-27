package service

import (
	"fmt"
	"letgo_repo/code_lists"
	"letgo_repo/data_access"
	"letgo_repo/service/type_def"
	"strconv"
	"strings"
)

type CodeServiceImpl struct {
}

func (c CodeServiceImpl) InitTodoCode(num int) {
	data_access.ProblemsMapper.InitInsertQuestionStatus(num)
}

func (c CodeServiceImpl) GetLinkedList(linkedLists string) (result []*code_lists.ListNode) {
	return code_lists.ArgsHandlerV1.GetLinkedList(linkedLists)
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

func (c CodeServiceImpl) Run(codeNum int, args2 type_def.Args) {
	args := code_lists.Args{}
	args.ListNodes = code_lists.ArgsHandlerV1.GetLinkedList(args2.LinkedLists)

	codeChallengeI := code_lists.CodeChallengeList.GetByCodeNum(codeNum)
	var codeChallenge code_lists.CodeChallenge
	if codeChallengeI == nil {
		fmt.Printf("查无此题[%d]", codeNum)
		return
	}
	codeChallenge = codeChallengeI.(code_lists.CodeChallenge)

	codeChallenge.Run(args)
}

func (c CodeServiceImpl) RunDemo(codeNum int) {
	codeChallengeI := code_lists.CodeChallengeList.GetByCodeNum(codeNum)
	var codeChallenge code_lists.CodeChallenge
	if codeChallengeI == nil {
		fmt.Printf("查无此题[%d]", codeNum)
		return
	}
	codeChallenge = codeChallengeI.(code_lists.CodeChallenge)

	codeChallenge.RunDemo()
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
