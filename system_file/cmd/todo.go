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
	"strconv"
	"strings"
	"time"
)

var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "展示待做计划,series 表示一个系列，例如本周的任务，应表示问本系列，所以使用任何参数前请携带--series",
	Long:  `展示待做计划`,
	Run: func(cmd *cobra.Command, args []string) {

		if todoParam.CaseAdd() {
			// todo 添加系列的同时应该可以支持设置默认，并且创建好对应的文件 包括cmd 文件夹下的enter.go
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
			Series: todoParam.GetSeriesOrDefault(),
		})

		for _, todoSeries1 := range todoSeries {
			fmt.Printf("************* %s 系列 ***************\n\n", todoSeries1.Series)

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
				if creatTime.After(todo.CreatedAt) {
					creatTime = todo.CreatedAt
				}
			}

			fmt.Printf("此系列被创建已过去[%s]\n", time.Now().Sub(creatTime))

			fmt.Printf("分析: 困难题: %s \t中等题: %s \t简单题: %s \t  完成与总数 [%d/%d]\n\n", utils.GetColorRed(fmt.Sprintf("%d", hardCount)), utils.GetColorYellow(fmt.Sprintf("%d", mediumCount)), utils.GetColorGreen(fmt.Sprintf("%d", easyCount)), globalDone, globalTotal)

			strTable := make([][]string, 0)
			strTable = append(strTable, []string{"no.", "标题", "level", "url", "ctn", "tags", "status"}) // 题目打印数组
			for _, todo := range todoSeries1.Todos {
				if todo.ManifestTitle == "" {
					continue
				}
				doneNum := 0
				totalNum := 0
				// ▰▱_ •°↗→↘⇘⇗⇒

				progressStr := ""

				for _, todoQuestion := range todo.TodoQuestions {
					codeNum, _ := strconv.Atoi(todoQuestion.FrontendQuestionId)
					todoQuestionDetail := service.SGroupV1.CodeServiceI.GetByCodeNum(codeNum)
					totalNum += 1
					if todoQuestion.Status == enum.DONE {
						doneNum += 1
					}

					switch todoQuestion.Status {
					case enum.DONE:
						todoQuestion.Status += "✌️"
					case enum.INITIALIZED:
						todoQuestion.Status += "☠️"
					}

					strTable = append(strTable, []string{todoQuestion.FrontendQuestionId, todoQuestionDetail.TitleCn, todoQuestionDetail.Difficulty, todoQuestionDetail.Url, fmt.Sprintf("%d", todoQuestionDetail.CompanyTagNum), todoQuestionDetail.GetTags(), todoQuestion.Status})
				}

				for i := 1; i <= 26; i++ {
					if totalNum > 0 && i <= (doneNum*26/totalNum) {
						progressStr += "▰"
						continue
					}
					progressStr += "_"
				}
				fmt.Printf("清单:%s 进度: [%s]%d/%d 创建时间:[%s]\n", utils.GetColorCyan(todo.ManifestTitle), progressStr, doneNum, totalNum, todo.Model.CreatedAt.Format("2006-01-02 15:04:05"))
				utils.TablePrint(strTable, true)
				strTable = make([][]string, 0)
			}
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
