package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // Importing for side effects ( _ ), package is not refernced directly
	"github.com/mjshaffer117/gator/internal/config"
	"github.com/mjshaffer117/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("read config: %+v\n", cfg)
	// Output with field names and values, e.g. {DatabaseURL:postgres://localhost:5432/mydb CurrentUser:Michael}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	cmds := commands{
		make(map[string]func(*state, *command) error),
	}
	cmds.register("register", handlerRegister)
	cmds.register("login", handlerLogin)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	// Check for two arguments: First is the program name, second is the command
	if len(os.Args) < 2 {
		log.Fatalf("no command provided")
	}

	cmd := &command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	db, err := sql.Open("postgres", cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	err = cmds.run(&state{db: dbQueries, cfg: &cfg}, cmd)
	if err != nil {
		log.Fatalf("error running command: %v", err)
	}

	fmt.Printf("updated config file: %+v\n", cfg)
}
