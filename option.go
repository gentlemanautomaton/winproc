// +build windows

package winproc

type collectionOpts struct {
	Include            []Filter
	IncludeAncestors   bool
	IncludeDescendants bool
	CollectCommands    bool
	CollectSessions    bool
	CollectUsers       bool
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

// CollectCommands is an option that enables collection of process
// command line, path and argument information.
func CollectCommands(opts *collectionOpts) {
	opts.CollectCommands = true
}

// CollectSessions is an option that enables collection of process
// session information.
func CollectSessions(opts *collectionOpts) {
	opts.CollectSessions = true
}

// CollectUsers is an option that enables collection of process
// user information.
func CollectUsers(opts *collectionOpts) {
	opts.CollectUsers = true
}
