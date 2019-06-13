// +build windows

package winproc

import (
	"io"
	"strings"
	"sync"
	"syscall"

	"github.com/gentlemanautomaton/winproc/nativeapi"
	"github.com/gentlemanautomaton/winproc/processaccess"

	"github.com/gentlemanautomaton/cmdline/cmdlinewindows"
	"github.com/gentlemanautomaton/winproc/psapi"
)

// List returns a list of all running processes.
func List(options ...CollectionOption) ([]Process, error) {
	var opts collectionOpts
	for _, option := range options {
		option(&opts)
	}

	snapshot, err := psapi.CreateSnapshot(psapi.SnapProcess, 0)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(snapshot)

	// Step 1: Collect all processes from the system
	var (
		procs []Process
		entry psapi.ProcessEntry
	)
	for entry, err = psapi.FirstProcess(snapshot); err == nil; entry, err = psapi.NextProcess(snapshot) {
		procs = append(procs, Process{
			ID:       ID(entry.ProcessID),
			ParentID: ID(entry.ParentProcessID),
			Name:     entry.Name(),
			Threads:  int(entry.Threads),
		})
	}
	if err != io.EOF {
		return nil, err
	}

	// Step 2: Remove processes that don't match the provided filters
	if len(opts.Include) > 0 {
		procs = applyFilters(procs, opts)
	}

	// Step 3: Collect process commands and/or sessions if requested
	if opts.CollectCommands || opts.CollectSessions || opts.CollectUsers {
		var wg sync.WaitGroup
		wg.Add(len(procs))
		for i := range procs {
			go func(i int) {
				defer wg.Done()

				process, err := syscall.OpenProcess(uint32(processaccess.QueryLimitedInformation), false, uint32(procs[i].ID))
				if err != nil {
					return
				}
				defer syscall.CloseHandle(process)

				if opts.CollectCommands {
					if line, err := nativeapi.ProcessCommandLine(process); err == nil {
						procs[i].CommandLine = strings.TrimSpace(line)
						procs[i].Path, procs[i].Args = cmdlinewindows.SplitCommand(line)
					}
				}

				if opts.CollectSessions {
					if sessionID, err := nativeapi.ProcessSessionID(process); err == nil {
						procs[i].SessionID = sessionID
					}
				}

				if opts.CollectUsers {
					if user, err := userFromProcess(process); err == nil {
						procs[i].User = user
					}
				}
			}(i)
		}
		wg.Wait()
	}

	return procs, nil
}

func applyFilters(procs []Process, opts collectionOpts) (filtered []Process) {
	match := MatchAny(opts.Include...)

	// Fast path for basic filtering that doesn't require relationships
	if !opts.IncludeAncestors && !opts.IncludeDescendants {
		for i := range procs {
			if match(procs[i]) {
				filtered = append(filtered, procs[i])
			}
		}
		return filtered
	}

	// Pass 1: Build relationships and a map of direct matches
	parents := make(map[ID]ID)    // Maps process ID to parent ID
	children := make(map[ID][]ID) // Maps parent ID to child process ID
	matched := make(map[ID]bool)  // Matched processes
	for _, process := range procs {
		parents[process.ID] = process.ParentID
		children[process.ParentID] = append(children[process.ParentID], process.ID)
		if match(process) {
			matched[process.ID] = true
		}
	}

	// Pass 2: Include ancestors
	descendantMatched := make(map[ID]bool) // Processes with a matched descendant
	if opts.IncludeAncestors {
		for i := range procs {
			pid := procs[i].ID
			if !matched[pid] {
				continue
			}
			for {
				if descendantMatched[pid] {
					break // Already processed
				}
				descendantMatched[pid] = true
				parent, ok := parents[pid]
				if !ok || parent == pid {
					break
				}
				pid = parent
			}
		}
	}

	// Pass 3: Include descendants
	ancestorMatched := make(map[ID]bool) // Processes with a matched ancestor
	if opts.IncludeDescendants {
		for i := range procs {
			pid := procs[i].ID
			if !matched[pid] {
				continue
			}
			next, ok := children[pid]
			if !ok {
				continue
			}
			for len(next) > 0 {
				current := next
				next = nil
				for _, child := range current {
					if ancestorMatched[child] {
						continue // Already processed
					}
					ancestorMatched[child] = true
					next = append(next, children[child]...)
				}
			}
		}
	}

	// Pass 4: Union of all matches
	for i := range procs {
		pid := procs[i].ID
		if matched[pid] || ancestorMatched[pid] || descendantMatched[pid] {
			filtered = append(filtered, procs[i])
		}
	}

	return filtered
}
