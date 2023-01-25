/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"letgo_repo/system_file/cmd/param_def"
	"letgo_repo/system_file/service"
	"letgo_repo/system_file/service/manifest/type_def"
	"letgo_repo/system_file/utils"
	"strings"
)

// manifestCmd represents the manifest command
var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "查看题目清单",
	Long:  `查看清单`,
	Run: func(cmd *cobra.Command, args []string) {

		// 当前只提供查询功能

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
			service.CodeServiceGroupV1.ServiceManifest.Save(tmpManifest)

			return
		}

		if manifestParam.CaseShow() {
			printTable := make([][]string, 0)
			printTable = append(printTable, []string{"标题", "标记", "标签", "题目列表"})
			manifests := service.CodeServiceGroupV1.ServiceManifest.GetList()

			for _, manifest := range manifests {
				row := []string{manifest.Title, manifest.Mark, manifest.GetTags(), manifest.GetFrontIdsWithColor()}

				printTable = append(printTable, row)
			}

			utils.TablePrintColorHandler(printTable, []int{3})
		}
	},
}

var manifestParam param_def.ManifestCmdWrapper

func init() {
	rootCmd.AddCommand(manifestCmd)
	manifestCmd.Flags().StringVar(&manifestParam.Add, "add", "", "添加题目至清单")
	manifestCmd.Flags().StringVar(&manifestParam.Remove, "remove", "", "清除清单中的题目")
	manifestCmd.Flags().StringSliceVar(&manifestParam.Set, "set", nil, "将清单题目列表重置成传入参数")
	manifestCmd.Flags().BoolVar(&manifestParam.Create, "create", false, "创建清单")
	manifestCmd.Flags().BoolVar(&manifestParam.Show, "show", false, "打印")
	manifestCmd.Flags().BoolVar(&manifestParam.Clear, "clear", false, "清空这个清单")
	//manifestCmd.PersistentFlags().BoolVarP(&detail, "detail", "d", false, "详细打印")
}
