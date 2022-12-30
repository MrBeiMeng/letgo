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
		//
		question := service.CodeService.GetByCodeNum(codeNum)
		if question.CodeNum == 0 {
			println("查无此题")
			return
		}

		reqBody := "{\"query\":\"\\n    query questionEditorData($titleSlug: String!) {\\n  question(titleSlug: $titleSlug) {\\n    questionId\\n    questionFrontendId\\n    codeSnippets {\\n      lang\\n      langSlug\\n      code\\n    }\\n    envInfo\\n    enableRunCode\\n  }\\n}\\n    \",\"variables\":{\"titleSlug\":\"%s\"}}"
		questionDetail := utils.HttpPost(`https://leetcode.cn/graphql/`, utils.Cookies, utils.HeaderMap, fmt.Sprintf(reqBody, question.TitleSlug))

		var resp type_def.CodeTemplateResp
		err := json.Unmarshal(questionDetail, &resp)
		if err != nil {
			println(utils.GetColorRed(err.Error()))
			return
		}

		var golangCodeTemplate type_def.CodeTemplate
		for _, template := range resp.Data.Question.CodeSnippets {
			if template.LangSlug != "golang" {
				continue
			}
			golangCodeTemplate = template
		}

		utils.InitFile(question.TitleSlug, question.Url, question.TitleCn, question.CodeNum, golangCodeTemplate.Code)

		fmt.Printf(utils.GetColorGreen("DONE"))
		service.CodeService.OperateLog(fmt.Sprintf("添加新题:%v", codeNum), golangCodeTemplate.Code, enum.INIT_CODE)
	},
}

var sure bool

func init() {
	rootCmd.AddCommand(initCmd)
	//initCmd.PersistentFlags().BoolVarP(&sure, "yes", "y", false, "确认创建")
}
