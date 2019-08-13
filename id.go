// +build windows

package winproc

import "strconv"

// ID is a 32-bit windows process identifier.
type ID uint32

// String returns a string representation of the id.
func (id ID) String() string {
	return strconv.Itoa(int(id))
}
