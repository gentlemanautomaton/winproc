// +build windows

package winproc

type collectionOpts struct {
	CollectPaths bool
}

// CollectionOption is a process collection option.
type CollectionOption func(*collectionOpts)

// CollectPaths is an option that enables process path collection.
func CollectPaths(opts *collectionOpts) {
	opts.CollectPaths = true
}
