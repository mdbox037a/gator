package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mdbox037a/gator/internal/config"
)

const currentUser string = "matt"

func main() {
	currentConfig, err := config.Read()
	if err != nil {
		log.Fatalf("Error: failed to read config - %v", err)
	}
	currentState := state{currentConfig: &currentConfig}
	commandSet := commands{handlers: make(map[string]func(*state, command) error)}

	commandSet.register("login", handlerLogin)
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Error: no command name provided")
	}
	currentCommand := command{
		Name: args[1],
		Args: args[2:],
	}
	err = commandSet.run(&currentState, currentCommand)
	if err != nil {
		log.Fatalf("Error: failed to run command - %v", err)
	}

	fmt.Printf("Debug: resulting program state: %+v\n", currentConfig)
}
