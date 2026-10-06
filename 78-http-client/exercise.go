//go:build !solution

// Package httpclient contains exercises for https://gobyexample.com/http-client
//
// Run the tests for this section with:
//
//	go test ./78-http-client/
package httpclient

import "time"

// The page fetches https://gobyexample.com. The exercises take the URL as a
// parameter instead, and the tests start a small local HTTP server
// (net/http/httptest) to stand in for the internet.
//
// One rule for every exercise: after a successful request the response body
// must be closed - write "defer resp.Body.Close()" right after the error
// check, as the page does. A body that is never closed leaks its network
// connection.

// Status sends a GET request to url with http.Get and returns the status
// line of the response (resp.Status), e.g. "200 OK" or "404 Not Found".
//
// A response with an error status is still a response: only a request that
// fails altogether (bad URL, no connection) makes Status return an error.
func Status(url string) (string, error) {
	// TODO: issue the request, close the body, return the status.
	return "", nil
}

// FirstLines fetches url and returns the first n lines of the response body,
// without their line endings. A body with fewer than n lines is returned
// completely.
//
// This is the page's "print the first 5 lines of the response body": read
// resp.Body with a bufio.Scanner and stop after n lines; check scanner.Err()
// at the end.
func FirstLines(url string, n int) ([]string, error) {
	// TODO: issue the request, scan up to n lines of the body.
	return nil, nil
}

// Fetch returns the whole response body of a GET request to url as a string.
//
// Unlike Status, Fetch treats every status other than 200 OK as a failure:
// it then returns "" and an error whose message contains the status (e.g.
// "404 Not Found").
//
// Hint: io.ReadAll reads a body to the end; resp.StatusCode and
// http.StatusOK are numbers, resp.Status is text.
func Fetch(url string) (string, error) {
	// TODO: issue the request, check the status code, read the body.
	return "", nil
}

// FetchWithTimeout works like Fetch but gives up when the complete exchange
// (connecting, sending the request, reading the response) takes longer than
// timeout; it then returns "" and an error.
//
// http.Get uses http.DefaultClient, which waits forever. Create your own
// client instead: &http.Client{Timeout: timeout} has a Get method too.
func FetchWithTimeout(url string, timeout time.Duration) (string, error) {
	// TODO: build a client with a timeout and fetch with it.
	return "", nil
}

// FetchAll fetches base+path for every element of paths and returns the
// bodies in a map from path to body. base is the address of a server such as
// "http://localhost:8080", the paths look like "/hello".
//
// The requests must be made concurrently - one goroutine per path - so that
// the total time is that of the slowest request, not the sum of all of them.
// Collect the results over a channel (do not write to the map from several
// goroutines). When any request fails or answers with a status other than
// 200 OK, FetchAll returns nil and an error, but only after all goroutines
// have delivered their result.
func FetchAll(base string, paths []string) (map[string]string, error) {
	// TODO: start a goroutine per path, gather bodies and errors.
	return nil, nil
}
