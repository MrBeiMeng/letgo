// Package cmd /*
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"letgo_repo/service"
	"letgo_repo/service/type_def"
	"letgo_repo/utils"
	"letgo_repo/utils/enum"
	"strconv"
	"strings"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "添加新的题目",
	Long:  `添加新的题目到列表`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			println("请传入题号")
			return
		}

		codeNum, _ := strconv.Atoi(args[0])

		question := service.CodeService.GetByCodeNum(codeNum)
		if question.CodeNum == 0 {
			println("查无此题")
			return
		}

		golangCodeTemplate := getCodeTemplate(question.TitleSlug)

		if !sure {
			input, err := utils.GetInput(fmt.Sprintf("确定您想注册{ %d %s %s %s }吗? [y/n]", codeNum, question.TitleCn, question.Difficulty, question.Url), 0)
			if err != nil {
				panic(err)
			}

			if !strings.EqualFold(input, "y") && !strings.EqualFold(input, "yes") {
				println(utils.GetColorRed("用户取消操作"))
				return
			}
		}

		utils.InitFile(question.TitleSlug, question.Url, question.TitleCn, question.CodeNum, golangCodeTemplate.Code)

		fmt.Printf("%s| Inited %d %s %s", utils.GetColorGreen("DONE"), question.CodeNum, question.TitleCn, question.Url)
		service.CodeService.OperateLog(fmt.Sprintf("添加新题:%v", codeNum), golangCodeTemplate.Code, enum.INIT_CODE)
	},
}

func getCodeTemplate(titleSlug string) (golangCodeTemplate type_def.CodeTemplate) {
	reqBody := "{\"query\":\"\\n    query questionEditorData($titleSlug: String!) {\\n  question(titleSlug: $titleSlug) {\\n    questionId\\n    questionFrontendId\\n    codeSnippets {\\n      lang\\n      langSlug\\n      code\\n    }\\n    envInfo\\n    enableRunCode\\n  }\\n}\\n    \",\"variables\":{\"titleSlug\":\"%s\"}}"
	questionDetail := utils.HttpPost(`https://leetcode.cn/graphql/`, utils.Cookies, utils.HeaderMap, fmt.Sprintf(reqBody, titleSlug))

	var resp type_def.CodeTemplateResp
	err := json.Unmarshal(questionDetail, &resp)
	if err != nil {
		println(utils.GetColorRed(err.Error()))
		return
	}

	for _, template := range resp.Data.Question.CodeSnippets {
		if template.LangSlug == "golang" {
			golangCodeTemplate = template
			break
		}
	}
	return
}

var sure bool

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().BoolVarP(&sure, "yes", "y", false, "确认创建")
}
