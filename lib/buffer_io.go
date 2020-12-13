package lib

import (
	"bytes"
	"io"
)

type BufferIo struct {
	outWriter *bytes.Buffer
	errWriter *bytes.Buffer
	Io
}

func NewBufferIo() *BufferIo {
	outWriter := bytes.NewBufferString("")
	errWriter := bytes.NewBufferString("")

	return &BufferIo{
		outWriter: outWriter,
		errWriter: errWriter,
		Io: Io{
			outWriter: outWriter,
			errWriter: errWriter,
		},
	}
}

func (i *BufferIo) OutString() string {
	return i.outWriter.String()
}

func (i *BufferIo) ErrString() string {
	return i.errWriter.String()
}

// SetOut sets the destination for usage messages.
// If newOut is nil, os.Stdout is used.
func (i *BufferIo) SetOut(newOut io.Writer) {
	panic("method not implemented")
}

// SetErr sets the destination for error messages.
// If newErr is nil, os.Stderr is used.
func (i *BufferIo) SetErr(newErr io.Writer) {
	panic("method not implemented")
}
