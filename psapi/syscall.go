// +build windows

package psapi

import (
	"io"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procCreateToolhelp32Snapshot = modkernel32.NewProc("CreateToolhelp32Snapshot")
	procProcess32First           = modkernel32.NewProc("Process32FirstW")
	procProcess32Next            = modkernel32.NewProc("Process32NextW")
	procModule32First            = modkernel32.NewProc("Module32FirstW")
	procModule32Next             = modkernel32.NewProc("Module32NextW")
)

// CreateSnapshot prepares a process, heap or module snapshot according to the
// provided flags. It calls the CreateToolhelp32Snapshot windows API function.
//
// It is the caller's responsibility to close the returned snapshot handle when
// finished with it by calling syscall.CloseHandle().
//
// https://docs.microsoft.com/en-us/windows/desktop/api/tlhelp32/nf-tlhelp32-createtoolhelp32snapshot
func CreateSnapshot(flags uint32, processID uint32) (handle syscall.Handle, err error) {
	const errBadLength = 24
	for round := 0; round < 256; round++ {
		r0, _, e := syscall.Syscall(
			procCreateToolhelp32Snapshot.Addr(),
			2,
			uintptr(flags),
			uintptr(processID),
			0)
		handle = syscall.Handle(r0)
		if handle == syscall.InvalidHandle {
			if e != 0 {
				err = syscall.Errno(e)
			} else {
				err = syscall.EINVAL
			}
			if e == errBadLength {
				continue
			}
		}
		return
	}
	return
}

// FirstProcess returns the first process entry from a snapshot.
// It calls the Process32FirstW windows API function.
//
// FirstProcess returns io.EOF if there are no processes in the snapshot.
//
// https://docs.microsoft.com/en-us/windows/desktop/api/tlhelp32/nf-tlhelp32-process32firstw
func FirstProcess(snapshot syscall.Handle) (entry ProcessEntry, err error) {
	entry.Size = uint32(unsafe.Sizeof(entry))

	r0, _, e := syscall.Syscall(
		procProcess32First.Addr(),
		2,
		uintptr(snapshot),
		uintptr(unsafe.Pointer(&entry)),
		0)

	if r0 == 0 {
		switch e {
		case 0:
			err = syscall.EINVAL
		case syscall.ERROR_NO_MORE_FILES:
			err = io.EOF
		default:
			err = syscall.Errno(e)
		}
	}

	return
}

// NextProcess returns the next process entry from a snapshot.
// It calls the Process32NextW windows API function.
//
// NextProcess returns io.EOF if there are no more processes in the snapshot.
//
// https://docs.microsoft.com/en-us/windows/desktop/api/tlhelp32/nf-tlhelp32-process32nextw
func NextProcess(snapshot syscall.Handle) (entry ProcessEntry, err error) {
	entry.Size = uint32(unsafe.Sizeof(entry))

	r0, _, e := syscall.Syscall(
		procProcess32Next.Addr(),
		2,
		uintptr(snapshot),
		uintptr(unsafe.Pointer(&entry)),
		0)

	if r0 == 0 {
		switch e {
		case 0:
			err = syscall.EINVAL
		case syscall.ERROR_NO_MORE_FILES:
			err = io.EOF
		default:
			err = syscall.Errno(e)
		}
	}

	return
}

// FirstModule returns the first module entry from a snapshot.
// It calls the Module32FirstW windows API function.
//
// FirstModule returns io.EOF if there are no modules in the snapshot.
//
// https://docs.microsoft.com/en-us/windows/desktop/api/tlhelp32/nf-tlhelp32-module32firstw
func FirstModule(snapshot syscall.Handle) (entry ModuleEntry, err error) {
	entry.Size = uint32(unsafe.Sizeof(entry))

	r0, _, e := syscall.Syscall(
		procModule32First.Addr(),
		2,
		uintptr(snapshot),
		uintptr(unsafe.Pointer(&entry)),
		0)

	if r0 == 0 {
		switch e {
		case 0:
			err = syscall.EINVAL
		case syscall.ERROR_NO_MORE_FILES:
			err = io.EOF
		default:
			err = syscall.Errno(e)
		}
	}

	return
}

// NextModule returns the next module entry from a snapshot.
// It calls the Module32NextW windows API function.
//
// NextModule returns io.EOF if there are no more modules in the snapshot.
//
// https://docs.microsoft.com/en-us/windows/desktop/api/tlhelp32/nf-tlhelp32-module32nextw
func NextModule(snapshot syscall.Handle) (entry ModuleEntry, err error) {
	entry.Size = uint32(unsafe.Sizeof(entry))

	r0, _, e := syscall.Syscall(
		procModule32Next.Addr(),
		2,
		uintptr(snapshot),
		uintptr(unsafe.Pointer(&entry)),
		0)

	if r0 == 0 {
		switch e {
		case 0:
			err = syscall.EINVAL
		case syscall.ERROR_NO_MORE_FILES:
			err = io.EOF
		default:
			err = syscall.Errno(e)
		}
	}

	return
}
