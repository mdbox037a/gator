package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/mdbox037a/gator/internal/database"
)

const feedURL string = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	ctx := context.Background()
	feed, err := fetchFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("Error: failed to retrieve RSS feed from %s - %v", feedURL, err)
	}
	feedContents, err := xml.MarshalIndent(feed, "", "    ")
	if err != nil {
		fmt.Errorf("Error: failed to marshal RSSFeed contents to print - %v", err)
	}
	fmt.Printf("Debug: RSSFeed contents:\n%s\n", string(feedContents))
	return nil
}

func scrapeFeeds(s *state) error {
	ctx := context.Background()
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("Error: failed to get next feed to fetch from database - %v", err)
	}

	err = s.db.MarkFeedFetched(ctx, feed.ID)
	if err != nil {
		return fmt.Errorf("Error: failed to mark feed '%s' as fetched - %n", feed.Name, err)
	}

	feedContent, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return fmt.Errorf("Error: failed to retrieve RSS feed from %s - %v", feed.Url, err)
	}

	feedPrintout, err := xml.MarshalIndent(feedContent, "", "    ")
	if err != nil {
		fmt.Errorf("Error: failed to marshal RSSFeed contents to print - %v", err)
	}
	fmt.Printf("Debug: RSSFeed contents:\n%s\n", string(feedPrintout))
	return nil
}
