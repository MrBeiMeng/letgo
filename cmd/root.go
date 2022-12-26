package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "letgo",
	Short: "letgo 是一个方便的本地刷题工具",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		println("你好，先生,这里是letgo ，一个轻量的本地刷题工具。我可以干什么？\n - 可以自动导入题目\n - 快速分类查看统计 ")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, err := fmt.Fprintln(os.Stderr, err)
		if err != nil {
			panic(err.Error())
		}
		os.Exit(1)
	}
}
