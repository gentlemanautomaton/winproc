//go:build windows
// +build windows

package winproc

import "strings"

// A StringMatcher is a function that matches strings
type StringMatcher func(string) bool

// MatchAny returns true if any of the filters match the process.
//
// MatchAny returns true if no filters are provided.
func MatchAny(filters ...Filter) Filter {
	return func(process Process) bool {
		if len(filters) == 0 {
			return true
		}
		for _, filter := range filters {
			if filter(process) {
				return true
			}
		}
		return false
	}
}

// MatchAll returns true if all of the filters match the process.
//
// MatchAll returns true if no filters are provided.
func MatchAll(filters ...Filter) Filter {
	return func(process Process) bool {
		for _, filter := range filters {
			if !filter(process) {
				return false
			}
		}
		return true
	}
}

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
