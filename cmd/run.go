// Package cmd
package cmd

import (
	"github.com/spf13/cobra"
	"letgo_repo/service"
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

		service.CodeService.Run(codeNum, userArgs)
	},
}

var userArgs string

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&userArgs, "args", "a", "", "参数列表")
}
