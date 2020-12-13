package lib

import (
	"os"
	"testing"
)

func TestBufferIo_OutOrStdout(t *testing.T) {
	i := &BufferIo{}
	if out := i.OutOrStdout(); out != os.Stdout {
		t.Errorf("Expected nil output to default to stdout")
	}
}

func TestBufferIo_ErrOrStderr(t *testing.T) {
	i := &BufferIo{}
	if out := i.ErrOrStderr(); out != os.Stderr {
		t.Errorf("Expected nil error to default to stderr")
	}
}
