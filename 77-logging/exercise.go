//go:build !solution

// Package logging contains exercises for https://gobyexample.com/logging
//
// Run the tests for this section with:
//
//	go test ./77-logging/
package logging

import (
	"io"
	"log"
	"log/slog"
)

// NewLogger returns a logger from the standard log package that writes to w
// and puts prefix in front of every line - what the page does with
// log.New(os.Stdout, "my:", log.LstdFlags).
//
// One difference: use 0 as the flags argument instead of log.LstdFlags. The
// standard flags add the current date and time to every line, which no test
// could predict; with 0 a line is just prefix + message.
//
// Example: NewLogger(w, "my:").Println("from mylog") writes "my:from mylog\n".
func NewLogger(w io.Writer, prefix string) *log.Logger {
	// TODO: create the logger with log.New.
	return nil
}

// Capture logs every message in msgs on its own line with a logger that
// writes into a bytes.Buffer, and returns everything that was logged as one
// string. The logger uses the given prefix and the flags 0 (see NewLogger).
//
// Example: Capture("buf:", "hello", "world") is "buf:hello\nbuf:world\n".
//
// This is the page's "buflog": a logger can write to anything that is an
// io.Writer, and *bytes.Buffer is one.
func Capture(prefix string, msgs ...string) string {
	// TODO: declare a bytes.Buffer, log into it, return its content.
	return ""
}

// NewDebugLogger returns a logger writing to w whose lines show where in the
// source code the log call was made, followed by prefix and the message:
//
//	main.go:23: app: starting up
//
// That takes two flags combined with the | operator: log.Lshortfile adds the
// file name and line number of the caller, and log.Lmsgprefix moves the
// prefix from the start of the line to just before the message.
func NewDebugLogger(w io.Writer, prefix string) *log.Logger {
	// TODO: create the logger with both flags set.
	return nil
}

// NewJSONLogger returns a structured logger from the log/slog package that
// writes every record as one line of JSON to w, like the page's
//
//	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
//	myslog := slog.New(jsonHandler)
//
// Example: NewJSONLogger(w).Info("hello again", "key", "val", "age", 25)
// writes a line such as
//
//	{"time":"2023-08-22T10:45:16.9Z","level":"INFO","msg":"hello again","key":"val","age":25}
func NewJSONLogger(w io.Writer) *slog.Logger {
	// TODO: create a JSON handler for w and a logger that uses it.
	return nil
}

// LogRequest writes one structured log record about a handled HTTP request
// to logger. The message is "request" and the record carries three
// attributes with the keys "method", "path" and "status" (a number).
//
// The level depends on the status: 500 and above is logged with
// logger.Error, everything else with logger.Info.
//
// Example: LogRequest(logger, "GET", "/hello", 200) logs, with a JSON
// handler, something like
//
//	{"time":"...","level":"INFO","msg":"request","method":"GET","path":"/hello","status":200}
func LogRequest(logger *slog.Logger, method, path string, status int) {
	// TODO: pick the level, pass the attributes as key/value pairs.
}

// CountLevels reads a log in the JSON-lines format that slog's JSON handler
// produces (one JSON object per line) from r and counts the records per
// level: the keys of the result are the values of the "level" field, e.g.
// {"INFO": 12, "ERROR": 1}.
//
// Blank lines are skipped. A line that is not valid JSON makes CountLevels
// return nil and an error.
//
// This combines the line filter from an earlier section (bufio.Scanner) with
// encoding/json: decode each line into a small struct that has just the
// field you need.
func CountLevels(r io.Reader) (map[string]int, error) {
	// TODO: scan the lines, decode each one, count by level.
	return nil, nil
}
