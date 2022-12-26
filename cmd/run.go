// Package cmd
package cmd

import (
	"github.com/spf13/cobra"
	"letgo_repo/service"
	"letgo_repo/service/type_def"
	"strconv"
	"strings"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "运行或测试方法",
	Long:  `你可以不传参数表示运行demo，或者传递参数测试进行黑盒测试`,
	Run: func(cmd *cobra.Command, args []string) {
		var dataStructureArgs type_def.Args

		if len(args) == 0 {
			println("请传入题号")
			return
		}

		codeNum, _ := strconv.Atoi(args[0])

		if !strings.EqualFold(linkedLists, "") {
			dataStructureArgs.LinkedLists = linkedLists
		}

		if dataStructureArgs.IsEmpty() {
			service.CodeService.RunDemo(codeNum)
			return
		}

		service.CodeService.Run(codeNum, dataStructureArgs)
	},
}

var linkedLists string

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&linkedLists, "linked_list", "l", "", "传入n个链表")
}
