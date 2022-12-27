package code_lists

import (
	"letgo_repo/data_access"
	"letgo_repo/data_access/models"
	"strconv"
)

var CodeChallengeList CodeChallengeListObj

func init() {
	CodeChallengeList = append(CodeChallengeList, enterCodeChallenge(234, isPalindrome, "[1,2,3,2,1]"))
	CodeChallengeList = append(CodeChallengeList, enterCodeChallenge(1, twoSum, "[2,7,11,13],9"))
	// enter new code here
}

func enterCodeChallenge(codeNum int, runFunc interface{}, tests ...string) (resultC CodeChallenge) {
	question, questionStatus := data_access.ProblemsMapper.GetByCodeNum(codeNum)
	codeInfo := mergeToCodeInfo(question, questionStatus)

	resultC.CodeInfo = codeInfo
	resultC.RunFunc = runFunc
	resultC.CodeInfo.Tags = []string{"????"}
	resultC.CodeInfo.Tests = append(resultC.CodeInfo.Tests, tests...)
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
