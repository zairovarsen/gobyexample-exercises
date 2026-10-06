//go:build solution

package logging

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"strings"
)

func NewLogger(w io.Writer, prefix string) *log.Logger {
	return log.New(w, prefix, 0)
}

func Capture(prefix string, msgs ...string) string {
	var buf bytes.Buffer
	buflog := log.New(&buf, prefix, 0)
	for _, msg := range msgs {
		buflog.Println(msg)
	}
	return buf.String()
}

func NewDebugLogger(w io.Writer, prefix string) *log.Logger {
	return log.New(w, prefix, log.Lshortfile|log.Lmsgprefix)
}

func NewJSONLogger(w io.Writer) *slog.Logger {
	jsonHandler := slog.NewJSONHandler(w, nil)
	return slog.New(jsonHandler)
}

func LogRequest(logger *slog.Logger, method, path string, status int) {
	logAt := logger.Info
	if status >= 500 {
		logAt = logger.Error
	}
	logAt("request", "method", method, "path", path, "status", status)
}

func CountLevels(r io.Reader) (map[string]int, error) {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var record struct {
			Level string `json:"level"`
		}
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			return nil, err
		}
		counts[record.Level]++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return counts, nil
}
