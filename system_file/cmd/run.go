// Package cmd
package cmd

import (
	"github.com/spf13/cobra"
	"letgo_repo/system_file/service"
	"letgo_repo/system_file/service/type_def"
	"letgo_repo/system_file/utils"
	"letgo_repo/system_file/utils/logger"
	"strconv"
	"time"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "运行或测试方法",
	Long:  `你可以不传参数表示运行demo，或者传递参数测试进行黑盒测试`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			logger.Logger.Break("请传入题号")
		}

		codeNum, _ := strconv.Atoi(args[0])

		service.CodeService.Run(type_def.RunWrapper{
			CodeNum:     codeNum,
			ArgsStr:     userArgs,
			SaveAll:     saveAll,
			RightAnswer: rightAnswer,
			Done:        done,
		})

		threadPool := utils.ThreadUtil

		if !threadPool.IsAllDone() {
			print("waiting go thread")
			for !threadPool.IsAllDone() {
				time.Sleep(time.Millisecond * 100)
			}
			println("done")
		}
	},
}

var userArgs string
var rightAnswer string
var saveAll bool
var done bool

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&userArgs, "args", "a", "", "参数列表")
	runCmd.Flags().StringVarP(&rightAnswer, "rightAnswer", "r", "", "正确结果")
	runCmd.PersistentFlags().BoolVarP(&saveAll, "saveAll", "s", false, "保存所有结果")
	runCmd.PersistentFlags().BoolVarP(&done, "done", "d", false, "完成当前题目")
}
