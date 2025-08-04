package scraper

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"telegram-search-cli/utils"

	"github.com/PuerkitoBio/goquery"
)

type Bot struct {
	Title       string
	URL         string
	Description string
}

func GetBots(keyword string, pages, perPage, totalResultLimit int) ([]Bot, error) {
	baseURL := "https://lyzem.com/search"
	allBots := make([]Bot, 0)
	client := &http.Client{Timeout: 20 * time.Second}

	// Show progress for multi-page searches
	var progressBar *utils.ProgressBar
	if pages > 1 {
		progressBar = utils.NewProgressBar(pages, "Fetching bots")
		utils.PrintInfo(fmt.Sprintf("Searching for bots with keyword: '%s'", keyword))
		utils.PrintInfo(fmt.Sprintf("Will fetch %d pages with %d results per page", pages, perPage))
	}

	for page := 1; page <= pages; page++ {
		if totalResultLimit > 0 && len(allBots) >= totalResultLimit {
			break
		}

		// Update progress
		if progressBar != nil {
			progressBar.Update(page)
		}

		params := url.Values{}
		params.Set("q", keyword)
		params.Set("f", "bots")
		params.Set("l", "")
		params.Set("p", strconv.Itoa(page))
		params.Set("per-page", strconv.Itoa(perPage))

		resp, err := client.Get(baseURL + "?" + params.Encode())
		if err != nil {
			return nil, utils.WrapNetworkError(err, fmt.Sprintf("fetch page %d", page))
		}

		if resp.StatusCode != 200 {
			resp.Body.Close()
			if resp.StatusCode == http.StatusTooManyRequests {
				return nil, utils.WrapHTTPError(resp.StatusCode, "search bots")
			}
			if resp.StatusCode == http.StatusNotFound {
				return nil, utils.WrapNotFoundError("search results", keyword)
			}
			// Continue for other status codes
			continue
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, utils.WrapInternalError("parse HTML", err)
		}

		pageResults := 0
		doc.Find("ul.search-results > li.search-result").EachWithBreak(func(i int, s *goquery.Selection) bool {
			if totalResultLimit > 0 && len(allBots) >= totalResultLimit {
				return false
			}
			titleTag := s.Find("p.search-result-title a")
			title := strings.TrimSpace(titleTag.Text())
			url, _ := titleTag.Attr("href")
			descrTag := s.Find("p.search-result-descr a")
			description := strings.TrimSpace(descrTag.Text())
			allBots = append(allBots, Bot{
				Title:       title,
				URL:         url,
				Description: description,
			})
			pageResults++
			return true
		})

		// Show page results if verbose
		if progressBar != nil && pageResults > 0 {
			utils.PrintInfo(fmt.Sprintf("Page %d: Found %d bots", page, pageResults))
		}
	}

	if progressBar != nil {
		utils.PrintSuccess(fmt.Sprintf("Completed! Found %d bots total", len(allBots)))
	}

	return allBots, nil
}
