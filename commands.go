package main

import (
	"github.com/mdbox037a/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}
