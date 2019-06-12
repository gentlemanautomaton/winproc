// +build windows

package winproc_test

import (
	"testing"

	"github.com/gentlemanautomaton/winproc"
)

func BenchmarkTree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		winproc.Tree()
	}
}

func BenchmarkTreeWithCommands(b *testing.B) {
	for n := 0; n < b.N; n++ {
		winproc.Tree(winproc.CollectCommands)
	}
}
