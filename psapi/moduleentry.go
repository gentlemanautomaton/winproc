// +build windows

package psapi

import "syscall"

// ModuleEntry holds information about a module within a process.
//
// https://docs.microsoft.com/en-us/windows/desktop/api/tlhelp32/ns-tlhelp32-moduleentry32w
type ModuleEntry struct {
	Size        uint32
	ModuleID    uint32
	ProcessID   uint32
	GlobalUsage uint32 // Unused
	ProcUsage   uint32 // Unused
	BaseAddr    uintptr
	BaseSize    uint32
	Handle      syscall.Handle
	NameBuffer  [MaxModuleName + 1]uint16
	PathBuffer  [syscall.MAX_PATH]uint16
}

// Name returns the module name as a string.
func (entry *ModuleEntry) Name() string {
	return syscall.UTF16ToString(entry.NameBuffer[:])
}

// Path returns the module path as a string.
func (entry *ModuleEntry) Path() string {
	return syscall.UTF16ToString(entry.PathBuffer[:])
}
