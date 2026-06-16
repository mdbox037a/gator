package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db.DeleteUsers(ctx)
	if err != nil {
		return fmt.Errorf("Error: failed to reset gator users database - %v", err)
	}
	fmt.Print("Info: gator users database reset successully")
	return nil
}
