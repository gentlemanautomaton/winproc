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
	Threads     int
}

// String returns a string representation of the process.
func (p Process) String() string {
	value := fmt.Sprintf("[%d] PID %d", p.SessionID, p.ID)
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
