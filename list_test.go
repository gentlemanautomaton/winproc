// +build windows

package winproc_test

import (
	"testing"

	"github.com/gentlemanautomaton/winproc"
)

func BenchmarkList(b *testing.B) {
	for n := 0; n < b.N; n++ {
		winproc.List()
	}
}

func BenchmarkListWithPaths(b *testing.B) {
	for n := 0; n < b.N; n++ {
		winproc.List(winproc.CollectPaths)
	}
}