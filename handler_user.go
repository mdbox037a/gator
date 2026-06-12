package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.New("Error: please provide username")
	}

	err := s.CurrentConfig.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("Error: failed to set username - %v", err)
	}

	fmt.Print("Info: username set successfully")
	return nil
}
