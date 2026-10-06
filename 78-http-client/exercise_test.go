package httpclient

import (
	"fmt"
	"maps"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"
)

const page = "line 1\nline 2\nline 3\nline 4\nline 5\nline 6\nline 7\n"

// newServer starts a local HTTP server for one test:
//
//	/page     200, the seven lines of the page constant
//	/crlf     200, three lines with Windows line endings
//	/empty    200, no body
//	/missing  404
//	/broken   500
func newServer(t *testing.T) *httptest.Server {
	t.Helper()
	mux := http.NewServeMux()
	mux.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, page)
	})
	mux.HandleFunc("/crlf", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "a\r\nb\r\nc\r\n")
	})
	mux.HandleFunc("/empty", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("/missing", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "nothing here", http.StatusNotFound)
	})
	mux.HandleFunc("/broken", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "kaputt", http.StatusInternalServerError)
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

// badURL cannot be fetched at all: its scheme is unknown, so the request
// fails immediately and without touching the network.
const badURL = "gbe-no-such-scheme://nowhere/"

// within runs f in a goroutine and returns its result, failing the test
// instead of hanging when f does not return within 2 seconds.
func within[T any](t *testing.T, what string, f func() T) T {
	t.Helper()
	done := make(chan T, 1)
	go func() { done <- f() }()
	select {
	case v := <-done:
		return v
	case <-time.After(2 * time.Second):
		t.Fatalf("%s did not return within 2 seconds", what)
	}
	panic("unreachable")
}

func TestStatus(t *testing.T) {
	srv := newServer(t)
	tests := []struct {
		path string
		want string
	}{
		{"/page", "200 OK"},
		{"/missing", "404 Not Found"},
		{"/broken", "500 Internal Server Error"},
	}
	for _, tt := range tests {
		got, err := Status(srv.URL + tt.path)
		if err != nil {
			t.Errorf("Status(server%s) returned error: %v", tt.path, err)
			continue
		}
		if got != tt.want {
			t.Errorf("Status(server%s) = %q, want %q", tt.path, got, tt.want)
		}
	}

	if got, err := Status(badURL); err == nil {
		t.Errorf("Status(%q) = %q, nil - want the error from http.Get", badURL, got)
	}
}

func TestFirstLines(t *testing.T) {
	srv := newServer(t)
	tests := []struct {
		path string
		n    int
		want []string
	}{
		{"/page", 5, []string{"line 1", "line 2", "line 3", "line 4", "line 5"}},
		{"/page", 1, []string{"line 1"}},
		{"/page", 100, []string{"line 1", "line 2", "line 3", "line 4", "line 5", "line 6", "line 7"}},
		{"/crlf", 2, []string{"a", "b"}},
		{"/page", 0, nil},
		{"/empty", 5, nil},
	}
	for _, tt := range tests {
		got, err := FirstLines(srv.URL+tt.path, tt.n)
		if err != nil {
			t.Errorf("FirstLines(server%s, %d) returned error: %v", tt.path, tt.n, err)
			continue
		}
		if !slices.Equal(got, tt.want) {
			t.Errorf("FirstLines(server%s, %d) = %q, want %q", tt.path, tt.n, got, tt.want)
		}
	}

	if got, err := FirstLines(badURL, 5); err == nil {
		t.Errorf("FirstLines(%q, 5) = %q, nil - want the error from http.Get", badURL, got)
	}
}

func TestFetch(t *testing.T) {
	srv := newServer(t)

	got, err := Fetch(srv.URL + "/page")
	if err != nil {
		t.Fatalf("Fetch(server/page) returned error: %v", err)
	}
	if got != page {
		t.Errorf("Fetch(server/page) = %q, want %q", got, page)
	}

	for path, status := range map[string]string{
		"/missing": "404 Not Found",
		"/broken":  "500 Internal Server Error",
	} {
		got, err := Fetch(srv.URL + path)
		if err == nil {
			t.Errorf("Fetch(server%s) = %q, nil - want an error for the status %s", path, got, status)
			continue
		}
		if got != "" {
			t.Errorf("Fetch(server%s) returned the body %q together with an error, want \"\"", path, got)
		}
		if !strings.Contains(err.Error(), status) {
			t.Errorf("Fetch(server%s) error = %q, want it to contain the status %q", path, err, status)
		}
	}

	if got, err := Fetch(badURL); err == nil {
		t.Errorf("Fetch(%q) = %q, nil - want the error from http.Get", badURL, got)
	}
}

