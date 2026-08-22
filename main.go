package main

import (
	"fmt"
	"log"

	"github.com/mjshaffer117/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	// Output with field names and values, e.g. {DatabaseURL:postgres://localhost:5432/mydb CurrentUser:Michael}

	err = cfg.SetUser("Michael")
	if err != nil {
		log.Fatalf("Error setting user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	fmt.Printf("Updated config file: %+v\n", cfg)
}
