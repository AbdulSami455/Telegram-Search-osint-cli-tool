package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "telegram-search-cli",
	Short: "A simple CLI tool for Telegram search",
	Long: `A command line interface tool built with Cobra for 
performing various Telegram-related operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Telegram Search CLI!")
		fmt.Println("Use --help to see available commands.")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	
}
