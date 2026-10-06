package logging

import (
	"bytes"
	"encoding/json"
	"log"
	"log/slog"
	"maps"
	"regexp"
	"strings"
	"testing"
)

// decodeLines parses output as JSON lines: one JSON object per line.
func decodeLines(t *testing.T, output string) []map[string]any {
	t.Helper()
	var records []map[string]any
	for _, line := range strings.Split(output, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		var record map[string]any
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			t.Fatalf("log line is not valid JSON: %q (%v)", line, err)
		}
		records = append(records, record)
	}
	return records
}

// checkField compares one field of a decoded log record. Numbers decoded
// from JSON are float64 values.
func checkField(t *testing.T, record map[string]any, key string, want any) {
	t.Helper()
	got, ok := record[key]
	if !ok {
		t.Errorf("log record %v has no field %q", record, key)
	} else if got != want {
		t.Errorf("log record field %q = %v, want %v", key, got, want)
	}
}

func TestNewLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewLogger(&buf, "my:")
	if logger == nil {
		t.Fatal(`NewLogger(w, "my:") returned nil`)
	}
	if got := logger.Flags(); got != 0 {
		t.Errorf("logger.Flags() = %d, want 0 (no date and time)", got)
	}
	if got := logger.Prefix(); got != "my:" {
		t.Errorf("logger.Prefix() = %q, want %q", got, "my:")
	}

	logger.Println("from mylog")
	logger.Printf("%d + %d = %d", 1, 2, 3)
	if got, want := buf.String(), "my:from mylog\nmy:1 + 2 = 3\n"; got != want {
		t.Errorf("the logger wrote %q, want %q", got, want)
	}

	// SetPrefix changes the prefix of an existing logger, as on the page.
	buf.Reset()
	logger.SetPrefix("ohmy:")
	logger.Println("from mylog")
	if got, want := buf.String(), "ohmy:from mylog\n"; got != want {
		t.Errorf("after SetPrefix the logger wrote %q, want %q", got, want)
	}
}

func TestCapture(t *testing.T) {
	tests := []struct {
		prefix string
		msgs   []string
		want   string
	}{
		{"buf:", []string{"hello", "world"}, "buf:hello\nbuf:world\n"},
		{"", []string{"plain"}, "plain\n"},
		{"[job] ", []string{"started", "50% done", "finished"}, "[job] started\n[job] 50% done\n[job] finished\n"},
		{"buf:", nil, ""},
	}
	for _, tt := range tests {
		if got := Capture(tt.prefix, tt.msgs...); got != tt.want {
			t.Errorf("Capture(%q, %q...) = %q, want %q", tt.prefix, tt.msgs, got, tt.want)
		}
	}
}

func TestNewDebugLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewDebugLogger(&buf, "app: ")
	if logger == nil {
		t.Fatal(`NewDebugLogger(w, "app: ") returned nil`)
	}
	if got, want := logger.Flags(), log.Lshortfile|log.Lmsgprefix; got != want {
		t.Errorf("logger.Flags() = %d, want %d (log.Lshortfile|log.Lmsgprefix)", got, want)
	}

	logger.Println("starting up")
	// The call above is in this file, so the line reads
	// "exercise_test.go:<line>: app: starting up".
	want := regexp.MustCompile(`^exercise_test\.go:\d+: app: starting up\n$`)
	if got := buf.String(); !want.MatchString(got) {
		t.Errorf(`the logger wrote %q, want "exercise_test.go:<line>: app: starting up\n"`, got)
	}
}

func TestNewJSONLogger(t *testing.T) {
	var buf bytes.Buffer
	logger := NewJSONLogger(&buf)
	if logger == nil {
		t.Fatal("NewJSONLogger(w) returned nil")
	}

	logger.Info("hi there")
	logger.Info("hello again", "key", "val", "age", 25)
	logger.Warn("careful")

	records := decodeLines(t, buf.String())
	if len(records) != 3 {
		t.Fatalf("the logger wrote %d JSON lines, want 3; output:\n%s", len(records), buf.String())
	}

	checkField(t, records[0], "msg", "hi there")
	checkField(t, records[0], "level", "INFO")
	if _, ok := records[0]["time"]; !ok {
		t.Errorf("log record %v has no \"time\" field", records[0])
	}

	checkField(t, records[1], "msg", "hello again")
	checkField(t, records[1], "level", "INFO")
	checkField(t, records[1], "key", "val")
	checkField(t, records[1], "age", 25.0)

	checkField(t, records[2], "msg", "careful")
	checkField(t, records[2], "level", "WARN")
}

func TestLogRequest(t *testing.T) {
	tests := []struct {
		method, path string
		status       int
		wantLevel    string
	}{
		{"GET", "/hello", 200, "INFO"},
		{"POST", "/users", 201, "INFO"},
		{"GET", "/missing", 404, "INFO"},
		{"GET", "/boom", 500, "ERROR"},
		{"PUT", "/gateway", 503, "ERROR"},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		logger := slog.New(slog.NewJSONHandler(&buf, nil))

		LogRequest(logger, tt.method, tt.path, tt.status)

		records := decodeLines(t, buf.String())
		if len(records) != 1 {
			t.Errorf("LogRequest(logger, %q, %q, %d) logged %d records, want 1; output: %q",
				tt.method, tt.path, tt.status, len(records), buf.String())
			continue
		}
		checkField(t, records[0], "msg", "request")
		checkField(t, records[0], "level", tt.wantLevel)
		checkField(t, records[0], "method", tt.method)
		checkField(t, records[0], "path", tt.path)
		checkField(t, records[0], "status", float64(tt.status))
	}
}

func TestCountLevels(t *testing.T) {
	// Produce a real slog JSON log to analyse.
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, nil))
	logger.Info("starting")
	logger.Warn("disk almost full", "free", "3%")
	logger.Error("write failed", "path", "/tmp/x")
	logger.Info("retrying", "attempt", 2)
	logger.Error("write failed", "path", "/tmp/x")
	logger.Error("giving up")

	got, err := CountLevels(&buf)
	if err != nil {
		t.Fatalf("CountLevels(slog output) returned error: %v", err)
	}
	want := map[string]int{"INFO": 2, "WARN": 1, "ERROR": 3}
	if !maps.Equal(got, want) {
		t.Errorf("CountLevels(slog output) = %v, want %v", got, want)
	}

	// Blank lines and Windows line endings are fine.
	input := "{\"level\":\"INFO\",\"msg\":\"a\"}\r\n\r\n{\"level\":\"DEBUG\",\"msg\":\"b\"}\r\n\n{\"level\":\"INFO\",\"msg\":\"c\"}"
	got, err = CountLevels(strings.NewReader(input))
	if err != nil {
		t.Fatalf("CountLevels(%q) returned error: %v", input, err)
	}
	want = map[string]int{"INFO": 2, "DEBUG": 1}
	if !maps.Equal(got, want) {
		t.Errorf("CountLevels(%q) = %v, want %v", input, got, want)
	}

	if got, err := CountLevels(strings.NewReader("")); err != nil || len(got) != 0 {
		t.Errorf(`CountLevels("") = %v, %v, want an empty result and a nil error`, got, err)
	}

	input = "{\"level\":\"INFO\",\"msg\":\"a\"}\nthis line is not JSON\n"
	if got, err := CountLevels(strings.NewReader(input)); err == nil {
		t.Errorf("CountLevels(%q) = %v, nil - want an error for the invalid line", input, got)
	}
}
