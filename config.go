package logger

import loggerv1 "github.com/omcrgnt/proto/gen/go/logger/v1"

// Config configures the process logger (level, format).
type Config struct {
	Level  loggerv1.Level  `ecfg:"LEVEL"`
	Format loggerv1.Format `ecfg:"FORMAT"`
}

// Build returns a Logger resource for res.Add and SDI.
func (c Config) Build() (any, error) { //nolint:govet // ecfg value receiver; only GetValue() used
	level := c.Level.GetValue()
	format := c.Format.GetValue()
	level, format = mergeDefaults(level, format)
	if err := validateLevel(level); err != nil {
		return nil, err
	}
	if err := validateFormat(format); err != nil {
		return nil, err
	}
	return &Logger{
		level:  level,
		format: format,
	}, nil
}
