// +build windows

package nativeapi

import (
	"syscall"
	"unsafe"
)

// unicodeString is a structure that describes a utf16 buffer.
type unicodeString struct {
	Length    uint16
	MaxLength uint16
	Buffer    uintptr
}

func utf16BytesToString(s []byte) string {
	p := (*[0xffff]uint16)(unsafe.Pointer(&s[0]))
	return syscall.UTF16ToString(p[:len(s)/2])
}
