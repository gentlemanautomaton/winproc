// +build windows

package winproc

// List returns a list of running processes. Collection options can be
// provided to filter the list and collect additional process information.
// Options will be evaluated in order.
//
// If a filter relies on process information gathered by one or more
// collector options, those options must be included before the filter.
func List(options ...CollectionOption) ([]Process, error) {
	// Collect all processes from the system
	procs, err := scan()
	if err != nil {
		return nil, err
	}

	// Form a collection
	col := Collection{
		Procs:    procs,
		Excluded: make([]bool, len(procs)),
	}

	// Apply each collection option in order
	for i := 0; i < len(options); i++ {
		opt := options[i]

		// Merge adjacent options when possible.
		//
		// This can improve efficiency by combining several options that work
		// more efficiently as a batch.
		for i+1 < len(options) {
			mergable, ok := opt.(MergableCollectionOption)
			if !ok {
				break
			}
			merged, ok := mergable.Merge(options[i+1])
			if !ok {
				break
			}
			opt = merged
			i++
		}

		opt.Apply(&col)
	}

	// Count the number of matches
	total := 0
	for i := range col.Excluded {
		if !col.Excluded[i] {
			total++
		}
	}

	// If all procs matched just return the original slice
	if len(col.Procs) == total {
		return col.Procs, nil
	}

	// Return the matches
	matched := make([]Process, 0, total)
	for i := range col.Procs {
		if col.Excluded[i] {
			continue
		}
		matched = append(matched, col.Procs[i])
	}
	return matched, nil
}
