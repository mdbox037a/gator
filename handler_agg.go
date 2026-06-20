package main

import (
	"context"
	"encoding/xml"
	"errors"
	"fmt"
	"time"
)

const feedURL string = "https://www.wagslane.dev/index.xml"

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Error: please supply an update interval (for example 1s, 2m, 8h, etc...)")
	}

	time_between_reqs, _ := time.ParseDuration(cmd.Args[0])
	fmt.Printf("Info: collecting RSS feeds every %v\n", time_between_reqs)
	ticker := time.NewTicker(time_between_reqs)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

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
