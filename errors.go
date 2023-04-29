//go:build windows
// +build windows

package winproc

import "errors"

var (
	// ErrClosed is returned when a process handle needed for an action has
	// already been closed.
	ErrClosed = errors.New("the process handle has been closed")
)
