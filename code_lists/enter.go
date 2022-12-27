package code_lists

import (
	"letgo_repo/data_access"
	"letgo_repo/data_access/models"
	"strconv"
)

var CodeChallengeList CodeChallengeListObj

func init() {
	CodeChallengeList = append(CodeChallengeList, enterCodeChallenge(PalindromeLinkedList{}))
	CodeChallengeList = append(CodeChallengeList, enterCodeChallenge(TwoSum{}))
	// enter new code here
}

func enterCodeChallenge(baseCode BaseCode) (resultC CodeChallenge) {
	question, questionStatus := data_access.ProblemsMapper.GetByCodeNum(baseCode.GetCodeNum())
	codeInfo := mergeToCodeInfo(question, questionStatus)

	resultC.CodeInfo = codeInfo
	resultC.BaseCode = baseCode
	resultC.CodeInfo.Tags = baseCode.GetTags()
	return resultC
}

func mergeToCodeInfo(question models.Question, questionStatus models.QuestionStatus) CodeInfo {
	codeInfo := CodeInfo{}
	codeInfo.Title = question.TranslatedTitle
	codeInfo.CodeNum, _ = strconv.Atoi(question.Id)
	codeInfo.Level = question.Level
	codeInfo.Description = question.TranslatedContent
	codeInfo.Visible = true
	codeInfo.Url = "https://leetcode.cn/problems/" + question.TitleSlug
	codeInfo.Star = questionStatus.Star
	codeInfo.Status = questionStatus.Status
	codeInfo.Visible = questionStatus.Visible
	codeInfo.EnglishTitleSlug = question.TitleSlug
	return codeInfo
}
