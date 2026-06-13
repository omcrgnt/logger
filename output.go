package logger

import (
	"io"
	"os"
)

// Output is a sealed log sink port for SDI.
type Output interface {
	io.Writer
	markerLoggerOutput()
}

type stdoutOutput struct{}

func (stdoutOutput) Write(p []byte) (int, error) { return os.Stdout.Write(p) }

func (stdoutOutput) markerLoggerOutput() {}
