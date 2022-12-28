package type_def

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"strings"
)

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
	table, err := gotable.Create("no.", "title", "level", "status", "star", "tags", "url")
	if err != nil {
		println(err.Error())
		return
	}
	for _, proj := range q {
		strArr := getStrArr([]any{proj.CodeNum, proj.Title, proj.Level, proj.Status, proj.Star, strings.Join([]string{"?"}, "·"), proj.Url})
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
		strArr := getStrArr([]any{proj.CodeNum, proj.Title, proj.Level, strings.Join([]string{"?"}, "·"), proj.Url})
		err = table.AddRow(strArr)
		if err != nil {
			println(err.Error())
		}
	}

	println(table.String())
}

type Question struct {
	RunFunc          interface{}
	Level            string
	Star             string
	Tests            []string
	Status           string
	Visible          bool
	CodeNum          int
	Title            string
	EnglishTitleSlug string
	Description      string
	Tags             []string
	Url              string
}
