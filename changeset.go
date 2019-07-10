// +build windows

package winproc

import "time"

// ChangeSet holds a set of changes to a process list.
type ChangeSet struct {
	Changes []Change
	Time    time.Time
	Err     error
}
