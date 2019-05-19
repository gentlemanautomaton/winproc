// +build windows

package nativeapi

import (
	"syscall"
	"unsafe"

	"github.com/gentlemanautomaton/winproc/ntstatus"
	"github.com/gentlemanautomaton/winproc/processinfo"
	"golang.org/x/sys/windows"
)

var (
	modntdll = windows.NewLazySystemDLL("ntdll.dll")

	procQueryInformationProcess = modntdll.NewProc("NtQueryInformationProcess")
)

// ProcessCommandLine requests the command line of a process from the
// NT kernel. It calls ProcessInfo.
//
// This call is only supported on Windows 10 1511 or newer.
func ProcessCommandLine(process syscall.Handle) (commandLine string, err error) {
	var (
		b      [2048]byte
		buffer = b[:]
		length uint32
	)

	for i := 0; i < 3; i++ {
		length, err = ProcessInfo(process, processinfo.CommandLineInfo, buffer)
		switch err {
		case ntstatus.InfoLengthMismatch:
			buffer = make([]byte, int(length))
		case nil:
			// A successful ProcessInfo call fills the buffer with a
			// UnicodeString structure followed by the command line as utf16.
			const start = unsafe.Sizeof(unicodeString{})
			return utf16BytesToString(buffer[start:length]), nil
		default:
			return "", err
		}
	}

	return "", err
}

// ProcessInfo requests information about a process from the NT kernel.
// It calls the NtQueryInformationProcess NT native API function.
//
// The type of information to be retrieved is defined by the given
// information class.
//
// https://docs.microsoft.com/en-us/windows/desktop/api/winternl/nf-winternl-ntqueryinformationprocess
func ProcessInfo(process syscall.Handle, class processinfo.Class, buffer []byte) (n uint32, err error) {
	if len(buffer) == 0 {
		return 0, ErrEmptyBuffer
	}
	if err := procQueryInformationProcess.Find(); err != nil {
		return 0, err
	}

	r0, _, _ := syscall.Syscall6(
		procQueryInformationProcess.Addr(),
		5,
		uintptr(process),
		uintptr(class),
		uintptr(unsafe.Pointer(&buffer[0])),
		uintptr(len(buffer)),
		uintptr(unsafe.Pointer(&n)),
		0)
	if r0 != 0 {
		err = ntstatus.Value(r0)
	}
	return
}
