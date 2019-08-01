// +build windows

package winproc

import (
	"fmt"
	"strings"
)

// Process holds information about a windows process.
type Process struct {
	ID          ID
	ParentID    ID
	Name        string
	Path        string
	Args        []string
	CommandLine string
	SessionID   uint32
	User        User
	Threads     int
	Times       Times
}

// UniqueID returns a unique identifier for the process by combining its
// creation time and process ID.
//
// The process must contain a creation time for this function to return a
// unique value. This can be accomplished by supplying the CollectTimes
// option when collecting processes.
func (p Process) UniqueID() UniqueID {
	return UniqueID{
		Creation: p.Times.Creation.UnixNano(),
		ID:       p.ID,
	}
}

// Protected returns true if p represents a protected process of some kind:
//
//	All processes in session 0
//	All processes running as Local System, NT Authority or Network Service
//	All processes for which the SID has not been collected
//	The process with ID 0
//	The process with ID 4
//
// TODO: Skip anything with the CREATE_PROTECTED_PROCESS flag?
func (p Process) Protected() bool {
	// https://brianbondy.com/blog/100/understanding-windows-at-a-deeper-level-sessions-window-stations-and-desktops

	// Anything in session zero is a system process
	if p.SessionID == 0 {
		return true
	}

	// The System Idle Process
	if p.ID == 0 {
		return true
	}

	// The System Process
	if p.ID == 4 {
		return true
	}

	// All unidentified processes and processes running with system security
	// identifiers.
	if p.User.SID == "" || p.User.System() {
		return true
	}

	// The Evolution of Protected Processes: http://www.alex-ionescu.com/?p=97

	return false
}

// String returns a string representation of the process.
func (p Process) String() string {
	value := fmt.Sprintf("[%d] PID %d", p.SessionID, p.ID)
	if user := p.User.String(); user != "" {
		value = fmt.Sprintf("%s (%s)", value, user)
	}
	if !p.Times.Creation.IsZero() {
		if p.Times.Exit.IsZero() {
			value = fmt.Sprintf("%s (created %s)", value, p.Times.Creation)
		} else {
			value = fmt.Sprintf("%s (created %s, exited %s)", value, p.Times.Creation, p.Times.Exit)
		}
	}
	if p.Times.Kernel != 0 || p.Times.User != 0 {
		value = fmt.Sprintf("%s (%s user %s kernel)", value, p.Times.Kernel, p.Times.User)
	}
	switch {
	case p.CommandLine != "":
		value = fmt.Sprintf("%s: %s", value, p.CommandLine)
	case p.Path != "" && len(p.Args) > 0:
		value = fmt.Sprintf("%s: %s %s", value, p.Path, strings.Join(p.Args, " "))
	case p.Path != "":
		value = fmt.Sprintf("%s: %s", value, p.Path)
	default:
		value = fmt.Sprintf("%s: %s", value, p.Name)
	}
	return value
}
