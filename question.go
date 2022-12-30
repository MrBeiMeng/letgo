package main

type Question struct {
	Typename           string  `json:"__typename"`
	AcRate             float64 `json:"acRate"`
	Difficulty         string  `json:"difficulty"`
	FreqBar            int     `json:"freqBar"`
	PaidOnly           bool    `json:"paidOnly"`
	Status             string  `json:"status"`
	FrontendQuestionId string  `json:"frontendQuestionId"`
	IsFavor            bool    `json:"isFavor"`
	SolutionNum        int     `json:"solutionNum"`
	Title              string  `json:"title"`
	TitleCn            string  `json:"titleCn"`
	TitleSlug          string  `json:"titleSlug"`
	TopicTags          []struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		Slug           string `json:"slug"`
		NameTranslated string `json:"nameTranslated"`
		Typename       string `json:"__typename"`
	} `json:"topicTags"`
	Extra struct {
		CompanyTagNum    int  `json:"companyTagNum"`
		HasVideoSolution bool `json:"hasVideoSolution"`
		TopCompanyTags   []struct {
			ImgUrl   string `json:"imgUrl"`
			Slug     string `json:"slug"`
			Typename string `json:"__typename"`
		} `json:"topCompanyTags"`
		Typename string `json:"__typename"`
	} `json:"extra"`
}
