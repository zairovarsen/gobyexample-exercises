//go:build solution

package httpserver

import (
	"fmt"
	"net/http"
	"sync"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func Headers(w http.ResponseWriter, r *http.Request) {
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "missing name", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "hello, %s\n", name)
}

func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", Hello)
	mux.HandleFunc("/headers", Headers)
	mux.HandleFunc("/greet", Greet)
	return mux
}

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	c.n++
	n := c.n
	c.mu.Unlock()

	fmt.Fprintf(w, "count: %d\n", n)
}
