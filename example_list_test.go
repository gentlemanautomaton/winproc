// +build windows

package winproc_test

import (
	"fmt"

	"github.com/gentlemanautomaton/winproc"
)

func ExampleList() {
	procs, err := winproc.List(
		winproc.Include(winproc.ContainsName("winlogon")),
		winproc.IncludeAncestors)
	if err != nil {
		fmt.Printf("Failed to retrieve process list: %v\n", err)
		return
	}

	for _, proc := range procs {
		fmt.Printf("%s\n", proc)
	}
}
