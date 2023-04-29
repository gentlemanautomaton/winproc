//go:build windows
// +build windows

package winproc

import (
	"strings"
	"sync"

	"github.com/gentlemanautomaton/cmdline/cmdlinewindows"
)

// A Collector is a collection option that collects additional information
// about a process.
//
// Information will only be collected for processes that have not been
// excluded by previous filtering options.
type Collector int

const (
	// CollectCommands is an option that enables collection of process
	// command line, path and argument information.
	CollectCommands Collector = 1 << iota

	// CollectSessions is an option that enables collection of process
	// session information.
	CollectSessions

	// CollectUsers is an option that enables collection of process
	// user information.
	CollectUsers

	// CollectTimes is an option that enables collection of process
	// time information.
	CollectTimes

	// CollectCriticality is an option that enables collection of process
	// criticality information.
	CollectCriticality
)

// Contains returns true if c contains b.
func (c Collector) Contains(b Collector) bool {
	return c&b == b
}

// Apply applies the collector to the collection.
func (c Collector) Apply(col *Collection) {
	if c == 0 {
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(col.Procs))

	for i := range col.Procs {
		if col.Excluded[i] {
			wg.Done()
			continue
		}
		go func(i int) {
			defer wg.Done()
			proc := &col.Procs[i]

			ref, err := Open(proc.ID)
			if err != nil {
				return
			}
			defer ref.Close()

			if c.Contains(CollectCommands) {
				if line, err := ref.CommandLine(); err == nil {
					proc.CommandLine = strings.TrimSpace(line)
					proc.Path, proc.Args = cmdlinewindows.SplitCommand(line)
				}
			}

			if c.Contains(CollectSessions) {
				if sessionID, err := ref.SessionID(); err == nil {
					proc.SessionID = sessionID
				}
			}

			if c.Contains(CollectUsers) {
				if user, err := ref.User(); err == nil {
					proc.User = user
				}
			}

			if c.Contains(CollectTimes) {
				if times, err := ref.Times(); err == nil {
					proc.Times = times
				}
			}

			if c.Contains(CollectCriticality) {
				if critical, err := ref.Critical(); err == nil {
					proc.Critical = critical
				}
			}
		}(i)
	}

	wg.Wait()
}

// Merge attempts to merge the collector with the next option. It returns true
// if successful.
func (c Collector) Merge(next CollectionOption) (merged CollectionOption, ok bool) {
	n, ok := next.(Collector)
	if !ok {
		return nil, false
	}
	return c | n, true
}
