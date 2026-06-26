/*
Package logger provides structured logging (log/slog) with resources in unique.

# Bootstrap

Blank-import the use subpackage at the app composition root (not logger itself):

	import _ "github.com/omcrgnt/logger/use"

	logger.Info(ctx, "started", "port", 8080)

logger/use registers [DefaultLog] and [DefaultStdout] via unique.MustAddReplaceable.
Without logger/use, package funcs are no-op until the logger resource is wired (Inject sets active).

# User override

Build and register user resources with unique.Add (replaces system defaults of the same type):

	logBuilt, err := cfg.Logger.Build()
	outBuilt, err := cfg.LogFile.Build()
	unique.Global().Add(outBuilt)
	unique.Global().Add(logBuilt)

Types [Config], [OutputStdoutConfig], and [OutputFileConfig] implement Build for override.

# Usage

Package funcs (Info, Debug, …) and [Default] delegate to the logger wired by Inject.

For DI, declare the port in Deps:

	func (c *Controller) Deps() []any {
	    return []any{(*logger.Logger)(nil)}
	}

# Wiring

Concrete logger implements [Logger]; Deps returns (*Output)(nil).
[DefaultLog] and [DefaultStdout] are for logger/use system registration.
*/
package logger
