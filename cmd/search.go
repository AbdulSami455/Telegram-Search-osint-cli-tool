package cmd

import (
	"fmt"
	"math"
	"strings"

	"telegram-search-cli/scraper"
	"telegram-search-cli/utils"

	"github.com/spf13/cobra"
)

var (
	channelSearchFlag bool
	groupSearchFlag   bool
	botSearchFlag     bool
	messageSearchFlag bool
	limitFlag         int
	pagesFlag         int
	perPageFlag       int
	exportFlag        string
	outputFlag        string
)

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search for content on Telegram",
	Long: `Search for messages, channels, groups, bots, or users on Telegram.
Example: telegram-search-cli search --channel "pakistan" --limit 12
Example: telegram-search-cli search --group "pakistan" --limit 12
Example: telegram-search-cli search --bot "sports" --limit 12
Example: telegram-search-cli search --message "pakistan" --limit 100
Example: telegram-search-cli search --bot "sports" --limit 50 --export json --output results.json`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]

		// Validate input
		if strings.TrimSpace(query) == "" {
			utils.PrintError("Search query cannot be empty")
			utils.PrintInfo("Please provide a valid search term")
			return
		}

		// Auto-calculate pages and per-page if limit is specified
		if limitFlag > 0 {
			if pagesFlag <= 0 {
				pagesFlag = int(math.Ceil(float64(limitFlag) / 10.0))
			}
			if perPageFlag <= 0 {
				perPageFlag = 10
			}
		} else {
			// Default values if no limit specified
			if pagesFlag <= 0 {
				pagesFlag = 1
			}
			if perPageFlag <= 0 {
				perPageFlag = 10
			}
		}

		// Show search configuration
		utils.PrintInfo(fmt.Sprintf("Search configuration: %d pages, %d results per page", pagesFlag, perPageFlag))
		if limitFlag > 0 {
			utils.PrintInfo(fmt.Sprintf("Result limit: %d", limitFlag))
		}

		if channelSearchFlag {
			utils.PrintHeader("Channel Search Results")

			channels, err := scraper.GetChannels(query, pagesFlag, perPageFlag, limitFlag)
			if err != nil {
				utils.HandleError(err, "search channels")
				return
			}

			// Export if requested
			if exportFlag != "" {
				utils.PrintInfo("Preparing to export results...")
				data := scraper.ConvertChannelsToInterface(channels)
				if err := utils.ExportToFile(data, query, "channels", exportFlag, outputFlag); err != nil {
					utils.HandleError(err, "export channels")
					return
				}
			}

			// Display results with better formatting
			utils.PrintSubHeader(fmt.Sprintf("Found %d channels for '%s'", len(channels), query))
			if len(channels) == 0 {
				utils.PrintWarning("No channels found for your search query")
				utils.PrintInfo("Try using different keywords or check your search terms")
				return
			}

			for i, ch := range channels {
				fmt.Printf("\n%d. %s\n", i+1, ch.Title)
				if ch.URL != "" {
					fmt.Printf("   🔗 URL: %s\n", ch.URL)
				}
				if ch.Description != "" {
					fmt.Printf("   📝 Description: %s\n", ch.Description)
				}
				fmt.Println(strings.Repeat("-", 50))
			}

			utils.PrintSuccess(fmt.Sprintf("Channel search completed successfully! Found %d results", len(channels)))
			return
		}

		if groupSearchFlag {
			utils.PrintHeader("Group Search Results")

			groups, err := scraper.GetGroups(query, pagesFlag, perPageFlag, limitFlag)
			if err != nil {
				utils.HandleError(err, "search groups")
				return
			}

			// Export if requested
			if exportFlag != "" {
				utils.PrintInfo("Preparing to export results...")
				data := scraper.ConvertGroupsToInterface(groups)
				if err := utils.ExportToFile(data, query, "groups", exportFlag, outputFlag); err != nil {
					utils.HandleError(err, "export groups")
					return
				}
			}

			// Display results with better formatting
			utils.PrintSubHeader(fmt.Sprintf("Found %d groups for '%s'", len(groups), query))
			if len(groups) == 0 {
				utils.PrintWarning("No groups found for your search query")
				utils.PrintInfo("Try using different keywords or check your search terms")
				return
			}

			for i, group := range groups {
				fmt.Printf("\n%d. %s\n", i+1, group.Title)
				if group.URL != "" {
					fmt.Printf("   🔗 URL: %s\n", group.URL)
				}
				if group.Description != "" {
					fmt.Printf("   📝 Description: %s\n", group.Description)
				}
				fmt.Println(strings.Repeat("-", 50))
			}

			utils.PrintSuccess(fmt.Sprintf("Group search completed successfully! Found %d results", len(groups)))
			return
		}

		if botSearchFlag {
			utils.PrintHeader("Bot Search Results")

			bots, err := scraper.GetBots(query, pagesFlag, perPageFlag, limitFlag)
			if err != nil {
				utils.HandleError(err, "search bots")
				return
			}

			// Export if requested
			if exportFlag != "" {
				utils.PrintInfo("Preparing to export results...")
				data := scraper.ConvertBotsToInterface(bots)
				if err := utils.ExportToFile(data, query, "bots", exportFlag, outputFlag); err != nil {
					utils.HandleError(err, "export bots")
					return
				}
			}

			// Display results with better formatting
			utils.PrintSubHeader(fmt.Sprintf("Found %d bots for '%s'", len(bots), query))
			if len(bots) == 0 {
				utils.PrintWarning("No bots found for your search query")
				utils.PrintInfo("Try using different keywords or check your search terms")
				return
			}

			for i, bot := range bots {
				fmt.Printf("\n%d. %s\n", i+1, bot.Title)
				if bot.URL != "" {
					fmt.Printf("   🔗 URL: %s\n", bot.URL)
				}
				if bot.Description != "" {
					fmt.Printf("   📝 Description: %s\n", bot.Description)
				}
				fmt.Println(strings.Repeat("-", 50))
			}

			utils.PrintSuccess(fmt.Sprintf("Bot search completed successfully! Found %d results", len(bots)))
			return
		}

		if messageSearchFlag {
			utils.PrintHeader("Message Search Results")

			messages, err := scraper.GetMessages(query, pagesFlag, perPageFlag, limitFlag)
			if err != nil {
				utils.HandleError(err, "search messages")
				return
			}

			// Export if requested
			if exportFlag != "" {
				utils.PrintInfo("Preparing to export results...")
				data := scraper.ConvertMessagesToInterface(messages)
				if err := utils.ExportToFile(data, query, "messages", exportFlag, outputFlag); err != nil {
					utils.HandleError(err, "export messages")
					return
				}
			}

			// Display results with better formatting
			utils.PrintSubHeader(fmt.Sprintf("Found %d messages for '%s'", len(messages), query))
			if len(messages) == 0 {
				utils.PrintWarning("No messages found for your search query")
				utils.PrintInfo("Try using different keywords or check your search terms")
				return
			}

			for i, msg := range messages {
				fmt.Printf("\n%d. %s\n", i+1, msg.Title)
				if msg.URL != "" {
					fmt.Printf("   🔗 URL: %s\n", msg.URL)
				}
				if msg.Description != "" {
					fmt.Printf("   📝 Description: %s\n", msg.Description)
				}
				fmt.Println(strings.Repeat("-", 50))
			}

			utils.PrintSuccess(fmt.Sprintf("Message search completed successfully! Found %d results", len(messages)))
			return
		}

		utils.PrintWarning("No search type specified")
		utils.PrintInfo("Please specify what you want to search for:")
		utils.PrintInfo("  --channel, -c  : Search for Telegram channels")
		utils.PrintInfo("  --group, -g    : Search for Telegram groups")
		utils.PrintInfo("  --bot, -b      : Search for Telegram bots")
		utils.PrintInfo("  --message, -m  : Search for Telegram messages")
		utils.PrintInfo("")
		utils.PrintInfo("Example: telegram-search-cli search --channel 'pakistan' --limit 12")
		utils.PrintInfo("Example: telegram-search-cli search --group 'technology' --export json")
	},
}

func init() {
	searchCmd.Flags().BoolVarP(&channelSearchFlag, "channel", "c", false, "Search for Telegram channels")
	searchCmd.Flags().BoolVarP(&groupSearchFlag, "group", "g", false, "Search for Telegram groups")
	searchCmd.Flags().BoolVarP(&botSearchFlag, "bot", "b", false, "Search for Telegram bots")
	searchCmd.Flags().BoolVarP(&messageSearchFlag, "message", "m", false, "Search for Telegram messages")
	searchCmd.Flags().IntVarP(&limitFlag, "limit", "l", 0, "Maximum number of results to return (0 = no limit)")
	searchCmd.Flags().IntVarP(&pagesFlag, "pages", "p", 0, "Number of pages to fetch (auto-calculated if limit is specified)")
	searchCmd.Flags().IntVarP(&perPageFlag, "per-page", "n", 0, "Results per page (default: 10)")
	searchCmd.Flags().StringVarP(&exportFlag, "export", "e", "", "Export format: json, csv, txt")
	searchCmd.Flags().StringVarP(&outputFlag, "output", "o", "", "Output filename (auto-generated if not specified)")
	rootCmd.AddCommand(searchCmd)
}
