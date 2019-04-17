// +build windows

package winproc

type collectionOpts struct {
	Include            []Filter
	IncludeAncestors   bool
	IncludeDescendants bool
	CollectPaths       bool
}

// CollectionOption is a process collection option.
type CollectionOption func(*collectionOpts)

// Include is an option that causes a process list to include processes
// matching a filter.
func Include(filter Filter) CollectionOption {
	return func(opts *collectionOpts) {
		opts.Include = append(opts.Include, filter)
	}
}

// IncludeAncestors is an option that includes all ancestors of
// matched processes.
func IncludeAncestors(opts *collectionOpts) {
	opts.IncludeAncestors = true
}

// IncludeDescendants is an option that includes all descendants of
// matched processes.
func IncludeDescendants(opts *collectionOpts) {
	opts.IncludeDescendants = true
}

// CollectPaths is an option that enables process path collection.
func CollectPaths(opts *collectionOpts) {
	opts.CollectPaths = true
}
