package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd *command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("login command requires a username argument")
	}
	err := s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to set user: %v", err)
	}
	fmt.Printf("user has been set: %s\n", cmd.Args[0])
	return nil
}
