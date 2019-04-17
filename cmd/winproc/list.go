package main

import (
	"context"
	"fmt"

	"github.com/gentlemanautomaton/winproc"
)

func list(ctx context.Context, opts ...winproc.CollectionOption) {
	procs, err := winproc.List(opts...)
	if err != nil {
		fmt.Printf("Failed to retrieve process list: %v\n", err)
		return
	}
	for _, proc := range procs {
		fmt.Printf("%s\n", proc)
	}
}
