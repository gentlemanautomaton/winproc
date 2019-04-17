// +build windows

package winproc

// TreeAction is a function that takes action on a tree.
type TreeAction func(tree []Node) error
