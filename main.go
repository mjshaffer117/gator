package main

import (
	"fmt"

	cfg "main.go/internal/config"
)

func main() {
	conf, err := cfg.Read()
	if err != nil {
		panic(err)
	}

	if err := conf.SetUser("Michael"); err != nil {
		panic(err)
	}

	conf, err = cfg.Read()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", conf)
}