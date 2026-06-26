package logger

import "os"

// OutputStdoutConfig registers stdout as the logger Output (user resource via Build).
type OutputStdoutConfig struct{}

// Build returns an Output writing to os.Stdout.
func (c OutputStdoutConfig) Build() (any, error) {
	return DefaultStdout(), nil
}

// DefaultStdout returns the system Output resource for logger/use registration.
func DefaultStdout() any {
	return stdoutOutput{}
}

type stdoutOutput struct{}

func (stdoutOutput) Write(p []byte) (int, error) { return os.Stdout.Write(p) }

func (stdoutOutput) markerLoggerOutput() {}
