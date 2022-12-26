package models

type Question struct {
	Id                string
	Title             string
	TitleSlug         string
	ArticleLive       string
	ArticleSlug       string
	Level             string
	TotalSubmitted    string
	TotalAcs          string
	FrontendId        string
	TranslatedTitle   string
	Content           string
	TranslatedContent string
	CodeSnippets      string
}

//func (q Question) ConvQuestionToCodeInfo() (result code_lists.CodeInfo) {
//	result.Title = q.TranslatedTitle
//	result.CodeNum, _ = strconv.Atoi(q.Id)
//	result.Level = q.Level
//	result.Description = q.TranslatedContent
//	result.Visible = true
//	result.Url = "https://leetcode.cn/problems/" + q.TitleSlug
//	return result
//}

type QuestionStatus struct {
	Id         int
	QuestionId string
	Star       string
	Status     string
	Visible    bool
}
