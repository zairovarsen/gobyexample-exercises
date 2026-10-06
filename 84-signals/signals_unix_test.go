//go:build !windows

package signals

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

// On Unix the test process can send itself a real signal.
func TestNewSignalChannelReceivesSIGTERM(t *testing.T) {
	// Registering a channel of our own first keeps SIGTERM from terminating
	// the test when NewSignalChannel does not register anything.
	guard := make(chan os.Signal, 1)
	signal.Notify(guard, syscall.SIGTERM)
	defer signal.Stop(guard)

	sigs := NewSignalChannel()
	if sigs == nil {
		t.Fatal("NewSignalChannel() returned a nil channel")
	}
	defer signal.Stop(sigs)

	if err := syscall.Kill(syscall.Getpid(), syscall.SIGTERM); err != nil {
		t.Fatalf("could not send SIGTERM to the test process: %v", err)
	}
	select {
	case sig := <-sigs:
		if sig != syscall.SIGTERM {
			t.Errorf("received %v, want %v", sig, syscall.SIGTERM)
		}
	case <-time.After(2 * time.Second):
		t.Error("SIGTERM was not delivered on the channel - is it registered with signal.Notify?")
	}
}
