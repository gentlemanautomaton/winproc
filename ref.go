// +build windows

package winproc

import (
	"sync"
	"syscall"

	"github.com/gentlemanautomaton/winproc/processaccess"
)

// A Ref is a reference to a running or exited process. It manages an open
// system handle internally. As long as a reference is open it prevents the
// associated process ID from being recycled.
//
// Each reference must be closed when it is no longer needed.
type Ref struct {
	mutex  sync.RWMutex
	handle syscall.Handle
}

// Open returns a reference to the process with the given process ID and
// access rights.
//
// If one or more access rights are provided, they will be combined. If no
// access rights are provided, the QueryLimitedInformation right will be used.
//
// It is the caller's responsibility to close the handle when finished with it.
func Open(pid ID, rights ...processaccess.Rights) (*Ref, error) {
	var combinedRights processaccess.Rights
	if len(rights) > 0 {
		for _, r := range rights {
			combinedRights |= r
		}
	} else {
		combinedRights = processaccess.QueryLimitedInformation
	}

	handle, err := syscall.OpenProcess(uint32(combinedRights), false, uint32(pid))
	if err != nil {
		return nil, err
	}
	return &Ref{handle: handle}, nil
}

// Close releases the process handle maintained by ref.
//
// If ref has already been closed it will return ErrClosed.
func (ref *Ref) Close() error {
	ref.mutex.Lock()
	defer ref.mutex.Unlock()

	if ref.handle == syscall.InvalidHandle {
		return ErrClosed
	}

	if err := syscall.CloseHandle(ref.handle); err != nil {
		return err
	}
	ref.handle = syscall.InvalidHandle

	return nil
}
