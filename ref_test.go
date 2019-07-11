// +build windows

package winproc_test

import (
	"context"
	"fmt"
	"os/exec"
	"testing"
	"time"

	"github.com/gentlemanautomaton/winproc"
	"github.com/gentlemanautomaton/winproc/processaccess"
)

var delays = []time.Duration{
	100 * time.Millisecond,
	200 * time.Millisecond,
	400 * time.Millisecond,
}

func TestWait(t *testing.T) {
	for _, delay := range delays {
		delay := delay // Capture range variable
		t.Run(fmt.Sprintf("Delay=%s", delay), func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), delay)
			defer cancel()

			command := exec.CommandContext(ctx, "notepad")
			if err := command.Start(); err != nil {
				t.Error(err)
			}

			ref, err := winproc.Open(winproc.ID(command.Process.Pid), processaccess.Synchronize)
			if err != nil {
				t.Error(err)
			}
			defer ref.Close()

			if err := ref.Wait(context.Background()); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestWaitCancellation(t *testing.T) {
	for _, delay := range delays {
		delay := delay // Capture range variable
		t.Run(fmt.Sprintf("Delay=%s", delay), func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()

			command := exec.Command("notepad")
			if err := command.Start(); err != nil {
				t.Error(err)
			}
			defer command.Process.Kill()

			ref, err := winproc.Open(winproc.ID(command.Process.Pid), processaccess.Synchronize)
			if err != nil {
				t.Error(err)
			}
			defer ref.Close()

			if err := ref.Wait(ctx); err != context.DeadlineExceeded {
				t.Error(err)
			}
		})
	}
}
