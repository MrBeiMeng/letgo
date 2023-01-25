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

type Question struct {
	gorm.Model
	Typename           string           `json:"__typename"`
	AcRate             float64          `json:"acRate"`
	Difficulty         string           `json:"difficulty"`
	FreqBar            int              `json:"freqBar"`
	PaidOnly           bool             `json:"paidOnly"`
	Status             string           `json:"status" gorm:"type:varchar(255)"`
	FrontendQuestionId string           `json:"frontendQuestionId" gorm:"type:varchar(255);index:unique"`
	IsFavor            bool             `json:"isFavor"`
	SolutionNum        int              `json:"solutionNum"`
	Title              string           `json:"title" gorm:"unique" gorm:"type:varchar(191)"`
	TitleCn            string           `json:"titleCn"`
	TitleSlug          string           `json:"titleSlug"`
	CompanyTagNum      int              `json:"companyTagNum"`
	HasVideoSolution   bool             `json:"hasVideoSolution"`
	TopCompanyTags     []TopCompanyTags `json:"topCompanyTags" gorm:"many2many:questions_top_company_tags;"`
	TopicTags          []TopicTags      `json:"topicTags" gorm:"many2many:questions_topic_tags;"`
}

type OperationRecords struct {
	gorm.Model
	Summary string
	Msg     string
	OpType  string
}

type ToDoQuestion struct {
	gorm.Model
	Theme    string
	CodeNums string
	Sort     int
	Master   bool
}

type QuestionTest struct {
	gorm.Model
	FrontendQuestionId string `gorm:"index:idx_name,unique"`
	Args               string `gorm:"index:idx_name,unique"`
	RightAnswer        string
}

type Manifest struct {
	gorm.Model
	QuestionsFrontIds string `gorm:"type:varchar(1000)"`
	Title             string `gorm:"type:varchar(255);index:unique"`
	Mark              string `gorm:"type:varchar(500)"`
	Tags              string `gorm:"type:varchar(1000);commit:'标签以逗号分割'"`
}
