package main

import (
	"fmt"

	"github.com/mdbox037a/gator/internal/config"
	"github.com/mdbox037a/gator/internal/database"
)

type state struct {
	db            *database.Queries
	currentConfig *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.Name]
	if !ok {
		return fmt.Errorf("Error: command '%s' does not exist", cmd.Name)
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
