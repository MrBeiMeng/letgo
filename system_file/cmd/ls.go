package cmd

import (
	"github.com/spf13/cobra"
	"letgo_repo/system_file/service"
	"letgo_repo/system_file/service/type_def"
	"letgo_repo/system_file/utils"
)

/*
ls
*/

// 子参数注册
func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().StringVarP(&title, "title", "t", "", "您欲搜索的题目")
	lsCmd.Flags().IntVarP(&no, "no", "n", 0, "您欲搜索的题号")
	lsCmd.Flags().StringVarP(&level, "level", "l", "", "题目难度等级")
	lsCmd.Flags().StringVarP(&star, "star", "r", "", "题目评价")
	lsCmd.Flags().StringVarP(&status, "status", "s", "", "完成情况")
	lsCmd.Flags().StringSliceVarP(&tags, "tags", "g", make([]string, 0), "标签筛选")
	lsCmd.PersistentFlags().BoolVarP(&detail, "detail", "d", false, "详细打印")
	//lsCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "搜索完整题目") // todo
	lsCmd.PersistentFlags().BoolVarP(&hidden, "hidden", "i", false, "显示隐藏题目")
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "打印题目列表",
	Long:  `可以添加过滤条件`,
	// 具体执行代码
	Run: func(cmd *cobra.Command, args []string) {
		queryWrapper := getQueryWrapper()

		projects := service.CodeService.Search(queryWrapper)
		if len(projects) == 0 {
			utils.TPrint("没有符合条件的条目")
			return
		}

		if detail {
			projects.Print()
			return
		}

		projects.EasyPrint()
	},
}

func getQueryWrapper() type_def.CodeQueryWrapper {
	return type_def.CodeQueryWrapper{
		Level:      level,
		Star:       star,
		Status:     status,
		ShowHidden: hidden,
		CodeTitle:  title,
		CodeNum:    no,
		Tags:       tags,
	}
}

var level string
var star string
var status string
var hidden bool
var all bool
var no int
var title string
var tags []string

var detail bool
