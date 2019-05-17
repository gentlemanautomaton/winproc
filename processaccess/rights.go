package processaccess

// Rights hold a set of windows process access rights.
type Rights uint32

// Windows process access rights.
//
// https://docs.microsoft.com/en-us/windows/desktop/procthread/process-security-and-access-rights
const (
	Terminate               Rights = 0x0001 // PROCESS_TERMINATE
	CreateThread            Rights = 0x0002 // PROCESS_CREATE_THREAD
	SetSessionID            Rights = 0x0004 // PROCESS_SET_SESSIONID
	VirtualMemoryOperation  Rights = 0x0008 // PROCESS_VM_OPERATION
	VirtualMemoryRead       Rights = 0x0010 // PROCESS_VM_READ
	VirtualMemoryWrite      Rights = 0x0020 // PROCESS_VM_WRITE
	CreateProcess           Rights = 0x0080 // PROCESS_CREATE_PROCESS
	SetQuota                Rights = 0x0100 // PROCESS_SET_QUOTA
	SetInformation          Rights = 0x0200 // PROCESS_SET_INFORMATION
	QueryInformation        Rights = 0x0400 // PROCESS_QUERY_INFORMATION
	SetPort                 Rights = 0x0800 // PROCESS_SET_PORT
	SuspendResume           Rights = 0x0800 // PROCESS_SUSPEND_RESUME
	QueryLimitedInformation Rights = 0x1000 // PROCESS_QUERY_LIMITED_INFORMATION
)
