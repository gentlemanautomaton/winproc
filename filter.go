// +build windows

package winproc

import "strings"

// A Filter returns true if it matches a process.
type Filter func(Process) bool

// MatchPID returns a filter that matches a process ID.
func MatchPID(pid int) Filter {
	return func(process Process) bool {
		return process.ID == pid
	}
}

// MatchName returns a filter that matches a process name.
func MatchName(name string) Filter {
	return func(process Process) bool {
		return strings.EqualFold(process.Name, name)
	}
}

// ContainsName returns a filter that matches part of a process name.
func ContainsName(name string) Filter {
	name = strings.ToLower(name)
	return func(process Process) bool {
		return strings.Contains(strings.ToLower(process.Name), name)
	}
}
