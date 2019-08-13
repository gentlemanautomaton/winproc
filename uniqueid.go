// +build windows

package winproc

import (
	"encoding/binary"
	"strconv"
)

// UniqueID is a 12-byte identifier that uniquely identifies a process
// on a host machine by combining its creation time and process ID.
type UniqueID struct {
	Creation int64 // Nanoseconds since (00:00:00 UTC, January 1, 1970)
	ID       ID
}

// String returns a string representation of the identifier.
func (uid UniqueID) String() string {
	return strconv.FormatInt(uid.Creation, 10) + "." + strconv.FormatUint(uint64(uid.ID), 10)
}

// Bytes returns uid as a sequence of bytes.
func (uid UniqueID) Bytes() (b [12]byte) {
	binary.LittleEndian.PutUint64(b[0:8], uint64(uid.Creation))
	binary.LittleEndian.PutUint32(b[8:12], uint32(uid.ID))
	return b
}
