package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mjshaffer117/gator/internal/config"
)

type state struct {
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
	cmds.register("login", handlerLogin)

	// Check for two arguments: First is the program name, second is the command
	if len(os.Args) < 2 {
		log.Fatalf("no command provided")
	}

	cmd := &command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}

	err = cmds.run(&state{cfg: &cfg}, cmd)
	if err != nil {
		log.Fatalf("error running command: %v", err)
	}

	fmt.Printf("updated config file: %+v\n", cfg)
}
