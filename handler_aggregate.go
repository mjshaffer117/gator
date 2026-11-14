package main

import (
	"context"
	"fmt"
)

func aggregateHandler(s *state, cmd *command) error {
	res, err := fetchFeed(context.Background(), "https://wagslane.dev/index.xml")
	if err != nil {
		fmt.Printf("Error fetching feed: %v\n", err)
		return err
	}
	fmt.Printf("Feed results: %+v\n", res)
	return nil
}
