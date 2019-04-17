// +build windows

package winproc

// MatchAny returns true if any of the filters match the process.
//
// MatchAny returns true if no filters are provided.
func MatchAny(process Process, filters ...Filter) bool {
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

// MatchAll returns true if all of the filters match the process.
//
// MatchAll returns true if no filters are provided.
func MatchAll(process Process, filters ...Filter) bool {
	for _, filter := range filters {
		if !filter(process) {
			return false
		}
	}
	return true
}
