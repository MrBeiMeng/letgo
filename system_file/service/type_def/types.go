package type_def

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/liushuochen/gotable"
	"io/ioutil"
	"letgo_repo/system_file/data_access/models"
	"letgo_repo/system_file/utils"
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
	models.Question
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

type QuestionTest struct {
	FrontendQuestionId string
	Args               string
	RightAnswer        string
	Saved              bool
}

type RunWrapper struct {
	CodeNum     int
	ArgsStr     string
	SaveAll     bool // 强制保存
	RightAnswer string
	Done        bool
}

// VersionBody 用于解析version.json
type VersionBody struct {
	ProjectName string `json:"project_name"`
	Versions    []struct {
		VersionNo string `json:"version_no"`
		Type      string `json:"type"`
		Date      string `json:"date"`
		Log       string `json:"log"`
	} `json:"version"`
}

func (v *VersionBody) InitByJsonFile(filePath string) error {
	byteStr, _ := ioutil.ReadFile(filePath)
	//fmt.Println(string(byteStr))

	err := json.Unmarshal(byteStr, &v)
	if err != nil {
		println(err.Error())
		return errors.New("version 文件发生了一些错误")
	}
	if len(v.Versions) == 0 {
		println(err.Error())
		msg := utils.GetColorRed("version 历史信息为空!")
		return errors.New(msg)
	}

	return nil
}
