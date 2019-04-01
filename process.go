// +build windows

package winproc

import "fmt"

// Process holds information about a windows process.
type Process struct {
	ID       int
	ParentID int
	Name     string
	Path     string
	Threads  int
}

// String returns a string representation of the process.
func (p Process) String() string {
	if p.Path == "" {
		return fmt.Sprintf("PID %d: %s", p.ID, p.Name)
	}
	return fmt.Sprintf("PID %d: %s (%s)", p.ID, p.Name, p.Path)
}
