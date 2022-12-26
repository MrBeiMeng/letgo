package code_lists

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"strings"
)

// BaseCode
//
//	@Description: 实现一个力扣题目需要实现的接口
type BaseCode interface {
	RunDemo()
	Run(args Args)
	GetCodeNum() int
	GetTags() []string
}

// CodeInfo
// @Description: 一个力扣题目应该要包含的信息
type CodeInfo struct {
	Level            string
	Star             string
	Status           string
	Visible          bool
	CodeNum          int
	Title            string
	EnglishTitleSlug string
	Description      string
	Tags             []string
	Url              string
}

func (c CodeInfo) Print() {
	printing := `题目:%s\t难度:%s\t
- - - - - - - - - 
%s
-----------------`
	fmt.Printf(printing, c.Title, c.Level, c.Description)
}

type CodeChallenge struct {
	BaseCode
	CodeInfo
}

type CodeChallengeListObj []CodeChallenge

func (c CodeChallengeListObj) GetByCodeNum(codeNum int) interface{} {
	for _, codeChallenge := range c {
		if codeChallenge.CodeNum == codeNum {
			return codeChallenge
		}
	}

	return nil
}

func (c CodeChallengeListObj) Print() {
	table, err := gotable.Create("no.", "title", "level", "status", "star", "tags", "url")
	if err != nil {
		println(err.Error())
		return
	}
	for _, proj := range c {
		strArr := getStrArr([]any{proj.GetCodeNum(), proj.Title, proj.Level, proj.Status, proj.Star, strings.Join(proj.GetTags(), "·"), proj.Url})
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

func (c CodeChallengeListObj) EasyPrint() {
	table, err := gotable.Create("no.", "title", "level", "tags", "url")
	if err != nil {
		println(err.Error())
		return
	}
	for _, proj := range c {
		strArr := getStrArr([]any{proj.GetCodeNum(), proj.Title, proj.Level, strings.Join(proj.GetTags(), "·"), proj.Url})
		err = table.AddRow(strArr)
		if err != nil {
			println(err.Error())
		}
	}

	println(table.String())
}
