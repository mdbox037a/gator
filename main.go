package main

import (
	"database/sql"
	// "encoding/json"
	// "fmt"
	"log"
	"os"

	"github.com/mdbox037a/gator/internal/config"
	"github.com/mdbox037a/gator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	currentConfig, err := config.Read()
	if err != nil {
		log.Fatalf("Error: failed to read config - %v", err)
	}

	db, err := sql.Open("postgres", currentConfig.DbURL)
	if err != nil {
		log.Fatalf("Error: failed to open database connection at %v", err)
	}
	dbQueries := database.New(db)

	currentState := state{db: dbQueries, currentConfig: &currentConfig}
	commandSet := commands{handlers: make(map[string]func(*state, command) error)}

	commandSet.register("login", handlerLogin)
	commandSet.register("register", handlerRegister)
	commandSet.register("reset", handlerReset)
	commandSet.register("users", handlerUsers)
	commandSet.register("agg", handlerAgg)
	commandSet.register("addfeed", handlerAddFeed)
	commandSet.register("feeds", handlerFeeds)
	commandSet.register("follow", handlerFollow)

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

	// debugConfig, err := json.MarshalIndent(currentConfig, "", "    ")
	// fmt.Printf("Debug: current .gatorconfig contents:\n%s\n", string(debugConfig))
}
