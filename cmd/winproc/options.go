package main

import (
	"github.com/gentlemanautomaton/winproc"
)

func makeOptions(pids []uint32, names []string, ancestors, descendants bool) (opts []winproc.CollectionOption) {
	for _, pid := range pids {
		opts = append(opts, winproc.Include(winproc.MatchID(winproc.ID(pid))))
	}

	for _, name := range names {
		opts = append(opts, winproc.Include(winproc.ContainsName(name)))
	}

	if ancestors {
		opts = append(opts, winproc.IncludeAncestors)
	}

	if descendants {
		opts = append(opts, winproc.IncludeDescendants)
	}

	opts = append(opts, winproc.CollectCommands, winproc.CollectSessions)

	return
}
