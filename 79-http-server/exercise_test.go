package httpserver

import (
	"io"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"
)

// serve calls a handler directly, without any network, and returns the
// recorded response.
func serve(handler http.HandlerFunc, req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	handler(rec, req)
	return rec
}

// sortedLines splits a body into its lines and sorts them.
func sortedLines(body string) []string {
	lines := strings.Split(strings.TrimSuffix(body, "\n"), "\n")
	slices.Sort(lines)
	return lines
}

func TestHello(t *testing.T) {
	rec := serve(Hello, httptest.NewRequest("GET", "/hello", nil))
	if rec.Code != http.StatusOK {
		t.Errorf("Hello responded with status %d, want %d", rec.Code, http.StatusOK)
	}
	if got, want := rec.Body.String(), "hello\n"; got != want {
		t.Errorf("Hello responded with the body %q, want %q", got, want)
	}
}

func TestHeaders(t *testing.T) {
	req := httptest.NewRequest("GET", "/headers", nil)
	req.Header.Set("User-Agent", "gbe-test/1.0")
	req.Header.Add("Accept", "text/plain")
	req.Header.Add("Accept", "text/html")
	req.Header.Set("X-Gbe-Test", "yes")

	rec := serve(Headers, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Headers responded with status %d, want %d", rec.Code, http.StatusOK)
	}
	body := rec.Body.String()
	if body != "" && !strings.HasSuffix(body, "\n") {
		t.Errorf("Headers responded with %q, want every line to end in \"\\n\"", body)
	}
	want := []string{
		"Accept: text/html",
		"Accept: text/plain",
		"User-Agent: gbe-test/1.0",
		"X-Gbe-Test: yes",
	}
	if got := sortedLines(body); !slices.Equal(got, want) {
		t.Errorf("Headers responded with the lines (sorted) %q, want %q", got, want)
	}
}

func TestGreet(t *testing.T) {
	tests := []struct {
		target   string
		wantCode int
		wantBody string
	}{
		{"/greet?name=gopher", http.StatusOK, "hello, gopher\n"},
		{"/greet?name=Go+by+Example", http.StatusOK, "hello, Go by Example\n"},
		{"/greet?other=1&name=ann", http.StatusOK, "hello, ann\n"},
		{"/greet", http.StatusBadRequest, "missing name\n"},
		{"/greet?name=", http.StatusBadRequest, "missing name\n"},
	}
	for _, tt := range tests {
		rec := serve(Greet, httptest.NewRequest("GET", tt.target, nil))
		if rec.Code != tt.wantCode {
			t.Errorf("Greet for %s responded with status %d, want %d", tt.target, rec.Code, tt.wantCode)
		}
		if got := rec.Body.String(); got != tt.wantBody {
			t.Errorf("Greet for %s responded with the body %q, want %q", tt.target, got, tt.wantBody)
		}
	}
}

func TestNewMux(t *testing.T) {
	mux := NewMux()
	if mux == nil {
		t.Fatal("NewMux() returned nil")
	}

	// A real server on a free local port, and a real client talking to it.
	srv := httptest.NewServer(mux)
	defer srv.Close()
	client := srv.Client()
	client.Timeout = 2 * time.Second

	get := func(path string) (int, string) {
		t.Helper()
		req, err := http.NewRequest("GET", srv.URL+path, nil)
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("X-Gbe-Test", "over the network")
		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("GET %s failed: %v", path, err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("GET %s: reading the body failed: %v", path, err)
		}
		return resp.StatusCode, string(body)
	}

	if code, body := get("/hello"); code != http.StatusOK || body != "hello\n" {
		t.Errorf("GET /hello = %d %q, want 200 %q", code, body, "hello\n")
	}
	if code, body := get("/greet?name=mux"); code != http.StatusOK || body != "hello, mux\n" {
		t.Errorf("GET /greet?name=mux = %d %q, want 200 %q", code, body, "hello, mux\n")
	}
	if code, _ := get("/greet"); code != http.StatusBadRequest {
		t.Errorf("GET /greet = %d, want 400", code)
	}
	if code, body := get("/headers"); code != http.StatusOK || !strings.Contains(body, "X-Gbe-Test: over the network\n") {
		t.Errorf("GET /headers = %d %q, want 200 and a body containing %q", code, body, "X-Gbe-Test: over the network\n")
	}
	if code, _ := get("/nope"); code != http.StatusNotFound {
		t.Errorf("GET /nope = %d, want 404 for a path without a handler", code)
	}
}

func TestCounter(t *testing.T) {
	counter := &Counter{}
	var handler http.Handler = counter // *Counter must satisfy http.Handler

	hit := func() string {
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, httptest.NewRequest("GET", "/count", nil))
		return rec.Body.String()
	}

	for _, want := range []string{"count: 1\n", "count: 2\n", "count: 3\n"} {
		if got := hit(); got != want {
			t.Fatalf("Counter responded with %q, want %q", got, want)
		}
	}

	// 200 requests at the same time: without the mutex some increments get
	// lost and the next count comes out too low.
	const parallel = 200
	bodies := make(chan string, parallel)
	var wg sync.WaitGroup
	for range parallel {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bodies <- hit()
		}()
	}
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("200 concurrent requests did not finish within 2 seconds - is the mutex unlocked again?")
	}
	close(bodies)

	seen := make(map[string]bool)
	for body := range bodies {
		if seen[body] {
			t.Errorf("two concurrent requests both got %q, want every request to get its own count", body)
		}
		seen[body] = true
	}
	if got, want := hit(), "count: 204\n"; got != want {
		t.Errorf("after 203 requests Counter responded with %q, want %q", got, want)
	}

	// Each Counter value counts on its own.
	other := &Counter{}
	rec := httptest.NewRecorder()
	other.ServeHTTP(rec, httptest.NewRequest("GET", "/count", nil))
	if got, want := rec.Body.String(), "count: 1\n"; got != want {
		t.Errorf("a fresh Counter responded with %q, want %q", got, want)
	}
}
