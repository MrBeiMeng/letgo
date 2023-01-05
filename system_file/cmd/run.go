// Package cmd
package cmd

import (
	"github.com/spf13/cobra"
	"letgo_repo/system_file/service"
	"strconv"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "运行或测试方法",
	Long:  `你可以不传参数表示运行demo，或者传递参数测试进行黑盒测试`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			println("请传入题号")
			return
		}

		codeNum, _ := strconv.Atoi(args[0])

		if rightAnswer != "" {
			service.CodeService.Run(codeNum, userArgs, saveAll, rightAnswer)
			return
		}

		service.CodeService.Run(codeNum, userArgs, saveAll)
	},
}

var userArgs string
var rightAnswer string
var saveAll bool

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&userArgs, "args", "a", "", "参数列表")
	runCmd.Flags().StringVarP(&rightAnswer, "rightAnswer", "r", "", "正确结果")
	runCmd.PersistentFlags().BoolVarP(&saveAll, "saveAll", "s", false, "保存所有结果")
}