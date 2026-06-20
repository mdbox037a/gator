package main

import (
	"context"
	"fmt"

	"github.com/mdbox037a/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		ctx := context.Background()
		user, err := s.db.GetUser(ctx, s.currentConfig.CurrentUserName)
		if err != nil {
			return fmt.Errorf("Error: failed to retrieve user info for %s - %v", s.currentConfig.CurrentUserName, err)
		}

		err = handler(s, cmd, user)
		if err != nil {
			return fmt.Errorf("Error: call to handler function failed - %v", err)
		}
		return nil
	}
}
