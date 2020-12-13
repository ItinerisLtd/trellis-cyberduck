package lib

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

type Io struct {
	// outWriter is a writer defined by the user that replaces stdout
	outWriter io.Writer
	// errWriter is a writer defined by the user that replaces stderr
	errWriter io.Writer
}

func NewIo() *Io {
	return &Io{
		outWriter: os.Stdout,
		errWriter: os.Stderr,
	}
}

func NewIoFromCobraCommand(cmd *cobra.Command) *Io {
	return &Io{
		outWriter: cmd.OutOrStdout(),
		errWriter: cmd.ErrOrStderr(),
	}
}

// SetOut sets the destination for usage messages.
// If newOut is nil, os.Stdout is used.
func (i *Io) SetOut(newOut io.Writer) {
	i.outWriter = newOut
}

// SetErr sets the destination for error messages.
// If newErr is nil, os.Stderr is used.
func (i *Io) SetErr(newErr io.Writer) {
	i.errWriter = newErr
}

// OutOrStdout returns output to stdout.
func (i *Io) OutOrStdout() io.Writer {
	if i.outWriter != nil {
		return i.outWriter
	}
	return os.Stdout
}

// ErrOrStderr returns output to stderr
func (i *Io) ErrOrStderr() io.Writer {
	if i.errWriter != nil {
		return i.errWriter
	}
	return os.Stderr
}
