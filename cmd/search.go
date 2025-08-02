package cmd

import (
	"fmt"
	"strings"

	"telegram-search-cli/scraper"

	"github.com/spf13/cobra"
)

var (
	channelSearchFlag bool
	groupSearchFlag   bool
	limitFlag         int
	pagesFlag         int
	perPageFlag       int
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search for content on Telegram",
	Long: `Search for messages, channels, groups, or users on Telegram.
Example: telegram-search-cli search --channel "pakistan" --limit 12
Example: telegram-search-cli search --group "pakistan" --limit 12`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		if pagesFlag <= 0 {
			pagesFlag = 1
		}
		if perPageFlag <= 0 {
			perPageFlag = 10
		}

		if channelSearchFlag {
			channels, err := scraper.GetChannels(query, pagesFlag, perPageFlag, limitFlag)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Printf("Found %d channels for keyword '%s':\n", len(channels), query)
			for i, ch := range channels {
				fmt.Printf("\n%d. %s\n", i+1, ch.Title)
				if ch.URL != "" {
					fmt.Printf("   URL: %s\n", ch.URL)
				}
				if ch.Description != "" {
					fmt.Printf("   Description: %s\n", ch.Description)
				}
				fmt.Println(strings.Repeat("-", 40))
			}
			return
		}

		if groupSearchFlag {
			groups, err := scraper.GetGroups(query, pagesFlag, perPageFlag, limitFlag)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			fmt.Printf("Found %d groups for keyword '%s':\n", len(groups), query)
			for i, group := range groups {
				fmt.Printf("\n%d. %s\n", i+1, group.Title)
				if group.URL != "" {
					fmt.Printf("   URL: %s\n", group.URL)
				}
				if group.Description != "" {
					fmt.Printf("   Description: %s\n", group.Description)
				}
				fmt.Println(strings.Repeat("-", 40))
			}
			return
		}

		fmt.Printf("Searching for: %s\n", query)
		fmt.Println("This is a placeholder for actual Telegram search functionality.")
		fmt.Println("In a real implementation, this would connect to Telegram API.")
		fmt.Println("Use --channel or --group flag to search for channels or groups.")
	},
}

func init() {
	searchCmd.Flags().BoolVarP(&channelSearchFlag, "channel", "c", false, "Search for Telegram channels")
	searchCmd.Flags().BoolVarP(&groupSearchFlag, "group", "g", false, "Search for Telegram groups")
	searchCmd.Flags().IntVarP(&limitFlag, "limit", "l", 0, "Maximum number of results to return (0 = no limit)")
	searchCmd.Flags().IntVarP(&pagesFlag, "pages", "p", 1, "Number of pages to fetch")
	searchCmd.Flags().IntVarP(&perPageFlag, "per-page", "n", 10, "Results per page")
	rootCmd.AddCommand(searchCmd)
}
