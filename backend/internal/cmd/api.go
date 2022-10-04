package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Mock exam API",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api function called")
		Runner()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
