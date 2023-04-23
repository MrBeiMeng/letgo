package cmd

import (
	"github.com/spf13/cobra"
	"letgo_repo/system_file/cmd/param_def"
	"letgo_repo/system_file/service"
)

var versionWrapper param_def.VersionWrapper

func init() {
	rootCmd.AddCommand(versionCmd)
	versionCmd.Flags().BoolVarP(&versionWrapper.Detail, "detail", "d", false, "历史版本与版本公告")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "The version command is used to display version information.",
	Long:  `version 命令揭示软件的来历与历程，流转在岁月长河中的沧桑变化。`,
	Run: func(cmd *cobra.Command, args []string) {

		// 读取version.json 文件.
		// 打印当前版本
		histories := service.SGroupV1.CommonService.GetVersionHistories()

		for i, versionInfo := range histories {
			if !versionWrapper.Detail && i > 0 {
				break
			}

			println(versionInfo)
		}
		// [版本命名规范:https://blog.csdn.net/waynelu92/article/details/73604172]

	},
}
