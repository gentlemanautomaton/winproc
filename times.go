//go:build windows
// +build windows

package winproc

import (
	"syscall"
	"time"
)

// Times holds time information about a windows process.
type Times struct {
	Creation time.Time     // Process creation time
	Exit     time.Time     // Process exit time
	Kernel   time.Duration // Time spent in kernel mode
	User     time.Duration // Time spent in user mode
}

func timeFromFiletime(ft syscall.Filetime) time.Time {
	if ft.HighDateTime == 0 && ft.LowDateTime == 0 {
		return time.Time{}
	}
	return time.Unix(0, ft.Nanoseconds())
}

func durationFromFiletime(ft syscall.Filetime) time.Duration {
	// 100-nanosecond intervals since January 1, 1601
	nsec := int64(ft.HighDateTime)<<32 + int64(ft.LowDateTime)
	// convert into nanoseconds
	nsec *= 100
	return time.Duration(nsec)
}
