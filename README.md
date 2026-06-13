# logger

Structured logging for omcrgnt services: singleton `log/slog`, SDI wiring, system stdout default.

## Quick start

```go
import _ "github.com/omcrgnt/logger"

logger.Info(ctx, "started", "port", 8080)
```

Blank import registers **stdout** `Output` (`res.AddBuiltin`) and configures `slog.Default` (level **info**, format **json**).

## Optional user override (app)

```go
type AppConfig struct {
    Logger  logger.Config           `ecfg:"LOGGER"`
    LogFile logger.OutputFileConfig `ecfg:"LOG_FILE"`
}
```

Pipeline:

```
builder.Build(cfg, res.Default)
sdi.Resolve(res.Default)  // Dedup Output: user file replaces system stdout
```

## SDI

- `Logger.Deps()` → `(*Output)(nil)`
- `OutputFileConfig.Build()` → user `Output`
- System `stdoutOutput` registered in `init`

## Design

See [demo/docs/res-sdi-coupling.md](https://github.com/omcrgnt/demo/blob/main/docs/res-sdi-coupling.md).

## Dev

```bash
task test
task lint
```

System module: use `lint/go/config/system.golangci.yml` (no `res.Add` in init-only paths; tests may use `res.Add`).
