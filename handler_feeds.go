package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd *command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("feeds command does not accept arguments")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds: %w", err)
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID.UUID)
		if err != nil {
			return fmt.Errorf("failed to get user for feed %s: %w", feed.Name, err)
		}
		fmt.Printf("* Feed Name: %s - Feed URL: %s - Username: %s\n", feed.Name, feed.Url, user.Name)
	}
	return nil
}
