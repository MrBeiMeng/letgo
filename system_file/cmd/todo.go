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

		table := make([][]string, 0)
		table = append(table, []string{"主题", "题目列表", "进度"})
		for _, toDoQuestion := range service.CodeService.GetToDos() {

			if toDoQuestion.Master {
				table = append(table, []string{fmt.Sprintf("< %s >", toDoQuestion.Theme), "——", toDoQuestion.Progress})
				continue
			}

			table = append(table, []string{toDoQuestion.Theme, toDoQuestion.CodeNums, toDoQuestion.Progress})
		}

		utils.TablePrint(table)
	},
}

func init() {
	rootCmd.AddCommand(todoCmd)
}
