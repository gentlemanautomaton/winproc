package main

import (
	"github.com/gentlemanautomaton/winproc"
)

func makeOptions(pids []uint32, names []string, ancestors, descendants bool) (opts []winproc.CollectionOption) {
	var filters []winproc.Filter

	for _, pid := range pids {
		filters = append(filters, winproc.MatchID(winproc.ID(pid)))
	}

	for _, name := range names {
		filters = append(filters, winproc.ContainsName(name))
	}

	if len(filters) > 0 {
		opts = append(opts, winproc.Include(winproc.MatchAny(filters...)))
	}

	if ancestors {
		opts = append(opts, winproc.IncludeAncestors)
	}

	if descendants {
		opts = append(opts, winproc.IncludeDescendants)
	}

	opts = append(opts, winproc.CollectCommands, winproc.CollectSessions, winproc.CollectUsers, winproc.CollectTimes)

	return
}
