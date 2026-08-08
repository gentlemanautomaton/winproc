package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gentlemanautomaton/winproc"
)

// WatchCmd watches the windows process list.
type WatchCmd struct {
	IncludePIDs        []uint32      `kong:"optional,name='pid',help='Include processes with a particular ID.'"`
	IncludeNames       []string      `kong:"optional,name='name',help='Include processes with a particular name.'"`
	IncludeAncestors   bool          `kong:"optional,name='ancestors',short='a',help='Include ancestors of matching processes.'"`
	IncludeDescendents bool          `kong:"optional,name='descendents',short='d',help='Include descendants of matching processes.'"`
	Interval           time.Duration `kong:"optional,name='interval',short='i',default='1s',help='Interval between updates.'"`
}

// Run executes the watch command.
func (cmd WatchCmd) Run(ctx context.Context) error {
	opts := makeOptions(cmd.IncludePIDs, cmd.IncludeNames, cmd.IncludeAncestors, cmd.IncludeDescendents)
	for cs := range winproc.Watch(ctx, cmd.Interval, 8, opts...) {
		if cs.Err != nil {
			switch cs.Err {
			case context.Canceled, context.DeadlineExceeded:
				return nil
			default:
				return fmt.Errorf("failed to retrieve process tree: %v\n", cs.Err)
			}
		}

		for _, change := range cs.Changes {
			if change.Removed {
				fmt.Printf("STOP: %s\n", change.Process)
			} else {
				fmt.Printf("START: %s\n", change.Process)
			}
		}
	}
	return nil
}
