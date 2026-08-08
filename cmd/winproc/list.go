package main

import (
	"context"
	"fmt"

	"github.com/gentlemanautomaton/winproc"
)

// ListCmd provides a list view of the windows process list.
type ListCmd struct {
	IncludePIDs        []uint32 `kong:"optional,name='pid',help='Include processes with a particular ID.'"`
	IncludeNames       []string `kong:"optional,name='name',help='Include processes with a particular name.'"`
	IncludeAncestors   bool     `kong:"optional,name='ancestors',short='a',help='Include ancestors of matching processes.'"`
	IncludeDescendents bool     `kong:"optional,name='descendents',short='d',help='Include descendants of matching processes.'"`
}

// Run executes the list command.
func (cmd ListCmd) Run(ctx context.Context) error {
	opts := makeOptions(cmd.IncludePIDs, cmd.IncludeNames, cmd.IncludeAncestors, cmd.IncludeDescendents)
	procs, err := winproc.List(opts...)
	if err != nil {
		return fmt.Errorf("failed to retrieve process list: %v\n", err)
	}
	for _, proc := range procs {
		fmt.Printf("%s\n", proc)
	}
	return nil
}
