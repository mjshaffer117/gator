package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd *command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("reset command does not accept arguments")
	}
	if err := s.db.DeleteAll(context.Background()); err != nil {
		return fmt.Errorf("failed to delete users: %v", err)
	}
	fmt.Println("all users have been deleted")
	return nil
}
