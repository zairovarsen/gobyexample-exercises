//go:build !solution

// Package httpserver contains exercises for https://gobyexample.com/http-server
//
// Run the tests for this section with:
//
//	go test ./79-http-server/
package httpserver

import (
	"net/http"
	"sync"
)

// A handler is a function with the signature
//
//	func(w http.ResponseWriter, r *http.Request)
//
// It answers a request by writing to w. The page registers its handlers
// with http.HandleFunc and blocks in http.ListenAndServe(":8090", nil); the
// tests call the handlers directly or start a server on a free port
// (net/http/httptest), so nothing here listens on a fixed port.

// Hello is the simplest handler: it responds with the body "hello\n".
//
// Hint: w is an io.Writer, so fmt.Fprintf(w, ...) works.
func Hello(w http.ResponseWriter, r *http.Request) {
	// TODO: write the response body.
}

// Headers echoes the headers of the request: for every header value it
// writes one line "Name: value\n" to the response. A header that was sent
// with several values produces several lines. The order of the lines does not
// matter.
//
// Example: a request with the header "Accept: text/plain" gets a body
// containing the line "Accept: text/plain\n".
//
// Hint: r.Header is a map from the header name to a slice of values.
func Headers(w http.ResponseWriter, r *http.Request) {
	// TODO: loop over the headers and their values, write a line for each.
}

// Greet reads the query parameter "name" from the request URL and responds
// with "hello, <name>\n": a request for /greet?name=gopher gets
// "hello, gopher\n".
//
// When the parameter is missing or empty, Greet responds with the status
// 400 Bad Request and the body "missing name\n" instead.
//
// Hints: r.URL.Query().Get("name") returns the (decoded) parameter;
// http.Error(w, message, status) sends an error response and appends the
// "\n" itself.
func Greet(w http.ResponseWriter, r *http.Request) {
	// TODO: read the parameter, answer with the greeting or with an error.
}

// NewMux returns a new request router (http.NewServeMux) with the three
// handlers registered on these paths:
//
//	/hello    Hello
//	/headers  Headers
//	/greet    Greet
//
// The page registers its handlers on the global default router with
// http.HandleFunc(path, handler); a ServeMux of your own has a method with
// the same name. A real program would then call
// http.ListenAndServe(":8090", NewMux()).
func NewMux() *http.ServeMux {
	// TODO: create the mux and register the handlers.
	return nil
}

// Counter is a handler with state: it counts the requests it has served.
// Because it has a ServeHTTP method it satisfies the http.Handler interface
// and can be registered with mux.Handle(path, counter).
//
// A server runs every request in its own goroutine, so n is shared between
// goroutines and must only be touched while holding mu.
type Counter struct {
	mu sync.Mutex
	n  int
}

// ServeHTTP increments the counter and responds with "count: N\n", where N
// is the number of requests served so far including this one: the first
// request gets "count: 1\n", the second "count: 2\n" and so on, also when
// many requests arrive at the same time.
func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: lock, increment, unlock, write the response.
}
