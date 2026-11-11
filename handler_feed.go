package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/inifomeeo/rssflow/internal/database"
)

func handlerAddFeed(s *state, cmd command) error {
	// Get the current user from the database
    user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
    if err != nil {
        return fmt.Errorf("error getting current user: %w", err)
    }

	if len(cmd.Arguments) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	
	name := cmd.Arguments[0]
	url := cmd.Arguments[1]

	feed, err := s.db.AddFeed(context.Background(), database.AddFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't add RSS feed to the database: %w", err)
	} 

	fmt.Printf("Feed created: %+v\n", feed)
	return nil
}

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("no arguments are required")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("feeds could not be retrieved: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %s\n", feed.Name)
        fmt.Printf("URL: %s\n", feed.Url)
        fmt.Printf("User: %s\n", feed.UserName)
        fmt.Println("--------------------------------")
	}

	return nil
}

func handlerFollowFeed(s *state, cmd command) error {
	// Get the current user from the database
    user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
    if err != nil {
        return fmt.Errorf("error getting current user: %w", err)
    }

	if len(cmd.Arguments) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}

	url := cmd.Arguments[0]

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed follow: %w", err)
	}

	fmt.Printf("%s is now following %s\n",
        feedFollow.UserName,
        feedFollow.FeedName,
    )

	return nil
}

func handlerFollowing(s *state, cmd command) error {
	// Get the current user from the database
    user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
    if err != nil {
        return fmt.Errorf("error getting current user: %w", err)
    }

	if len(cmd.Arguments) != 0 {
		return fmt.Errorf("no arguments are required")
	}

	feedFollows, err := s.db.GetFeedsFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %w", err)
	}

	fmt.Printf("%s is following:\n", user.Name)

	for _, feedFollow := range feedFollows {
		fmt.Printf("  - %s\n", feedFollow.FeedName)
	}

	return nil
}
