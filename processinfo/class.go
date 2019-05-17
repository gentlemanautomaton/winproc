package processinfo

// Class is a windows process information class.
type Class uint32

// Windows process information classes.
const (
	BasicInfo                        Class = 0  // ProcessBasicInformation
	QuotaLimits                      Class = 1  // ProcessQuotaLimits
	IOCounters                       Class = 2  // ProcessIoCounters
	VirtualMemoryCounters            Class = 3  // ProcessVmCounters
	Times                            Class = 4  // ProcessTimes
	BasePriority                     Class = 5  // ProcessBasePriority
	RaisePriority                    Class = 6  // ProcessRaisePriority
	DebugPort                        Class = 7  // ProcessDebugPort
	ExceptionPort                    Class = 8  // ProcessExceptionPort
	AccessToken                      Class = 9  // ProcessAccessToken
	LocalDescriptorTableInfo         Class = 10 // ProcessLdtInformation
	LocalDescriptorTableSize         Class = 11 // ProcessLdtSize
	DefaultHardErrorMode             Class = 12 // ProcessDefaultHardErrorMode
	IOPortHandlers                   Class = 13 // ProcessIoPortHandlers
	PooledUsageAndLimits             Class = 14 // ProcessPooledUsageAndLimits
	WorkingSetWatch                  Class = 15 // ProcessWorkingSetWatch
	UserModeIOPrivilegeLevel         Class = 16 // ProcessUserModeIOPL
	EnableAlignmentFaultFixup        Class = 17 // ProcessEnableAlignmentFaultFixup
	PriorityClass                    Class = 18 // ProcessPriorityClass
	Wx86Info                         Class = 19 // ProcessWx86Information
	HandleCount                      Class = 20 // ProcessHandleCount
	AffinityMask                     Class = 21 // ProcessAffinityMask
	PriorityBoost                    Class = 22 // ProcessPriorityBoost
	DeviceMap                        Class = 23 // ProcessDeviceMap
	SessionInfo                      Class = 24 // ProcessSessionInformation
	ForegroundInfo                   Class = 25 // ProcessForegroundInformation
	Wow64Info                        Class = 26 // ProcessWow64Information
	ImageFileName                    Class = 27 // ProcessImageFileName
	LocallyUniqueIDDeviceMapsEnabled Class = 28 // ProcessLUIDDeviceMapsEnabled
	BreakOnTermination               Class = 29 // ProcessBreakOnTermination
	DebugObjectHandle                Class = 30 // ProcessDebugObjectHandle
	DebugFlags                       Class = 31 // ProcessDebugFlags
	HandleTracing                    Class = 32 // ProcessHandleTracing
	IOPriority                       Class = 33 // ProcessIoPriority
	ExecuteFlags                     Class = 34 // ProcessExecuteFlags
	ResourceManagement               Class = 35 // ProcessResourceManagement
	Cookie                           Class = 36 // ProcessCookie
	ImageInfo                        Class = 37 // ProcessImageInformation
	CycleTime                        Class = 38 // ProcessCycleTime
	PagePriority                     Class = 39 // ProcessPagePriority
	InstrumentationCallback          Class = 40 // ProcessInstrumentationCallback
	ThreadStackAllocation            Class = 41 // ProcessThreadStackAllocation
	WorkingSetWatchEx                Class = 42 // ProcessWorkingSetWatchEx
	ImageFileNameWin32               Class = 43 // ProcessImageFileNameWin32
	ImageFileMapping                 Class = 44 // ProcessImageFileMapping
	AffinityUpdateMode               Class = 45 // ProcessAffinityUpdateMode
	MemoryAllocationMode             Class = 46 // ProcessMemoryAllocationMode
	GroupInfo                        Class = 47 // ProcessGroupInformation
	TokenVirtualizationEnabled       Class = 48 // ProcessTokenVirtualizationEnabled
	ConsoleHostProcess               Class = 49 // ProcessConsoleHostProcess
	WindowInfo                       Class = 50 // ProcessWindowInformation
	HandleInfo                       Class = 51 // ProcessHandleInformation
	MitigationPolicy                 Class = 52 // ProcessMitigationPolicy
	DynamicFunctionTableInfo         Class = 53 // ProcessDynamicFunctionTableInformation
	HandleCheckingMode               Class = 54 // ProcessHandleCheckingMode
	KeepAliveCount                   Class = 55 // ProcessKeepAliveCount
	RevokeFileHandles                Class = 56 // ProcessRevokeFileHandles
	WorkingSetControl                Class = 57 // ProcessWorkingSetControl
	HandleTable                      Class = 58 // ProcessHandleTable
	CheckStackExtentsMode            Class = 59 // ProcessCheckStackExtentsMode
	CommandLineInfo                  Class = 60 // ProcessCommandLineInformation
	ProtectionInfo                   Class = 61 // ProcessProtectionInformation
	MemoryExhaustion                 Class = 62 // ProcessMemoryExhaustion
	FaultInfo                        Class = 63 // ProcessFaultInformation
	TelemetryIDInfo                  Class = 64 // ProcessTelemetryIdInformation
	CommitReleaseInfo                Class = 65 // ProcessCommitReleaseInformation
	DefaultCPUSetsInfo               Class = 66 // ProcessDefaultCpuSetsInformation
	AllowedCPUSetsInfo               Class = 67 // ProcessAllowedCpuSetsInformation
	_                                Class = 68 // ProcessReserved1Information
	_                                Class = 69 // ProcessReserved2Information
	SubsystemProcess                 Class = 70 // ProcessSubsystemProcess
	JobMemoryInfo                    Class = 71 // ProcessJobMemoryInformation
	SubsystemInfo                    Class = 75 // ProcessSubsystemInformation
)
