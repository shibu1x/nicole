package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
)

type Video struct {
	Id                string
	Title             string
	ThumbnailURL      string
	OwnerId           string
	WatchURL          string
	IsPaymentRequired bool
	IsMuted           bool
}

func main() {
	log.Printf("Starting CLI job at %s", time.Now().Format(time.RFC3339))

	if err := runPeriodicTask(); err != nil {
		log.Fatal(err)
	}

	log.Println("Job completed successfully")
}

func scrapeNicovideoRanking() ([][]Video, error) {
	resp, err := http.Get("https://www.nicovideo.jp/ranking")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ranking page: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var rawData string
	doc.Find("#MatrixRanking-app").Each(func(i int, s *goquery.Selection) {
		if appData, exists := s.Attr("data-app"); exists {
			rawData = appData
		}
	})

	if rawData == "" {
		return nil, fmt.Errorf("ranking data not found")
	}

	// gjson を使用してJSONをパース
	result := gjson.Parse(rawData)

	videoLists := make([][]Video, 0)

	// Lanes の解析
	result.Get("lanes").ForEach(func(_, value gjson.Result) bool {
		if value.Get("title").String() == "その他" {
			return true
		}

		videoList := make([]Video, 0)

		// videoList の解析
		value.Get("videoList").ForEach(func(_, videoValue gjson.Result) bool {
			video := Video{
				Id:                videoValue.Get("id").String(),
				Title:             videoValue.Get("title").String(),
				ThumbnailURL:      videoValue.Get("thumbnail.listingUrl").String(),
				OwnerId:           videoValue.Get("owner.id").String(),
				IsPaymentRequired: videoValue.Get("isPaymentRequired").Bool(),
			}
			video.WatchURL = fmt.Sprintf("https://www.nicovideo.jp/watch/%s", video.Id)
			videoList = append(videoList, video)
			return true
		})

		videoLists = append(videoLists, videoList)
		return true
	})

	return videoLists, nil
}

func runPeriodicTask() error {
	videoList, err := scrapeNicovideoRanking()
	if err != nil {
		return fmt.Errorf("failed to scrape ranking: %w", err)
	}

	// Convert to JSON
	jsonData, err := json.MarshalIndent(videoList, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Create dist directory if it doesn't exist
	if err := os.MkdirAll("dist", 0755); err != nil {
		return fmt.Errorf("failed to create dist directory: %w", err)
	}

	// Save locally
	if err := os.WriteFile("dist/ranking.json", jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}
	log.Println("Ranking data saved to dist/ranking.json")

	return nil
}
