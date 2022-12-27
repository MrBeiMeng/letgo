// Package cmd /*
package cmd

import (
	"github.com/spf13/cobra"
	"letgo_repo/service"
	"letgo_repo/utils"
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

		codeInfo := service.CodeService.SearchInDBByNo(codeNum)
		if codeInfo.CodeNum == 0 {
			println("查无此题")
			return
		}

		//reader := bufio.NewReader(os.Stdin)
		//line, _, _ := reader.ReadLine()
		//inputStr := string(line)
		//if strings.EqualFold(inputStr, "yes") {
		//
		//}

		utils.InitFile(codeInfo.EnglishTitleSlug, codeInfo.Url, codeInfo.Title, codeNum)

		service.CodeService.InitTodoCode(codeNum)
	},
}

var sure bool

func init() {
	rootCmd.AddCommand(initCmd)
	//initCmd.PersistentFlags().BoolVarP(&sure, "yes", "y", false, "确认创建")
}
