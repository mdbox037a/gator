package main

import (
	"context"
	"encoding/xml"
	"fmt"
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
