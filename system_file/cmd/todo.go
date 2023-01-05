// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"letgo_repo/system_file/service"
	"letgo_repo/system_file/utils"
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "展示待做计划",
	Long:  `展示待做计划`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Printf("主题\t进度\t题目列表\n")
		todos := service.CodeService.GetToDos()

		maxLength := 0

		for _, toDoQuestion := range todos {
			maxLength = max(len(toDoQuestion.Theme), maxLength)
		}

		for _, todoQuestion := range todos {
			if todoQuestion.Master {
				fmt.Printf(">%s\t%s\t%s\n", ConvLengthTo(todoQuestion.Theme, maxLength), todoQuestion.Progress, utils.GetColorWhite("*********************"))
				continue
			}

			fmt.Printf("%s\t%s\t%s\n", ConvLengthTo(todoQuestion.Theme, maxLength), todoQuestion.Progress, todoQuestion.CodeNums)
		}

	},
}

func ConvLengthTo(str string, length int) string {
	for i := 0; i < (length-len(str))/9; i++ {
		str += "\t"
	}
	return str
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func init() {
	rootCmd.AddCommand(todoCmd)
}
