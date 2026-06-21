package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mdbox037a/gator/internal/database"
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

	for _, item := range feedContent.Channel.Item {
		parsedTime, err := convertRSSTime(item.PubDate)
		if err != nil {
			return fmt.Errorf("Error: could not parse post pulished date - %v", err)
		}
		var desc sql.NullString
		if item.Description != "" {
			desc = sql.NullString{
				String: item.Description,
				Valid:  true,
			}
		} else {
			desc = sql.NullString{
				Valid: false,
			}
		}
		postData := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         feed.Url,
			Description: desc,
			PublishedAt: parsedTime,
			FeedID:      feed.ID,
		}

		_, err = s.db.CreatePost(ctx, postData)
		if err != nil {
			if !strings.Contains(err.Error(), "duplicate key") {
				return fmt.Errorf("Error: failed to add post '%s' to database - %v", postData.Title, err)
			}
		}
	}

	// debug: printout
	feedPrintout, err := xml.MarshalIndent(feedContent, "", "    ")
	if err != nil {
		fmt.Errorf("Error: failed to marshal RSSFeed contents to print - %v", err)
	}
	fmt.Printf("Debug: RSSFeed contents:\n%s\n", string(feedPrintout))
	return nil
}

func convertRSSTime(pubDate string) (sql.NullTime, error) {
	parsedTime, err := time.Parse(time.RFC1123Z, pubDate)
	if err != nil {
		parsedTime, err = time.Parse(time.RFC1123, pubDate)
		if err != nil {
			return sql.NullTime{Valid: false}, err
		}
		return sql.NullTime{Time: parsedTime, Valid: true}, nil
	}
	return sql.NullTime{Time: parsedTime, Valid: true}, nil
}
