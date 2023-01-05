package type_def

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"letgo_repo/letgo_file/data_access/models"
	"strings"
)

// AnyFunc 一个方法
type AnyFunc interface{}

type ToDoQuestion struct {
	Theme    string
	CodeNums string
	Progress string
	Master   bool
}

type CodeQueryWrapper struct {
	Level      string
	Star       string
	Status     string
	ShowHidden bool
	CodeNum    int
	CodeTitle  string
	Tags       []string
}

type Questions []Question

func (q Questions) Print() {
	table, err := gotable.Create("no.", "title", "level", "status", "tags", "CTN", "url")
	if err != nil {
		println(err.Error())
		return
	}
	for _, proj := range q {
		strArr := getStrArr([]any{proj.CodeNum, proj.TitleCn, proj.Difficulty, proj.Status, strings.Join(proj.Tags, "·"), proj.CompanyTagNum, proj.Url})
		err = table.AddRow(strArr)
		if err != nil {
			println(err.Error())
		}
	}

	println(table.String())
}

func getStrArr(list []interface{}) (result []string) {
	for _, item := range list {
		result = append(result, fmt.Sprintf("%v", item))
	}

	return result
}

func (q Questions) EasyPrint() {
	table, err := gotable.Create("no.", "title", "level", "tags", "url")
	if err != nil {
		println(err.Error())
		return
	}
	for _, proj := range q {
		strArr := getStrArr([]any{proj.FrontendQuestionId, proj.Title, proj.Difficulty, strings.Join([]string{"?"}, "·"), proj.TopicTags})
		err = table.AddRow(strArr)
		if err != nil {
			println(err.Error())
		}
	}

	println(table.String())
}

type Question struct {
	RunFunc interface{}
	models.Questions
	Url              string
	Tags             []string
	TopUsedCompanies []string
	Tests            []string
	Visible          bool
	Star             string
	CodeNum          int
}

type CodeTemplateResp struct {
	Data struct {
		Question struct {
			QuestionId         string         `json:"questionId"`
			QuestionFrontendId string         `json:"questionFrontendId"`
			CodeSnippets       []CodeTemplate `json:"codeSnippets"`
			EnvInfo            string         `json:"envInfo"`
			EnableRunCode      bool           `json:"enableRunCode"`
		} `json:"question"`
	} `json:"data"`
}

type CodeTemplate struct {
	Lang     string `json:"lang"`
	LangSlug string `json:"langSlug"`
	Code     string `json:"code"`
}
