package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/mdbox037a/gator/internal/database"
)

func browse(s *state, cmd command) error {
	limit := 2
	if len(cmd.Args) > 0 {
		temp, err := strconv.Atoi(cmd.Args[0])
		if err != nil || temp < 0 {
			return fmt.Errorf("Error: please provide a positive integer value to browse - %v", err)
		}
		limit = temp
	}

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.currentConfig.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Error: failed to get user info for '%s' from database - %v", s.currentConfig.CurrentUserName, err)
	}
	userBrowse := database.GetUserPostsParams{
		UserID: user.ID,
		Limit:  int32(limit),
	}
	posts, err := s.db.GetUserPosts(ctx, userBrowse)
	if err != nil {
		return fmt.Errorf("Error: failed to get latest posts from user's followed feeds - %v", err)
	}

	for _, post := range posts {
		postContent, err := json.MarshalIndent(post, "", "    ")
		if err != nil {
			return fmt.Errorf("Error: failed to marshal post contents - %v", err)
		}
		fmt.Printf("Post contents:\n%s\n", string(postContent))
	}

	return nil
}
