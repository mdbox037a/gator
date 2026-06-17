package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mdbox037a/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return errors.New("Error: please provide a feed name and feed url")
	}
	feedName := strings.TrimSpace(cmd.Args[0])
	if feedName == "" {
		return errors.New("Error: feed name should not be blank")
	}
	feedURL := strings.TrimSpace(cmd.Args[1])
	if _, err := url.ParseRequestURI(feedURL); err != nil {
		return fmt.Errorf("Error: specified feed url '%s' is not valid", feedURL)
	}

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.currentConfig.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error: failed to query user info for %s in database - %v", s.currentConfig.CurrentUserName, err)
		// I suppose this shouldn't ever happen, because setting the user earlier will have failed first
		// then again, maybe something has gone wrong in the DB in the meantime
	}

	feedArgs := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(ctx, feedArgs)
	if err != nil {
		return fmt.Errorf("Error: failed to add feed to database - %v", err)
	}

	feedContents, err := json.MarshalIndent(feed, "", "    ")
	if err != nil {
		fmt.Errorf("Error: failed to marshal new RSSFeed contents to print - %v", err)
	}
	fmt.Printf("Debug: new RSSFeed contents:\n%s\n", string(feedContents))

	return nil
}
