package code_lists

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"strings"
)

// BaseCode
//
//	@Description: 实现一个力扣题目需要实现的接口
type BaseCode struct {
	RunFunc interface{}
}

// CodeInfo
// @Description: 一个力扣题目应该要包含的信息
type CodeInfo struct {
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

func (c CodeInfo) Print() {
	printing := `题目:%s\t难度:%s\t
- - - - - - - - - 
%s
-----------------`
	fmt.Printf(printing, c.Title, c.Level, c.Description)
}

func (c CodeInfo) GetTests() []string {
	return c.Tests
}

type CodeChallenge struct {
	BaseCode
	CodeInfo
}

type CodeChallengeListObj []CodeChallenge

// GetByCodeNum
//
//	@Description: 通过题号找题
//	@receiver c 题目列表
//	@param codeNum 题号
//	@return CodeChallenge 对应题目对象
//	@return bool 是否找到
func (c CodeChallengeListObj) GetByCodeNum(codeNum int) (CodeChallenge, bool) {
	for _, codeChallenge := range c {
		if codeChallenge.CodeNum == codeNum {
			return codeChallenge, true
		}
	}

	return CodeChallenge{}, false
}

func (c CodeChallengeListObj) Print() {
	table, err := gotable.Create("no.", "title", "level", "status", "star", "tags", "url")
	if err != nil {
		println(err.Error())
		return
	}
	for _, proj := range c {
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

func (c CodeChallengeListObj) EasyPrint() {
	table, err := gotable.Create("no.", "title", "level", "tags", "url")
	if err != nil {
		println(err.Error())
		return
	}
	for _, proj := range c {
		strArr := getStrArr([]any{proj.CodeNum, proj.Title, proj.Level, strings.Join([]string{"?"}, "·"), proj.Url})
		err = table.AddRow(strArr)
		if err != nil {
			println(err.Error())
		}
	}

	println(table.String())
}
