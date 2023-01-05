package generate

type QuestionList struct {
	Data struct {
		ProblemSetQuestionList struct {
			Typename  string      `json:"__typename"`
			Questions []Questions `json:"questions"`
			HasMore   bool        `json:"hasMore"`
			Total     int         `json:"total"`
		} `json:"problemsetQuestionList"`
	} `json:"data"`
}

type TopicTags struct {
	Name           string `json:"name" gorm:"unique"`
	Slug           string `json:"slug"`
	NameTranslated string `json:"nameTranslated"`
	Typename       string `json:"__typename"`
}

type TopCompanyTags struct {
	ImgUrl   string `json:"imgUrl"`
	Slug     string `json:"slug" gorm:"unique"`
	Typename string `json:"__typename"`
}

type Extras struct {
	QuestionId       uint             `json:"-"`
	CompanyTagNum    int              `json:"companyTagNum"`
	HasVideoSolution bool             `json:"hasVideoSolution"`
	TopCompanyTags   []TopCompanyTags `json:"topCompanyTags" gorm:"many2many:extras_top_company_tags;"`
	Typename         string           `json:"__typename"`
}

type Questions struct {
	Typename           string      `json:"__typename"`
	AcRate             float64     `json:"acRate"`
	Difficulty         string      `json:"difficulty"`
	FreqBar            int         `json:"freqBar"`
	PaidOnly           bool        `json:"paidOnly"`
	Status             string      `json:"status"`
	FrontendQuestionId string      `json:"frontendQuestionId"`
	IsFavor            bool        `json:"isFavor"`
	SolutionNum        int         `json:"solutionNum"`
	Title              string      `json:"title" gorm:"unique"`
	TitleCn            string      `json:"titleCn"`
	TitleSlug          string      `json:"titleSlug"`
	TopicTags          []TopicTags `json:"topicTags" gorm:"many2many:question_topic_tags;"`
	Extra              Extras      `json:"extra" gorm:"foreignKey:QuestionId"`
}
