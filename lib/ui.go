package lib

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

type Ui struct {
	// inReader is a reader defined by the user that replaces stdin
	inReader io.Reader
	// outWriter is a writer defined by the user that replaces stdout
	outWriter io.Writer
	// errWriter is a writer defined by the user that replaces stderr
	errWriter io.Writer
}

func NewUi() *Ui {
	return &Ui{
		inReader:  os.Stdin,
		outWriter: os.Stdout,
		errWriter: os.Stderr,
	}
}

func NewUiFromCobraCommand(cmd *cobra.Command) *Ui {
	return &Ui{
		inReader:  cmd.InOrStdin(),
		outWriter: cmd.OutOrStdout(),
		errWriter: cmd.ErrOrStderr(),
	}
}

// SetOut sets the destination for usage messages.
// If newOut is nil, os.Stdout is used.
func (u *Ui) SetOut(newOut io.Writer) {
	u.outWriter = newOut
}

// SetErr sets the destination for error messages.
// If newErr is nil, os.Stderr is used.
func (u *Ui) SetErr(newErr io.Writer) {
	u.errWriter = newErr
}

// SetIn sets the source for input data
// If newIn is nil, os.Stdin is used.
func (u *Ui) SetIn(newIn io.Reader) {
	u.inReader = newIn
}

// OutOrStdout returns output to stdout.
func (u *Ui) OutOrStdout() io.Writer {
	if u.outWriter != nil {
		return u.outWriter
	}
	return os.Stdout
}

// ErrOrStderr returns output to stderr
func (u *Ui) ErrOrStderr() io.Writer {
	if u.errWriter != nil {
		return u.errWriter
	}
	return os.Stderr
}

// InOrStdin returns input to stdin
func (u *Ui) InOrStdin() io.Reader {
	if u.inReader != nil {
		return u.inReader
	}
	return os.Stdin
}
