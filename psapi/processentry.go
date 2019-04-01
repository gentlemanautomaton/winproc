// +build windows

package psapi

import "syscall"

// ProcessEntry holds information about a process.
//
// https://docs.microsoft.com/en-us/windows/desktop/api/tlhelp32/ns-tlhelp32-processentry32w
type ProcessEntry struct {
	Size               uint32
	Usage              uint32 // Unused
	ProcessID          uint32
	DefaultHeapID      uintptr // Unused
	ModuleID           uint32  // Unused
	Threads            uint32
	ParentProcessID    uint32
	BaseThreadPriority int32
	Flags              uint32
	NameBuffer         [syscall.MAX_PATH]uint16
}

// Name returns the executable file name of the process as a string.
func (entry *ProcessEntry) Name() string {
	return syscall.UTF16ToString(entry.NameBuffer[:])
}
