// +build windows

package winproc

import (
	"sync"
	"syscall"

	"github.com/gentlemanautomaton/winproc/nativeapi"
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

// CommandLine returns the command line used to invoke the process.
//
// This call is only supported on Windows 10 1511 or newer.
//
// TODO: Consider adding support for older operating systems by using older,
// more arcane implementations when necessary.
//
// https://wj32.org/wp/2009/01/24/howto-get-the-command-line-of-processes/
// https://stackoverflow.com/questions/45891035/get-processid-from-binary-path-command-line-statement-in-c
func (ref *Ref) CommandLine() (command string, err error) {

	ref.mutex.RLock()
	defer ref.mutex.RUnlock()

	if ref.handle == syscall.InvalidHandle {
		return "", ErrClosed
	}

	return nativeapi.ProcessCommandLine(ref.handle)
}

// SessionID returns the ID of the windows session associated with the
// process.
func (ref *Ref) SessionID() (sessionID uint32, err error) {
	ref.mutex.RLock()
	defer ref.mutex.RUnlock()

	if ref.handle == syscall.InvalidHandle {
		return 0, ErrClosed
	}

	return nativeapi.ProcessSessionID(ref.handle)
}

// User returns information about the windows user associated with the
// process.
func (ref *Ref) User() (user User, err error) {
	ref.mutex.RLock()
	defer ref.mutex.RUnlock()

	if ref.handle == syscall.InvalidHandle {
		return User{}, ErrClosed
	}

	return userFromProcess(ref.handle)
}

// Times returns time information about the process.
func (ref *Ref) Times() (times Times, err error) {
	ref.mutex.RLock()
	defer ref.mutex.RUnlock()

	if ref.handle == syscall.InvalidHandle {
		return Times{}, ErrClosed
	}

	var creation, exit, kernel, user syscall.Filetime
	if err := syscall.GetProcessTimes(ref.handle, &creation, &exit, &kernel, &user); err != nil {
		return Times{}, err
	}

	return Times{
		Creation: timeFromFiletime(creation),
		Exit:     timeFromFiletime(exit),
		Kernel:   durationFromFiletime(kernel),
		User:     durationFromFiletime(user),
	}, nil
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
