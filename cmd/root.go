package cmd

import (
	"telegram-search-cli/utils"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "telegram-search-cli",
	Short: "A powerful CLI tool for searching Telegram content",
	Long: `Telegram Search CLI - A comprehensive command line interface tool for 
searching and discovering Telegram channels, groups, bots, and messages.

Features:
• Search for Telegram channels, groups, bots, and messages
• Export results in multiple formats (JSON, CSV, TXT)
• Progress indicators and visual feedback
• Human-readable error messages and suggestions
• Configurable pagination and result limits

Built with modern CLI design principles for the best user experience.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintHeader("Telegram Search CLI")
		utils.PrintInfo("Welcome! This tool helps you search Telegram content efficiently.")
		utils.PrintInfo("")
		utils.PrintInfo("Quick start:")
		utils.PrintInfo("  telegram-search-cli search --channel 'technology' --limit 10")
		utils.PrintInfo("  telegram-search-cli search --group 'programming' --export json")
		utils.PrintInfo("  telegram-search-cli search --bot 'news' --pages 3")
		utils.PrintInfo("")
		utils.PrintInfo("Use 'telegram-search-cli --help' for more information")
		utils.PrintInfo("Use 'telegram-search-cli search --help' for search options")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {

}
