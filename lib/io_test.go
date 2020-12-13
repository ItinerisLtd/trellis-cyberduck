package lib

import (
	"os"
	"testing"
)

func TestIo_SetOut(t *testing.T) {
	i := &Io{}
	i.SetOut(nil)
	if out := i.OutOrStdout(); out != os.Stdout {
		t.Errorf("Expected setting output to nil to revert back to stdout")
	}
}

func TestIo_SetErr(t *testing.T) {
	i := &Io{}
	i.SetErr(nil)
	if out := i.ErrOrStderr(); out != os.Stderr {
		t.Errorf("Expected setting error to nil to revert back to stderr")
	}
}
