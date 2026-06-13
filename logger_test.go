package logger

import (
	"bytes"
	"context"
	"io"
	"os"
	"strings"
	"testing"

	loggerv1 "github.com/omcrgnt/proto/gen/go/logger/v1"
)

type testOutput struct {
	buf bytes.Buffer
}

func (t *testOutput) Write(p []byte) (int, error) { return t.buf.Write(p) }

func (t *testOutput) markerLoggerOutput() {}

func TestConfig_defaults(t *testing.T) {
	raw, err := Config{}.Build()
	if err != nil {
		t.Fatal(err)
	}
	l := raw.(*Logger)
	l.Inject([]any{&testOutput{}})

	Info(context.Background(), "hello")
}

func TestConfig_invalidLevel(t *testing.T) {
	t.Parallel()

	_, err := Config{
		Level: loggerv1.Level{Value: "trace"},
	}.Build()
	if err == nil {
		t.Fatal("expected error for invalid level")
	}
}

func TestConfig_invalidFormat(t *testing.T) {
	t.Parallel()

	_, err := Config{
		Format: loggerv1.Format{Value: "yaml"},
	}.Build()
	if err == nil {
		t.Fatal("expected error for invalid format")
	}
}

func TestLogger_Inject(t *testing.T) {
	out := &testOutput{}
	raw, err := Config{
		Level:  loggerv1.Level{Value: "debug"},
		Format: loggerv1.Format{Value: "text"},
	}.Build()
	if err != nil {
		t.Fatal(err)
	}
	l := raw.(*Logger)
	l.Inject([]any{out})

	Debug(context.Background(), "debug-msg")
	if !strings.Contains(out.buf.String(), "debug-msg") {
		t.Fatalf("expected debug in output, got %q", out.buf.String())
	}
}

func TestOutputFileConfig(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/app.log"

	raw, err := OutputFileConfig{Path: path}.Build()
	if err != nil {
		t.Fatal(err)
	}
	out := raw.(Output)

	rawLog, err := Config{
		Format: loggerv1.Format{Value: "text"},
	}.Build()
	if err != nil {
		t.Fatal(err)
	}
	l := rawLog.(*Logger)
	l.Inject([]any{out})

	Info(context.Background(), "file-line")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "file-line") {
		t.Fatalf("file content: %q", data)
	}
}

func TestOutputFileConfig_emptyPath(t *testing.T) {
	_, err := OutputFileConfig{}.Build()
	if err == nil {
		t.Fatal("expected error for empty path")
	}
}

func TestInfo_defaultBootstrap(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	old := os.Stdout
	os.Stdout = w
	t.Cleanup(func() { os.Stdout = old })

	if err := applyHandler(defaultLevelValue, defaultFormatValue, stdoutOutput{}); err != nil {
		t.Fatal(err)
	}

	done := make(chan []byte, 1)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		done <- buf.Bytes()
	}()

	Info(context.Background(), "bootstrap-test")
	_ = w.Close()
	out := <-done
	if !strings.Contains(string(out), "bootstrap-test") {
		t.Fatalf("stdout: %q", out)
	}
}
