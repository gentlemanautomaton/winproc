//go:build windows
// +build windows

package winproc

// Change describes a process that is being added or removed from a process
// list.
type Change struct {
	Process Process
	Removed bool
}
