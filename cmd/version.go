package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of telegram-search-cli",
	Long:  `All software has versions. This is telegram-search-cli's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("telegram-search-cli v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
