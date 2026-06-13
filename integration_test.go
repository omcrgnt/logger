package logger_test

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/omcrgnt/logger"
	"github.com/omcrgnt/res"
	"github.com/omcrgnt/sdi"
	loggerv1 "github.com/omcrgnt/proto/gen/go/logger/v1"
)

func TestIntegration_dedupAndResolve(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/integration.log"

	fileRaw, err := logger.OutputFileConfig{Path: path}.Build()
	if err != nil {
		t.Fatal(err)
	}
	if err := res.Add(fileRaw); err != nil { //nolint:forbidigo // simulates app builder wiring
		t.Fatal(err)
	}

	logRaw, err := logger.Config{
		Level:  loggerv1.Level{Value: "info"},
		Format: loggerv1.Format{Value: "json"},
	}.Build()
	if err != nil {
		t.Fatal(err)
	}
	if err := res.Add(logRaw); err != nil { //nolint:forbidigo // simulates app builder wiring
		t.Fatal(err)
	}

	if err := sdi.Resolve(res.Default); err != nil {
		t.Fatal(err)
	}

	logger.Info(context.Background(), "after-resolve")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "after-resolve") {
		t.Fatalf("expected log in file, got %q", data)
	}
}
