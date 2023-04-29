//go:build windows
// +build windows

package winproc

import (
	"io"
	"syscall"

	"github.com/gentlemanautomaton/winproc/psapi"
)

// scan collects all processes from the system.
func scan() (procs []Process, err error) {
	// TODO: Use WTSEnumerateProcesses?
	// http://codexpert.ro/blog/2013/12/01/listing-processes-part-4-using-remote-desktop-services-api/
	// https://docs.microsoft.com/en-us/windows/win32/api/wtsapi32/nf-wtsapi32-wtsenumerateprocessesexw

	snapshot, err := psapi.CreateSnapshot(psapi.SnapProcess, 0)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(snapshot)

	var entry psapi.ProcessEntry
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

	return procs, nil
}
