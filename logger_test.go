package logger

import (
	"bytes"
	"context"
	"io"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/omcrgnt/res/unique"
	loggerv1 "github.com/omcrgnt/proto/gen/go/logger/v1"
)

type testOutput struct {
	buf bytes.Buffer
}

func (t *testOutput) Write(p []byte) (int, error) { return t.buf.Write(p) }

func (t *testOutput) markerLoggerOutput() {}

func setupUseDefaults(t *testing.T) *unique.Registry {
	t.Helper()
	active = noopLogger{}
	u := unique.New()
	if err := u.AddReplaceable(DefaultLog()); err != nil {
		t.Fatal(err)
	}
	if err := u.AddReplaceable(DefaultStdout()); err != nil {
		t.Fatal(err)
	}
	return u
}

func wireFromRegistry(t *testing.T, u *unique.Registry) {
	t.Helper()
	active = noopLogger{}
	logRaw, err := u.GetOneByInterface(reflect.TypeOf((*Logger)(nil)).Elem())
	if err != nil {
		t.Fatal(err)
	}
	outRaw, err := u.GetOneByInterface(reflect.TypeOf((*Output)(nil)).Elem())
	if err != nil {
		t.Fatal(err)
	}
	l, ok := logRaw.(*logger)
	if !ok {
		t.Fatalf("expected *logger, got %T", logRaw)
	}
	out, ok := outRaw.(Output)
	if !ok {
		t.Fatalf("expected Output, got %T", outRaw)
	}
	l.Inject([]any{out})
}

func TestNoop_withoutRegistry(t *testing.T) {
	active = noopLogger{}

	Info(context.Background(), "silent")
	Default().Info(context.Background(), "silent")
}

func TestConfig_build(t *testing.T) {
	active = noopLogger{}
	outRaw, err := OutputStdoutConfig{}.Build()
	if err != nil {
		t.Fatal(err)
	}
	logRaw, err := Config{
		Level:  &loggerv1.Level{Value: "info"},
		Format: &loggerv1.Format{Value: "json"},
	}.Build()
	if err != nil {
		t.Fatal(err)
	}
	l := logRaw.(*logger)
	l.Inject([]any{outRaw.(Output)})

	Info(context.Background(), "hello")
}

func TestDefault_fromRegistry(t *testing.T) {
	out := &testOutput{}
	active = noopLogger{}
	l := DefaultLog().(*logger)
	l.Inject([]any{out})

	Default().Info(context.Background(), "instance")
	if !strings.Contains(out.buf.String(), "instance") {
		t.Fatalf("expected instance log, got %q", out.buf.String())
	}
}

func TestLogImpl_Inject(t *testing.T) {
	out := &testOutput{}
	raw, err := Config{
		Level:  &loggerv1.Level{Value: "debug"},
		Format: &loggerv1.Format{Value: "text"},
	}.Build()
	if err != nil {
		t.Fatal(err)
	}
	l := raw.(*logger)
	l.Inject([]any{out})

	l.Debug(context.Background(), "debug-msg")
	if !strings.Contains(out.buf.String(), "debug-msg") {
		t.Fatalf("expected debug in output, got %q", out.buf.String())
	}
}

func TestOutputStdoutConfig(t *testing.T) {
	active = noopLogger{}
	outRaw, err := OutputStdoutConfig{}.Build()
	if err != nil {
		t.Fatal(err)
	}
	logRaw, err := Config{
		Level:  &loggerv1.Level{Value: "info"},
		Format: &loggerv1.Format{Value: "text"},
	}.Build()
	if err != nil {
		t.Fatal(err)
	}
	l := logRaw.(*logger)
	l.Inject([]any{outRaw.(Output)})

	Info(context.Background(), "stdout-line")
}

func TestOutputFileConfig(t *testing.T) {
	active = noopLogger{}

	dir := t.TempDir()
	path := dir + "/app.log"

	outRaw, err := OutputFileConfig{Path: path}.Build()
	if err != nil {
		t.Fatal(err)
	}
	logRaw, err := Config{
		Level:  &loggerv1.Level{Value: "info"},
		Format: &loggerv1.Format{Value: "text"},
	}.Build()
	if err != nil {
		t.Fatal(err)
	}
	l := logRaw.(*logger)
	l.Inject([]any{outRaw.(Output)})

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

func TestInfo_afterResolve(t *testing.T) {
	u := setupUseDefaults(t)
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	old := os.Stdout
	os.Stdout = w
	t.Cleanup(func() { os.Stdout = old })

	wireFromRegistry(t, u)

	Info(context.Background(), "bootstrap-test")

	done := make(chan []byte, 1)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		done <- buf.Bytes()
	}()
	_ = w.Close()
	out := <-done
	if !strings.Contains(string(out), "bootstrap-test") {
		t.Fatalf("stdout: %q", out)
	}
}

func TestRegistry_outputOverride(t *testing.T) {
	u := setupUseDefaults(t)
	dir := t.TempDir()
	path := dir + "/output-only.log"

	outBuilt, err := OutputFileConfig{Path: path}.Build()
	if err != nil {
		t.Fatal(err)
	}
	if err := u.Add(outBuilt); err != nil {
		t.Fatal(err)
	}
	logRaw, err := u.GetOneByInterface(reflect.TypeOf((*Logger)(nil)).Elem())
	if err != nil {
		t.Fatal(err)
	}
	logRaw.(*logger).Inject([]any{outBuilt.(Output)})

	Info(context.Background(), "output-only")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "output-only") {
		t.Fatalf("expected log in file, got %q", data)
	}
}

func TestRegistry_loggerOverride(t *testing.T) {
	u := setupUseDefaults(t)
	dir := t.TempDir()
	path := dir + "/integration.log"

	outBuilt, err := OutputFileConfig{Path: path}.Build()
	if err != nil {
		t.Fatal(err)
	}
	if err := u.Add(outBuilt); err != nil {
		t.Fatal(err)
	}
	logBuilt, err := Config{
		Level:  &loggerv1.Level{Value: "info"},
		Format: &loggerv1.Format{Value: "json"},
	}.Build()
	if err != nil {
		t.Fatal(err)
	}
	if err := u.Add(logBuilt); err != nil {
		t.Fatal(err)
	}
	logBuilt.(*logger).Inject([]any{outBuilt.(Output)})

	Info(context.Background(), "after-resolve")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "after-resolve") {
		t.Fatalf("expected log in file, got %q", data)
	}
}
