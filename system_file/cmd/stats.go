/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "stats",
	Short: "做题统计",
	Long:  `显示做题信息`,
	Run: func(cmd *cobra.Command, args []string) {
		// 显示总做题量、
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
