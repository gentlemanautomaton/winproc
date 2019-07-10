// +build windows

package winproc_test

import (
	"fmt"
	"strings"

	"github.com/gentlemanautomaton/winproc"
)

func ExampleTree() {
	procs, err := winproc.List(
		winproc.Include(winproc.ContainsName("svchost")),
		winproc.IncludeAncestors,
		winproc.IncludeDescendants,
		winproc.CollectCommands)
	if err != nil {
		fmt.Printf("Failed to retrieve process list: %v\n", err)
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
