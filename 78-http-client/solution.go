//go:build solution

package httpclient

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Status(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return resp.Status, nil
}

func FirstLines(url string, n int) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var lines []string
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; i < n && scanner.Scan(); i++ {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func Fetch(url string) (string, error) {
	return fetch(http.DefaultClient, url)
}

func FetchWithTimeout(url string, timeout time.Duration) (string, error) {
	client := &http.Client{Timeout: timeout}
	return fetch(client, url)
}

func fetch(client *http.Client, url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GET %s: %s", url, resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func FetchAll(base string, paths []string) (map[string]string, error) {
	type result struct {
		path string
		body string
		err  error
	}
	results := make(chan result)
	for _, path := range paths {
		go func() {
			body, err := Fetch(base + path)
			results <- result{path, body, err}
		}()
	}

	// Receive exactly one result per goroutine, also after a failure, so
	// that none of them stays blocked on its send.
	bodies := make(map[string]string, len(paths))
	var firstErr error
	for range paths {
		r := <-results
		if r.err != nil {
			if firstErr == nil {
				firstErr = r.err
			}
			continue
		}
		bodies[r.path] = r.body
	}
	if firstErr != nil {
		return nil, firstErr
	}
	return bodies, nil
}
