package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mdbox037a/gator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Error: please provide a feed URL")
	}
	feedURL := cmd.Args[0]

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.currentConfig.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error: failed to retrieve user from database - %v", err)
	}
	user_id := user.ID
	feed_id, err := s.db.GetFeed(ctx, feedURL)
	if err != nil {
		return fmt.Errorf("Error: failed to retrieve feed from database - %v", err)
	}
	feedFollowParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user_id,
		FeedID:    feed_id,
	}
	feedFollow, err := s.db.CreateFeedFollow(ctx, feedFollowParams)
	if err != nil {
		return fmt.Errorf("Error: failed to link user to feed - %v", err)
	}

	feedFollowContents, err := json.MarshalIndent(feedFollow, "", "    ")
	if err != nil {
		fmt.Errorf("Error: failed to marshal feed follow row for output - %v", err)
	}
	fmt.Printf("Feed '%s' now followed by user '%s'\n%s\n", feedFollow.FeedName, s.currentConfig.CurrentUserName, string(feedFollowContents))

	return nil
}
