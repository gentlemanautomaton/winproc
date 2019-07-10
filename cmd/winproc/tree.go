package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/gentlemanautomaton/winproc"
)

func tree(ctx context.Context, opts ...winproc.CollectionOption) {
	procs, err := winproc.List(opts...)
	if err != nil {
		fmt.Printf("Failed to retrieve process tree: %v\n", err)
		return
	}
	printChildren(0, winproc.Tree(procs))
}

func printChildren(depth int, nodes []winproc.Node) {
	for _, node := range nodes {
		fmt.Printf("%s%s\n", strings.Repeat("  ", depth), node.Process)
		printChildren(depth+1, node.Children)
	}
}
