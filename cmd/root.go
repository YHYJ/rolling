/*
File: root.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-17 16:08:19

Description: 程序未带子命令或参数时执行
*/

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rolling",
	Short: "For statistics and output system installation and update information",
	Long:  `rolling is a system installation and update statistics tool for Arch Linux.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// 定义全局Flag
	rootCmd.Flags().BoolP("help", "h", false, "help for rolling")
}
