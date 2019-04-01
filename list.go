// +build windows

package winproc

import (
	"io"
	"sync"
	"syscall"

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

	var (
		procs []Process
		entry psapi.ProcessEntry
	)
	for entry, err = psapi.FirstProcess(snapshot); err == nil; entry, err = psapi.NextProcess(snapshot) {
		procs = append(procs, Process{
			ID:       int(entry.ProcessID),
			ParentID: int(entry.ParentProcessID),
			Name:     entry.Name(),
		})
	}
	if err != io.EOF {
		return nil, err
	}

	if opts.CollectPaths {
		var wg sync.WaitGroup
		wg.Add(len(procs))
		for i := range procs {
			go func(i int) {
				defer wg.Done()
				if procs[i].ID == 0 {
					// PID 0 is the system process and its path can't be collected
					return
				}
				procs[i].Path = collectPath(uint32(procs[i].ID))
			}(i)
		}
		wg.Wait()
	}

	return procs, nil
}

func collectPath(processID uint32) string {
	snapshot, err := psapi.CreateSnapshot(psapi.SnapModule|psapi.SnapModule32, processID)
	if err != nil {
		return ""
	}
	defer syscall.CloseHandle(snapshot)

	entry, err := psapi.FirstModule(snapshot)
	if err != nil {
		return ""
	}
	return entry.Path()
}
