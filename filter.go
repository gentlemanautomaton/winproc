//go:build windows
// +build windows

package winproc

// A Filter returns true if it matches a process.
type Filter func(Process) bool

// Include is an inclusion filter.
type Include Filter

// Apply applies the inclusion filter to the collection.
func (include Include) Apply(col *Collection) {
	if include == nil {
		return
	}

	for i := range col.Procs {
		if !col.Excluded[i] {
			col.Excluded[i] = !include(col.Procs[i])
		}
	}
}

// Exclude is an exclusion filter.
type Exclude Filter

// Apply applies the exclusion filter to the collection.
func (exclude Exclude) Apply(col *Collection) {
	if exclude == nil {
		return
	}

	for i := range col.Procs {
		if !col.Excluded[i] {
			col.Excluded[i] = exclude(col.Procs[i])
		}
	}
}
