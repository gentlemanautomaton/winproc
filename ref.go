//go:build windows
// +build windows

package winproc

import (
	"context"
	"fmt"
	"os"
	"sync"
	"syscall"

	"github.com/gentlemanautomaton/winproc/nativeapi"
	"github.com/gentlemanautomaton/winproc/processaccess"
	"github.com/gentlemanautomaton/winproc/procthreadapi"
	"golang.org/x/sys/windows"
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

// ID returns the ID of the process.
func (ref *Ref) ID() (ID, error) {
	ref.mutex.RLock()
	defer ref.mutex.RUnlock()

	if ref.handle == syscall.InvalidHandle {
		return 0, ErrClosed
	}

	id, err := windows.GetProcessId(windows.Handle(ref.handle))
	if err != nil {
		return 0, err
	}
	return ID(id), nil
}

// UniqueID returns a unique identifier for the process by combining its
// creation time and process ID.
func (ref *Ref) UniqueID() (UniqueID, error) {
	ref.mutex.RLock()
	defer ref.mutex.RUnlock()

	if ref.handle == syscall.InvalidHandle {
		return UniqueID{}, ErrClosed
	}

	id, err := windows.GetProcessId(windows.Handle(ref.handle))
	if err != nil {
		return UniqueID{}, err
	}

	var creation, exit, kernel, user syscall.Filetime
	if err := syscall.GetProcessTimes(ref.handle, &creation, &exit, &kernel, &user); err != nil {
		return UniqueID{}, err
	}

	return UniqueID{
		Creation: creation.Nanoseconds(),
		ID:       ID(id),
	}, nil
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

// Critical returns true if the process is considered critical to the system's
// operation.
//
// This call is only supported on Windows 8.1 or newer.
func (ref *Ref) Critical() (bool, error) {
	ref.mutex.RLock()
	defer ref.mutex.RUnlock()

	if ref.handle == syscall.InvalidHandle {
		return false, ErrClosed
	}

	return procthreadapi.IsProcessCritical(ref.handle)
}

// Wait waits until the process terminates or ctx is cancelled. It
// returns nil if the process has terminated.
//
// If ctx is cancelled the value of ctx.Err() will be returned.
func (ref *Ref) Wait(ctx context.Context) error {
	// Hold a read lock for the duration of the call. This will block
	// calls to ref.Close() until we're done.
	ref.mutex.RLock()
	defer ref.mutex.RUnlock()

	// Exit if the context has already been cancelled
	if err := ctx.Err(); err != nil {
		return err
	}

	// Make sure the process hasn't been closed already
	if ref.handle == syscall.InvalidHandle {
		return ErrClosed
	}

	// We need to wait on the process handle and the context simultaneously,
	// so we'll prepare a windows event that will be signaled when ctx is
	// cancelled and rely on WaitForMultipleObjects to wake up when something
	// happens.
	//
	// The cancellation event will be signaled in a separate goroutine.
	cancelled, err := windows.CreateEvent(nil, 1, 0, nil) // 2nd and 3rd arguments are booleans
	if err != nil {
		return err
	}
	defer windows.CloseHandle(cancelled)

	// Derive a context that we'll cancel via defer to make sure the
	// goroutine always exits
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()

	// Use a WaitGroup to know when the goroutine is done
	var wg sync.WaitGroup
	wg.Add(1)

	// Create a goroutine that will signal the event when ctx is cancelled
	go func() {
		defer wg.Done()
		<-ctx.Done()
		windows.SetEvent(cancelled)
	}()

	// Wait for process termination or context cancellation
	result, err := windows.WaitForMultipleObjects([]windows.Handle{
		cancelled,                  // WAIT_OBJECT_0
		windows.Handle(ref.handle), // WAIT_OBJECT_1
	}, false, windows.INFINITE)

	// Tell the goroutine to exit (in case it hasn't already)
	cancel()

	// Wait for the goroutine to exit (in case it hasn't already)
	wg.Wait()

	switch result {
	case syscall.WAIT_OBJECT_0:
		// The context was cancelled
		return ctx.Err()
	case syscall.WAIT_OBJECT_0 + 1:
		// The process has terminated
		return nil
	case syscall.WAIT_FAILED:
		return os.NewSyscallError("WaitForMultipleObjects", err)
	default:
		return fmt.Errorf("winproc.Ref: unexpected result from WaitForMultipleObjects: %d", result)
	}
}

// Terminate instructs the operating system to terminate the process with the
// given exit code.
func (ref *Ref) Terminate(exitCode uint32) error {
	return procthreadapi.TerminateProcess(ref.handle, exitCode)
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
