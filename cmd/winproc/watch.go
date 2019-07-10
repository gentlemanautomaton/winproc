package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gentlemanautomaton/winproc"
)

func watch(ctx context.Context, interval time.Duration, opts ...winproc.CollectionOption) {
	for cs := range winproc.Watch(ctx, interval, 8, opts...) {
		switch cs.Err {
		case context.Canceled, context.DeadlineExceeded, nil:
		default:
			fmt.Printf("Failed to retrieve process tree: %v\n", cs.Err)
		}

		for _, change := range cs.Changes {
			if change.Removed {
				fmt.Printf("STOP: %s\n", change.Process)
			} else {
				fmt.Printf("START: %s\n", change.Process)
			}
		}
	}
}
