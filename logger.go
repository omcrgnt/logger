package logger

// Logger is an SDI resource that reconfigures the process slog singleton.
type Logger struct {
	level  string
	format string
	out    Output
}

// Deps returns interface stubs required for wiring.
func (l *Logger) Deps() []any {
	return []any{(*Output)(nil)}
}

// Inject assigns Output from the pool and reconfigures slog.Default.
func (l *Logger) Inject(args []any) {
	for _, arg := range args {
		if out, ok := arg.(Output); ok {
			l.out = out
			break
		}
	}
	if l.out == nil {
		l.out = stdoutOutput{}
	}
	if err := applyHandler(l.level, l.format, l.out); err != nil {
		panic(err)
	}
}
