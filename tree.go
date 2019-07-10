// +build windows

package winproc

// Tree creates a hierarchy out of a list of processes.
func Tree(procs []Process) []Node {
	// Build a lookup for each node and a map from parents to children
	nodes := make(map[ID]Process, len(procs))
	hierarchy := make(map[ID][]ID, len(procs))
	for _, proc := range procs {
		nodes[proc.ID] = proc
		if proc.ID != proc.ParentID {
			hierarchy[proc.ParentID] = append(hierarchy[proc.ParentID], proc.ID)
		}
	}

	// Build a tree from the roots
	var roots []Node
	for _, proc := range procs {
		if proc.ID != findRoot(proc.ID, proc.ParentID, nodes) {
			continue
		}
		roots = append(roots, Node{
			Process:  proc,
			Children: childNodes(proc.ID, proc.ID, nodes, hierarchy),
		})
	}
	return roots
}

func findRoot(pid, parent ID, nodes map[ID]Process) ID {
	// Keep track of the PIDs we've seen
	seen := make([]ID, 0, 8)

	for {
		// Processes that identify themselves as their own parent are roots
		if pid == parent {
			return pid
		}

		// Processes with an inaccessible parent are roots
		node, found := nodes[parent]
		if !found {
			return pid
		}

		// Add this PID to the list
		seen = append(seen, pid)

		// Advance
		pid, parent = parent, node.ParentID

		// If we encounter a PID more than once it means we're working on a
		// circular hierarchy of process IDs. This can happen if a parent
		// process dies and its process ID is recycled by a descendant.
		if lowest, repeated := isRepeat(pid, seen); repeated {
			return lowest // Arbitrarily (but deterministically) selected root
		}
	}
}

// isRepeat checks to see if pid is present in seen. The contents of seen
// is assumed to be in order of ancestry. It returns the lowest PID within
// the cyclic portion of the ancestry.
func isRepeat(pid ID, seen []ID) (lowest ID, repeated bool) {
	const maxID = ^ID(0)
	lowest = maxID
	for i := len(seen) - 1; i >= 0; i-- {
		if seen[i] < lowest {
			lowest = seen[i]
		}
		if seen[i] == pid {
			return lowest, true
		}
	}
	return lowest, false
}

func childNodes(root, parent ID, nodes map[ID]Process, hierarchy map[ID][]ID) []Node {
	childIDs := hierarchy[parent]
	if len(childIDs) == 0 {
		return nil
	}
	children := make([]Node, 0, len(childIDs))
	for _, child := range childIDs {
		if child == root {
			continue // Avoid infinite recursion back to the root
		}
		children = append(children, Node{
			Process:  nodes[child],
			Children: childNodes(root, child, nodes, hierarchy),
		})
	}
	return children
}
