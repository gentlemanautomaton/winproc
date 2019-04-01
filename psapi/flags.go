// +build windows

package psapi

// Snapshot flags.
const (
	SnapHeapList = 0x00000001 // TH32CS_SNAPHEAPLIST
	SnapProcess  = 0x00000002 // TH32CS_SNAPPROCESS
	SnapThread   = 0x00000004 // TH32CS_SNAPTHREAD
	SnapModule   = 0x00000008 // TH32CS_SNAPMODULE
	SnapModule32 = 0x00000010 // TH32CS_SNAPMODULE32
	Inherit      = 0x80000000 // TH32CS_INHERIT

	SnapAll = SnapHeapList | SnapModule | SnapProcess | SnapThread // TH32CS_SNAPALL
)
