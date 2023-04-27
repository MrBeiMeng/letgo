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
	"letgo_repo/system_file/utils/logger"
	"strings"
	"time"
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "展示待做计划,series 表示一个系列，例如本周的任务，应表示问本系列，所以使用任何参数前请携带--series",
	Long:  `展示待做计划`,
	Run: func(cmd *cobra.Command, args []string) {

		if todoParam.CaseAdd() {

			service.SGroupV1.ServiceTodo.Save(type_def.AddTodo{
				Series:       todoParam.GetSeriesOrDefault(),
				ManifestList: todoParam.Add,
			})

			return
		}

		if todoParam.CaseChangeDefault() {
			if strings.TrimSpace(todoParam.Series) == "" {
				logger.Logger.Break("必须携带--series 参数")
			}

			err := service.SGroupV1.ServiceTodo.ChangeDefaultSeries(todoParam.Series)
			if err != nil {
				panic(err)
			}

			logger.Logger.Info("update default series =  %s succeed !", todoParam.Series)
			return
		}

		// 检查是否有额外参数
		todoSeries := service.SGroupV1.ServiceTodo.GetList(type_def.QueryWrapper{
			todoParam.GetSeriesOrDefault(),
		})

		for _, todoSeries1 := range todoSeries {
			fmt.Printf("************* %s ***************\n", todoSeries1.Series)

			globalDone := 0
			globalTotal := 0
			hardCount := 0
			mediumCount := 0
			easyCount := 0
			creatTime := time.Now()

			for _, todo := range todoSeries1.Todos {
				for _, todoQuestion := range todo.TodoQuestions {
					if todoQuestion.Status == enum.DONE {
						globalDone += 1
					}

					if todoQuestion.Difficulty == "EASY" {
						easyCount += 1
					}
					if todoQuestion.Difficulty == "MEDIUM" {
						mediumCount += 1
					}
					if todoQuestion.Difficulty == "HARD" {
						hardCount += 1
					}

					globalTotal += 1
				}
				creatTime = todo.CreatedAt
			}

			fmt.Printf("困难题: %s \t中等题: %s \t简单题: %s \t  完成与总数 [%d/%d] 已创建[%s]天\n", utils.GetColorRed(fmt.Sprintf("%d", hardCount)), utils.GetColorYellow(fmt.Sprintf("%d", mediumCount)), utils.GetColorGreen(fmt.Sprintf("%d", easyCount)), globalDone, globalTotal, time.Now().Sub(creatTime))

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
					switch todoQuestion.Status {
					case enum.DONE:
						numSliceStr += "✌️"
					case enum.INITIALIZED:
						numSliceStr += "☠️"
					}
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
					if i <= (doneNum/totalNum)*5 {
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
	//todoCmd.Flags().StringVar(&todoParam.CreateSeries, "create_series", "", "创建一个系列,后接系列名称，系列名不可重复")
	todoCmd.Flags().StringSliceVarP(&todoParam.Add, "add", "a", nil, "为系列添加清单 example: letgo '系列名' todo add ")
	todoCmd.Flags().BoolVarP(&todoParam.Default, "default", "d", false, "设置默认的系列名 ")
	todoCmd.PersistentFlags().StringVarP(&todoParam.Series, "series", "s", "", "指定一个清单名 和 add 一起使用")
}
