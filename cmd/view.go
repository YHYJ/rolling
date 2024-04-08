/*
File: view.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-17 16:08:21

Description: 执行子命令 'view'
*/

package cmd

import (
	"github.com/spf13/cobra"
	"github.com/yhyj/rolling/cli"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View system information",
	Long:  `View the collected system installation and update information.`,
	Run: func(cmd *cobra.Command, args []string) {
		cli.SystemInfo()
	},
}

func init() {
	viewCmd.Flags().BoolP("help", "h", false, "help for view command")
	rootCmd.AddCommand(viewCmd)
}
