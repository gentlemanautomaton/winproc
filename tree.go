// +build windows

package winproc

// Tree returns a hierarchy of processes.
func Tree(options ...CollectionOption) ([]Node, error) {
	procs, err := List(options...)
	if err != nil {
		return nil, err
	}

	nodes := make(map[int]Process, len(procs))
	hierarchy := make(map[int][]int, len(procs))
	for _, proc := range procs {
		nodes[proc.ID] = proc
		hierarchy[proc.ParentID] = append(hierarchy[proc.ParentID], proc.ID)
	}

	var roots []Node
	for _, proc := range procs {
		if _, hasParent := nodes[proc.ParentID]; hasParent {
			continue
		}
		roots = append(roots, Node{
			Process:  proc,
			Children: childNodes(proc.ID, nodes, hierarchy),
		})
	}
	return roots, nil
}

func childNodes(parent int, nodes map[int]Process, hierarchy map[int][]int) []Node {
	childIDs := hierarchy[parent]
	if len(childIDs) == 0 {
		return nil
	}
	children := make([]Node, 0, len(childIDs))
	for _, child := range childIDs {
		children = append(children, Node{
			Process:  nodes[child],
			Children: childNodes(child, nodes, hierarchy),
		})
	}
	return children
}
