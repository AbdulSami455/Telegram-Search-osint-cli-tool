package cmd

import (
	"telegram-search-cli/utils"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of telegram-search-cli",
	Long:  `All software has versions. This is telegram-search-cli's`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintInfo("telegram-search-cli v1.0.0")
		utils.PrintInfo("Built with modern CLI design principles")
		utils.PrintInfo("For updates and documentation, visit the project repository")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
