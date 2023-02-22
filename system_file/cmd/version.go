package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"letgo_repo/system_file/cmd/param_def"
	"letgo_repo/system_file/utils"
)

var versionWrapper param_def.VersionWrapper

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVarP(&versionWrapper.Detail, "detail", "d", false, "历史版本与版本公告")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "letgo's version",
	Long:  `当前项目的版本信息`,
	Run: func(cmd *cobra.Command, args []string) {
		var versionBody param_def.VersionBody

		err := versionBody.InitByJsonFile("version.json")
		if err != nil {
			panic(err)
		}

		// 读取version.json 文件.
		// 打印当前版本
		if versionWrapper.Detail {
			detailFormatTemplate := "${project_name} v${version_no} -- ${type}[${date}] | ${log}"

			for i, perVersion := range versionBody.Version {
				argsMap := make(map[string]string)
				argsMap["project_name"] = versionBody.ProjectName
				argsMap["version_no"] = perVersion.VersionNo
				argsMap["type"] = perVersion.Type
				argsMap["date"] = perVersion.Date
				argsMap["log"] = perVersion.Log

				if i == 0 {
					fmt.Printf("\n🔎 ")
				} else {
					fmt.Printf("   ")
				}

				fmt.Println(utils.ReplaceAll(detailFormatTemplate, argsMap))
			}

		} else {
			lastVersion := versionBody.Version[0]
			fmt.Printf("\n🔎 v%s -- %s\n", lastVersion.VersionNo, lastVersion.Type)
		}

		println()
		// [版本命名规范:https://blog.csdn.net/waynelu92/article/details/73604172]
	},
}
