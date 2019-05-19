// +build windows

package winproc

import (
	"fmt"
	"syscall"

	"github.com/gentlemanautomaton/winproc/nativeapi"
	"github.com/gentlemanautomaton/winproc/processaccess"
)

// https://wj32.org/wp/2009/01/24/howto-get-the-command-line-of-processes/
// https://stackoverflow.com/questions/45891035/get-processid-from-binary-path-command-line-statement-in-c

// CommandLine attempts to retrieve the command line used to invoke a particular
// process.
//
// This call is only supported on Windows 10 1511 or newer.
//
// TODO: Add support for older operating systems by using an older, more
// arcane alternative when necessary.
func CommandLine(pid ID) (string, error) {
	process, err := syscall.OpenProcess(uint32(processaccess.QueryLimitedInformation), false, uint32(pid))
	if err != nil {
		return "", fmt.Errorf("failed to open handle for process %d: %v", pid, err)
	}
	defer syscall.CloseHandle(process)

	return nativeapi.ProcessCommandLine(process)
}
