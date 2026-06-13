/*
Package logger — structured logging singleton (log/slog) with SDI wiring.

System bootstrap (blank import):

	import _ "github.com/omcrgnt/logger"

Registers stdout Output via res.AddBuiltin and configures slog.Default (info, json).
Optional user override: logger.Config and logger.OutputFileConfig via builder.Build + sdi.Resolve.

Public API: Debug, Info, Warn, Error(ctx, msg, args...).
*/
package logger
