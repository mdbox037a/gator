package main

import (
	"fmt"
	"log"

	"github.com/mdbox037a/gator/internal/config"
)

const currentUser string = "matt"

func main() {
	currentConfig, err := config.Read()
	if err != nil {
		log.Fatalf("Error: failed to read config - %v", err)
	}

	err = currentConfig.SetUser(currentUser)
	if err != nil {
		log.Fatalf("Error: failed to set user - %v\n", err)
	}

	currentConfig, err = config.Read()
	if err != nil {
		log.Fatalf("Error: failed to read config - %v", err)
	}
	fmt.Printf("%+v\n", currentConfig)
}
