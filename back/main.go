package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

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
	videoLists, err := fetchVideoLists()
	if err != nil {
		log.Fatalf("Failed to retrieve video list: %v", err)
	}

	// Save to JSON file
	if err := saveVideoListsToJSON(videoLists, "dist/ranking.json"); err != nil {
		log.Fatalf("Failed to save JSON file: %v", err)
	} else {
		log.Println("Successfully saved JSON file: dist/ranking.json")
	}
}

// FetchVideoLists is a function that retrieves video lists from Niconico video rankings
func fetchVideoLists() ([][]Video, error) {
	// Niconico video ranking page URL
	url := "https://www.nicovideo.jp/ranking/custom"

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %v", err)
	}

	// Set User-Agent header (as needed)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")

	// Create HTTP client and execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute HTTP request: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid HTTP status code: %d", resp.StatusCode)
	}

	// Parse HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	// Search for <meta name="server-response" content="..."> element
	var jsonContent string
	doc.Find("meta[name='server-response']").Each(func(i int, s *goquery.Selection) {
		// Get content attribute value
		content, exists := s.Attr("content")
		if exists {
			jsonContent = content
		}
	})

	if jsonContent == "" {
		return nil, fmt.Errorf("server-response meta tag not found")
	}

	// Decode HTML entities (if needed)
	jsonContent = strings.ReplaceAll(jsonContent, "&quot;", "\"")

	// Parse JSON data
	result := gjson.Parse(jsonContent)

	// Extract $getCustomRankingRanking
	rankingData := result.Get("data.response.$getCustomRankingRanking")

	if !rankingData.Exists() {
		return nil, fmt.Errorf("$getCustomRankingRanking data not found")
	}

	videoLists := make([][]Video, 0)

	// Parse Lanes
	rankingData.ForEach(func(_, value gjson.Result) bool {
		if value.Get("data.title").String() == "Other" {
			return true
		}

		videoList := make([]Video, 0)

		// Parse videoList
		value.Get("data.videoList").ForEach(func(_, videoValue gjson.Result) bool {
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

// SaveVideoListsToJSON is a function that saves video lists to a JSON file
func saveVideoListsToJSON(videoLists [][]Video, filePath string) error {
	// Encode to JSON
	jsonData, err := json.MarshalIndent(videoLists, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %v", err)
	}

	// Ensure directory exists
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	// Create file for writing with gzip compression
	file, err := os.Create(filePath + ".gz")
	if err != nil {
		return fmt.Errorf("failed to create gzip file: %v", err)
	}
	defer file.Close()

	// Create gzip writer
	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	// Write JSON data to gzip writer
	if _, err := gzipWriter.Write(jsonData); err != nil {
		return fmt.Errorf("failed to write gzipped data: %v", err)
	}

	// Also save the uncompressed version for backward compatibility
	if err := os.WriteFile(filePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write uncompressed file: %v", err)
	}

	return nil
}
