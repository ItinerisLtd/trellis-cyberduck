package lib

import "io"

type OutErrWriter interface {
	SetOut(newOut io.Writer)
	SetErr(newErr io.Writer)
	OutOrStdout() io.Writer
	ErrOrStderr() io.Writer
}
