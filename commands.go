package main

import (
	"github.com/mdbox037a/gator/internal/config"
)

type state struct {
	CurrentConfig *config.Config
}

type command struct {
	Name string
	Args []string
}
