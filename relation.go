//go:build windows
// +build windows

package winproc

// Relation is a collection option that includes processes related to those
// already matched.
type Relation int

const (
	// IncludeAncestors is a collection option that includes all ancestors
	// of matched processes.
	IncludeAncestors Relation = 1 << iota

	// IncludeDescendants is a collection option that includes all descendants
	// of matched processes.
	IncludeDescendants
)

// Contains returns true if r contains b.
func (r Relation) Contains(b Relation) bool {
	return r&b == b
}

// Apply applies the relation to the collection.
func (r Relation) Apply(col *Collection) {
	if r == 0 {
		return
	}

	// Pass 1: Build relationships
	parents := make(map[ID]ID)    // Maps process ID to parent ID
	children := make(map[ID][]ID) // Maps parent ID to child process ID
	for _, process := range col.Procs {
		parents[process.ID] = process.ParentID
		children[process.ParentID] = append(children[process.ParentID], process.ID)
	}

	// Pass 2: Include ancestors
	descendantMatched := make(map[ID]bool) // Processes with a matched descendant
	if r.Contains(IncludeAncestors) {
		for i := range col.Procs {
			if col.Excluded[i] {
				continue
			}
			pid := col.Procs[i].ID
			for {
				if descendantMatched[pid] {
					break // Already processed
				}
				descendantMatched[pid] = true
				parent, ok := parents[pid]
				if !ok || parent == pid {
					break
				}
				pid = parent
			}
		}
	}

	// Pass 3: Include descendants
	ancestorMatched := make(map[ID]bool) // Processes with a matched ancestor
	if r.Contains(IncludeDescendants) {
		for i := range col.Procs {
			if col.Excluded[i] {
				continue
			}
			pid := col.Procs[i].ID
			next, ok := children[pid]
			if !ok {
				continue
			}
			for len(next) > 0 {
				current := next
				next = nil
				for _, child := range current {
					if ancestorMatched[child] {
						continue // Already processed
					}
					ancestorMatched[child] = true
					next = append(next, children[child]...)
				}
			}
		}
	}

	// Pass 4: Include matches
	for i := range col.Procs {
		pid := col.Procs[i].ID
		if ancestorMatched[pid] || descendantMatched[pid] {
			col.Excluded[i] = false
		}
	}
}

// Merge attempts to merge the relation with the next option. It returns true
// if successful.
func (r Relation) Merge(next CollectionOption) (merged CollectionOption, ok bool) {
	n, ok := next.(Relation)
	if !ok {
		return nil, false
	}
	return r | n, true
}
