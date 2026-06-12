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
		log.Fatal(err)
	}

	err = currentConfig.SetUser(currentUser)
	currentConfig, err = config.Read()
	fmt.Printf("%+v\n", currentConfig)
}
