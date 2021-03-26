package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "center [text]",
	Version: "0.0.3",
	Short:   "A cli tool to center text",
	Long:    "A cli tool to center text",
	Run: func(cmd *cobra.Command, args []string) {
		ExecuteCenter(cmd, args)
	},
}

func Execute() {
	_ = rootCmd.Execute()
}
