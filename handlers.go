package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mjshaffer117/gator/internal/database"
)

func handlerLogin(s *state, cmd *command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("login command requires a username argument")
	}
	name := cmd.Args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("failed to get user: %v", err)
	}
	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to set user: %v", err)
	}
	fmt.Printf("user has been set: %s\n", cmd.Args[0])
	return nil
}

func handlerRegister(s *state, cmd *command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("register command requires a username argument")
	}
	name := cmd.Args[0]
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	fmt.Printf("user has been registered: %+v\n", user)
	log.Printf("user has been registered: %+v\n", user)
	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to set user: %v", err)
	}
	return nil
}

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
