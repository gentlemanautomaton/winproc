// +build windows

package winproc

import "strings"

// A StringMatcher is a function that matches strings
type StringMatcher func(string) bool

// A Filter returns true if it matches a process.
type Filter func(Process) bool

// MatchID returns a filter that matches a process ID.
func MatchID(pid ID) Filter {
	return func(process Process) bool {
		return process.ID == pid
	}
}

// MatchName returns a filter that matches a process name.
func MatchName(matcher StringMatcher) Filter {
	return func(process Process) bool {
		return matcher(process.Name)
	}
}

// EqualsName returns a filter that matches a process name case-insensitively.
func EqualsName(name string) Filter {
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
