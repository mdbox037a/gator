package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mdbox037a/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Error: please provide username")
	}

	ctx := context.Background()
	_, err := s.db.GetUser(ctx, cmd.Args[0])
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("Error: user %s not found in database; please try again or register new user", cmd.Args[0])
		}
		return errors.New("Error: failed to query database")
	}

	err = s.currentConfig.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Error: failed to set username - %v", err)
	}

	fmt.Println("Info: username set successfully")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Error: please provide a username")
	}

	ctx := context.Background()
	userArgs := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      cmd.Args[0],
	}
	user, err := s.db.CreateUser(ctx, userArgs)
	if err != nil {
		return fmt.Errorf("Error: failed to add user to database %s - %v", userArgs.Name, err)
	}
	err = s.currentConfig.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Error: failed to set username in gatorconfig - %v", err)
	}

	debugUserInfo, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		fmt.Printf("Warning: failed to marshal user info for debugging")
	}
	fmt.Printf("Info: user %s successfully created", user.Name)
	fmt.Printf("Debug: new user info: \n%s\n", string(debugUserInfo))

	return nil
}
