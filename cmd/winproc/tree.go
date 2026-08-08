package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/gentlemanautomaton/winproc"
)

// TreeCmd provides a tree view of the windows process list.
type TreeCmd struct {
	IncludePIDs        []uint32 `kong:"optional,name='pid',help='Include processes with a particular ID.'"`
	IncludeNames       []string `kong:"optional,name='name',help='Include processes with a particular name.'"`
	IncludeAncestors   bool     `kong:"optional,name='ancestors',short='a',help='Include ancestors of matching processes.'"`
	IncludeDescendents bool     `kong:"optional,name='descendents',short='d',help='Include descendants of matching processes.'"`
}

// Run executes the tree command.
func (cmd TreeCmd) Run(ctx context.Context) error {
	opts := makeOptions(cmd.IncludePIDs, cmd.IncludeNames, cmd.IncludeAncestors, cmd.IncludeDescendents)
	procs, err := winproc.List(opts...)
	if err != nil {
		return fmt.Errorf("failed to retrieve process tree: %v\n", err)
	}
	printChildren(0, winproc.Tree(procs))
	return nil
}

func printChildren(depth int, nodes []winproc.Node) {
	for _, node := range nodes {
		fmt.Printf("%s%s\n", strings.Repeat("  ", depth), node.Process)
		printChildren(depth+1, node.Children)
	}
}
