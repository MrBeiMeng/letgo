/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"letgo_repo/system_file/cmd/param_def"
	model_type_def "letgo_repo/system_file/data_access/manifest/type_def"
	"letgo_repo/system_file/service"
	"letgo_repo/system_file/service/manifest/type_def"
	"letgo_repo/system_file/utils"
	"letgo_repo/system_file/utils/logger"
	"strings"
)

// manifestCmd represents the manifest command
var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "创建一个刷题清单 之后你可以将清单添加到你的 todos 中 ",
	Long:  `创建一个刷题清单 之后你可以将清单添加到你的 todos 中 `,
	Run: func(cmd *cobra.Command, args []string) {
		if manifestParam.CaseAdd() {
			// title,frontIds,tags,mark

			dataList := utils.RoughSplit(manifestParam.Add)

			if len(dataList) < 2 {
				println("命令add，提供的参数太少")
				return
			}

			frontIds := strings.Trim(dataList[1], "[]")
			tmpManifest := type_def.Manifest{
				Title:             dataList[0],
				QuestionsFrontIds: strings.Split(frontIds, ","),
			}

			if 2 < len(dataList) {
				if tmpManifest.TagMap == nil {
					tmpManifest.TagMap = make(map[string]struct{})
				}

				dataList[2] = strings.Trim(dataList[2], "[]")
				for _, tag := range strings.Split(dataList[2], ",") {
					tmpManifest.TagMap[tag] = struct{}{}
				}
			}

			if 3 < len(dataList) {
				tmpManifest.Mark = dataList[3]
			}

			// 构造对象，调用方法
			service.SGroupV1.ServiceManifest.Save(tmpManifest)

			logger.Logger.Success("创建成功[%s]共[%d]条记录", tmpManifest.Title, len(tmpManifest.QuestionsFrontIds))
			return
		}

		if manifestParam.CaseShow() {
			// 获取标题
			// 打印
			println(manifestParam.Show)
			manifests := service.SGroupV1.ServiceManifest.Get(model_type_def.QueryWrapper{TitleSlice: []string{manifestParam.Show}})
			if len(manifests) == 0 {
				logger.Logger.Break("标题输入有误")
				return
			}

			targetManifest := manifests[0]
			fmt.Printf("正在打印 %s 的题目列表mark[%s] | tags[%s] | 题量[%d]\n", utils.GetColorCyan(targetManifest.Title), targetManifest.Mark, targetManifest.GetTags(), len(targetManifest.QuestionsFrontIds))

			questions := service.SGroupV1.CodeServiceI.GetByCodeNums(targetManifest.QuestionsFrontIds)
			questions.Print()
			return
		}

		// 什么参数都不带的时候
		printTable := make([][]string, 0)
		printTable = append(printTable, []string{"标题", "标记", "标签", "题目数量"})
		manifests := service.SGroupV1.ServiceManifest.GetList()

		for _, manifest := range manifests {
			row := []string{manifest.Title, manifest.Mark, manifest.GetTags(), fmt.Sprint(len(manifest.QuestionsFrontIds))}

			printTable = append(printTable, row)
		}

		utils.TablePrint(printTable, true) // 修改打印方式
		return
	},
}

var manifestParam param_def.ManifestCmdWrapper

func init() {
	rootCmd.AddCommand(manifestCmd)
	manifestCmd.Flags().StringVar(&manifestParam.Add, "add", "", "添加题目至清单 title,frontIds,tags,mark(unneeded)")
	manifestCmd.Flags().StringVar(&manifestParam.Remove, "remove", "", "清除清单中的题目") // todo 逻辑错误
	manifestCmd.Flags().StringSliceVar(&manifestParam.Set, "set", nil, "将清单题目列表重置成传入参数")
	manifestCmd.Flags().BoolVar(&manifestParam.Create, "create", false, "创建清单")
	manifestCmd.Flags().StringVarP(&manifestParam.Show, "show", "s", "", "打印") // 打印某个清单
	manifestCmd.Flags().BoolVar(&manifestParam.Clear, "clear", false, "清空这个清单")
	//manifestCmd.PersistentFlags().BoolVarP(&detail, "detail", "d", false, "详细打印")
}
