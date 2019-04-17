package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gentlemanautomaton/winproc"
)

func watch(ctx context.Context, interval time.Duration, opts ...winproc.CollectionOption) {
	err := winproc.Watch(ctx, func(tree []winproc.Node) error {
		printChildren(0, tree)
		return nil
	}, interval, opts...)

	switch err {
	case context.Canceled, context.DeadlineExceeded:
	default:
		fmt.Printf("Failed to retrieve process tree: %v\n", err)
	}
}
