package models

import "gorm.io/gorm"

type TopicTags struct {
	gorm.Model
	Name           string `json:"name" gorm:"unique"`
	Slug           string `json:"slug"`
	NameTranslated string `json:"nameTranslated"`
	Typename       string `json:"__typename"`
}

type TopCompanyTags struct {
	gorm.Model
	ImgUrl   string `json:"imgUrl"`
	Slug     string `json:"slug" gorm:"unique"`
	Typename string `json:"__typename"`
}

type Questions struct {
	gorm.Model
	Typename           string           `json:"__typename"`
	AcRate             float64          `json:"acRate"`
	Difficulty         string           `json:"difficulty"`
	FreqBar            int              `json:"freqBar"`
	PaidOnly           bool             `json:"paidOnly"`
	Status             string           `json:"status"`
	FrontendQuestionId string           `json:"frontendQuestionId"`
	IsFavor            bool             `json:"isFavor"`
	SolutionNum        int              `json:"solutionNum"`
	Title              string           `json:"title" gorm:"unique"`
	TitleCn            string           `json:"titleCn"`
	TitleSlug          string           `json:"titleSlug"`
	CompanyTagNum      int              `json:"companyTagNum"`
	HasVideoSolution   bool             `json:"hasVideoSolution"`
	TopCompanyTags     []TopCompanyTags `json:"topCompanyTags" gorm:"many2many:questions_top_company_tags;"`
	TopicTags          []TopicTags      `json:"topicTags" gorm:"many2many:questions_topic_tags;"`
}

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
