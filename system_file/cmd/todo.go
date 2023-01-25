// Package cmd /*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"letgo_repo/system_file/cmd/param_def"
	"letgo_repo/system_file/service"
	"letgo_repo/system_file/service/todo/type_def"
	"letgo_repo/system_file/utils"
	"letgo_repo/system_file/utils/enum"
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "展示待做计划",
	Long:  `展示待做计划`,
	Run: func(cmd *cobra.Command, args []string) {
		if todoParam.CaseCreateSeries() {
			// 创建一个系列
			series := todoParam.CreateSeries
			service.ServiceGroupV1.CreateSeries(type_def.SeriesWrapper{Series: series})
			return
		}

		if todoParam.CaseAdd() {
			if todoParam.Series == "" {
				fmt.Printf("%s", utils.GetColorRed("请使用series 字段指定目标系列"))
				return
			}

			service.ServiceGroupV1.ServiceTodo.Save(type_def.AddTodo{
				Series:       todoParam.Series,
				ManifestList: todoParam.Add,
			})

			return
		}

		// 检查是否有额外参数
		todoSeries := service.ServiceGroupV1.ServiceTodo.GetList(type_def.QueryWrapper{})

		for _, todoSeries1 := range todoSeries {
			fmt.Printf("************* %s ***************\n", todoSeries1.Series)

			strTable := make([][]string, 0)
			strTable = append(strTable, []string{"标题", "标签", "题目列表", "进度条"})

			for _, todo := range todoSeries1.Todos {
				if todo.ManifestTitle == "" {
					continue
				}

				doneNum := 0
				totalNum := 0
				// ▰▱_ •°↗→↘⇘⇗⇒

				progressStr := ""
				numSliceStr := ""

				for _, todoQuestion := range todo.TodoQuestions {
					totalNum += 1
					if todoQuestion.Status == enum.DONE {
						doneNum += 1
					}

					if len(numSliceStr) != 0 {
						numSliceStr += ","
					}
					numSliceStr += todoQuestion.FrontendQuestionId
					switch todoQuestion.Difficulty {
					case "EASY":
						numSliceStr += "⇘"
					case "MEDIUM":
						numSliceStr += "⇒"
					case "HARD":
						numSliceStr += "⇗"
					}
				}

				for i := 1; i <= 5; i++ {
					if i <= doneNum {
						progressStr += "▰"
						continue
					}
					progressStr += "_"
				}

				progressStr = fmt.Sprintf("[%s]%d/%d", progressStr, doneNum, totalNum)

				strTable = append(strTable, []string{todo.ManifestTitle, todo.ManifestTag, numSliceStr, fmt.Sprintf("%s", progressStr)})
			}

			utils.TablePrint(strTable, true)
		}

	},
}

var todoParam param_def.TodoCmdWrapper

func init() {
	rootCmd.AddCommand(todoCmd)
	todoCmd.Flags().StringVar(&todoParam.CreateSeries, "create_series", "", "创建一个系列,后接系列名称，系列名不可重复")
	todoCmd.Flags().StringSliceVarP(&todoParam.Add, "add", "a", nil, "为系列添加清单 example: ... todo add ")
	todoCmd.PersistentFlags().StringVarP(&todoParam.Series, "series", "s", "", "指定一个清单名 和 add 一起使用")
}
