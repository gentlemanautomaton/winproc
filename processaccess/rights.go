package processaccess

// Rights hold a set of windows process access rights.
type Rights uint32

// Windows process access rights.
//
// https://docs.microsoft.com/en-us/windows/desktop/procthread/process-security-and-access-rights
const (
	Terminate               Rights = 0x00000001 // PROCESS_TERMINATE
	CreateThread            Rights = 0x00000002 // PROCESS_CREATE_THREAD
	SetSessionID            Rights = 0x00000004 // PROCESS_SET_SESSIONID
	VirtualMemoryOperation  Rights = 0x00000008 // PROCESS_VM_OPERATION
	VirtualMemoryRead       Rights = 0x00000010 // PROCESS_VM_READ
	VirtualMemoryWrite      Rights = 0x00000020 // PROCESS_VM_WRITE
	CreateProcess           Rights = 0x00000080 // PROCESS_CREATE_PROCESS
	SetQuota                Rights = 0x00000100 // PROCESS_SET_QUOTA
	SetInformation          Rights = 0x00000200 // PROCESS_SET_INFORMATION
	QueryInformation        Rights = 0x00000400 // PROCESS_QUERY_INFORMATION
	SetPort                 Rights = 0x00000800 // PROCESS_SET_PORT
	SuspendResume           Rights = 0x00000800 // PROCESS_SUSPEND_RESUME
	QueryLimitedInformation Rights = 0x00001000 // PROCESS_QUERY_LIMITED_INFORMATION
	Synchronize             Rights = 0x00100000 // SYNCHRONIZE
)
