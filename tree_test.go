//go:build windows
// +build windows

package winproc_test

import (
	"testing"

	"github.com/gentlemanautomaton/winproc"
)

func BenchmarkTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		list, err := winproc.List()
		if err != nil {
			b.Error(err)
		}
		winproc.Tree(list)
	}
}

func BenchmarkTreeWithCommands(b *testing.B) {
	for n := 0; n < b.N; n++ {
		list, err := winproc.List(winproc.CollectCommands)
		if err != nil {
			b.Error(err)
		}
		winproc.Tree(list)
	}
}
