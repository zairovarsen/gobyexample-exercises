//go:build solution

package contextx

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Wait(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func Poll(ctx context.Context, interval time.Duration, check func() bool) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if check() {
				return nil
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func Generate(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for n := 0; ; n++ {
			select {
			case ch <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}

func SlowHandler(work time.Duration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		select {
		case <-time.After(work):
			fmt.Fprintf(w, "hello\n")
		case <-ctx.Done():
			err := ctx.Err()
			internalError := http.StatusInternalServerError
			http.Error(w, err.Error(), internalError)
		}
	}
}

func Get(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
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
