package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mjshaffer117/gator/internal/database"
)

func handlerAddFeed(s *state, cmd *command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("addfeed command requires a name and url argument")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUser)
	if err != nil {
		return fmt.Errorf("failed to get current user: %v", err)
	}
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    uuid.NullUUID{UUID: user.ID, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("failed to create feed: %v", err)
	}
	fmt.Printf("feed has been added: %+v\n", feed)
	return nil
}
