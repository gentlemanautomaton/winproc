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

func BenchmarkListWithCommands(b *testing.B) {
	for n := 0; n < b.N; n++ {
		winproc.List(winproc.CollectCommands)
	}
}

func BenchmarkListWithSessions(b *testing.B) {
	for n := 0; n < b.N; n++ {
		winproc.List(winproc.CollectSessions)
	}
}

func BenchmarkListWithUsers(b *testing.B) {
	for n := 0; n < b.N; n++ {
		winproc.List(winproc.CollectUsers)
	}
}

func BenchmarkListWithAll(b *testing.B) {
	for n := 0; n < b.N; n++ {
		winproc.List(winproc.CollectCommands, winproc.CollectSessions, winproc.CollectUsers)
	}
}
