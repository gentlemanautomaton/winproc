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
