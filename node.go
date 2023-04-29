//go:build windows
// +build windows

package winproc

// Node is a node in a process tree.
type Node struct {
	Process
	Children []Node
}
