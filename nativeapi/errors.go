// +build windows

package nativeapi

import "errors"

var (
	// ErrEmptyBuffer is returned when a nil or zero-sized buffer is provided
	// to a system call.
	ErrEmptyBuffer = errors.New("nil or empty buffer provided")
)
