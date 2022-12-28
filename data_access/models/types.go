package models

type QuestionInfo struct {
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
	Tags              string
	TranslatedContent string
	CodeSnippets      string
}

type Question struct {
	QuestionInfo
	QuestionStatus
}

type QuestionStatus struct {
	QuestionId string
	Star       string
	Status     string
	Visible    bool
}
