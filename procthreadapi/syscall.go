package procthreadapi

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// https://devblogs.microsoft.com/oldnewthing/20180216-00/?p=98035

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procIsProcessCritical = modkernel32.NewProc("IsProcessCritical")
	procTerminateProcess  = modkernel32.NewProc("TerminateProcess")
)

// IsProcessCritical returns true if the given process handle represents
// a critical system process. It calls the IsProcessCritical windows API
// function.
//
// This call is only supported on Windows 8.1 or newer.
//
// https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-isprocesscritical
func IsProcessCritical(process syscall.Handle) (critical bool, err error) {
	if err := procIsProcessCritical.Find(); err != nil {
		return false, err
	}

	r0, _, e := syscall.Syscall(
		procIsProcessCritical.Addr(),
		2,
		uintptr(process),
		uintptr(unsafe.Pointer(&critical)),
		0)
	if r0 == 0 {
		if e != 0 {
			err = syscall.Errno(e)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

// TerminateProcess attempts to terminate the process with the given handle.
// It calls the TerminateProcess windows API function.
//
// TerminateProcess returns nil if successful.
//
// https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-terminateprocess
func TerminateProcess(process syscall.Handle, exitCode uint32) (err error) {
	r0, _, e := syscall.Syscall(
		procTerminateProcess.Addr(),
		2,
		uintptr(process),
		uintptr(exitCode),
		0)
	if r0 == 0 {
		if e != 0 {
			err = syscall.Errno(e)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