func TestFetchWithTimeout(t *testing.T) {
	srv := newServer(t)

	got, err := FetchWithTimeout(srv.URL+"/page", 2*time.Second)
	if err != nil {
		t.Fatalf("FetchWithTimeout(server/page, 2s) returned error: %v", err)
	}
	if got != page {
		t.Errorf("FetchWithTimeout(server/page, 2s) = %q, want %q", got, page)
	}

	if got, err := FetchWithTimeout(srv.URL+"/missing", 2*time.Second); err == nil {
		t.Errorf("FetchWithTimeout(server/missing, 2s) = %q, nil - want an error for the status 404", got)
	}

	// A server that does not answer until the client gives up (or the test
	// ends, whichever comes first).
	release := make(chan struct{})
	slow := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
		case <-release:
		case <-time.After(3 * time.Second):
		}
	}))
	defer slow.Close()
	defer close(release)

	type result struct {
		body string
		err  error
	}
	start := time.Now()
	res := within(t, "FetchWithTimeout(slow server, 50ms)", func() result {
		body, err := FetchWithTimeout(slow.URL, 50*time.Millisecond)
		return result{body, err}
	})
	if res.err == nil {
		t.Errorf("FetchWithTimeout(slow server, 50ms) = %q, nil - want a timeout error", res.body)
	}
	if elapsed := time.Since(start); elapsed > time.Second {
		t.Errorf("FetchWithTimeout(slow server, 50ms) took %v, want it to give up after about 50ms", elapsed)
	}
}

func TestFetchAll(t *testing.T) {
	paths := []string{"/a", "/b", "/c", "/d"}

	// Every handler waits until all four requests have arrived. That only
	// happens when the requests are made concurrently; after half a second
	// the handlers give up waiting so that a sequential FetchAll does not
	// hang - it fails the "together" check below instead.
	var (
		mu       sync.Mutex
		inFlight int
		together = make(chan struct{})
		allHere  sync.Once
		gaveUp   = make(chan struct{})
		giveUp   sync.Once
	)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		inFlight++
		if inFlight == len(paths) {
			allHere.Do(func() { close(together) })
		}
		mu.Unlock()

		select {
		case <-together:
		case <-gaveUp:
		case <-time.After(500 * time.Millisecond):
			giveUp.Do(func() { close(gaveUp) })
		}

		mu.Lock()
		inFlight--
		mu.Unlock()
		fmt.Fprintf(w, "body of %s", r.URL.Path)
	}))
	defer srv.Close()

	type result struct {
		bodies map[string]string
		err    error
	}
	res := within(t, "FetchAll(server, 4 paths)", func() result {
		bodies, err := FetchAll(srv.URL, paths)
		return result{bodies, err}
	})
	if res.err != nil {
		t.Fatalf("FetchAll(server, %q) returned error: %v", paths, res.err)
	}
	want := map[string]string{
		"/a": "body of /a",
		"/b": "body of /b",
		"/c": "body of /c",
		"/d": "body of /d",
	}
	if !maps.Equal(res.bodies, want) {
		t.Fatalf("FetchAll(server, %q) = %v, want %v", paths, res.bodies, want)
	}
	select {
	case <-together:
	default:
		t.Error("the four requests were never in flight at the same time; start one goroutine per path")
	}
}

func TestFetchAllError(t *testing.T) {
	srv := newServer(t)
	type result struct {
		bodies map[string]string
		err    error
	}

	paths := []string{"/page", "/missing", "/empty"}
	res := within(t, "FetchAll(server, one failing path)", func() result {
		bodies, err := FetchAll(srv.URL, paths)
		return result{bodies, err}
	})
	if res.err == nil {
		t.Errorf("FetchAll(server, %q) = %v, nil - want an error because /missing answers 404", paths, res.bodies)
	} else if res.bodies != nil {
		t.Errorf("FetchAll(server, %q) returned the map %v together with an error, want nil", paths, res.bodies)
	}

	paths = []string{"/page", "/empty"}
	res = within(t, "FetchAll(server, two good paths)", func() result {
		bodies, err := FetchAll(srv.URL, paths)
		return result{bodies, err}
	})
	want := map[string]string{"/page": page, "/empty": ""}
	if res.err != nil || !maps.Equal(res.bodies, want) {
		t.Errorf("FetchAll(server, %q) = %v, %v, want %v, nil", paths, res.bodies, res.err, want)
	}
}
