package scraper

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

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

	for page := 1; page <= pages; page++ {
		if totalResultLimit > 0 && len(allBots) >= totalResultLimit {
			break
		}
		params := url.Values{}
		params.Set("q", keyword)
		params.Set("f", "bots")
		params.Set("l", "")
		params.Set("p", strconv.Itoa(page))
		params.Set("per-page", strconv.Itoa(perPage))
		resp, err := client.Get(baseURL + "?" + params.Encode())
		if err != nil {
			return nil, fmt.Errorf("failed to fetch page %d: %w", page, err)
		}
		if resp.StatusCode != 200 {
			resp.Body.Close()
			continue
		}
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to parse HTML: %w", err)
		}
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
			return true
		})
	}
	return allBots, nil
}
