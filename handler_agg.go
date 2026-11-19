package main

import (
	"fmt"
	"time"
)

func handlerAggregate(s *state, cmd command) error {
	if len(cmd.Arguments) != 1 {
        return fmt.Errorf("usage: %v <time_between_reqs>", cmd.Name)
    }

	time_between_reqs := cmd.Arguments[0]

	timeBetweenRequests, err := time.ParseDuration(time_between_reqs)
	if err != nil {
        return fmt.Errorf("error parsing duration string: %w", err)
    }

    if timeBetweenRequests <= 0 {
        return fmt.Errorf("duration must be greater than zero")
    }

    fmt.Printf("Collecting feeds every %s\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	defer ticker.Stop()

	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
            fmt.Printf("error scraping feeds: %v\n", err)
        }
	}
}
