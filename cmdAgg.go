package main

import (
	rss "github.com/can-ek/gator/rss"
	"fmt"
	"context"
)

func handleAgg(s *state, cmd command) error {
	feed, err :=  rss.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
