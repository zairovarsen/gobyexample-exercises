//go:build solution

package deferx

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"
)

var (
	ErrInvalidAmount     = errors.New("amount must be positive")
	ErrInsufficientFunds = errors.New("insufficient funds")
)

func Bracket(w io.Writer, body func()) {
	fmt.Fprintln(w, "enter")
	defer fmt.Fprintln(w, "exit")
	body()
}

func ReverseOrder(names []string, record func(string)) {
	for _, name := range names {
		defer record(name)
	}
	record("body")
}

type Account struct {
	mu      sync.Mutex
	balance int
}

func (a *Account) Withdraw(amount int) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	if amount <= 0 {
		return ErrInvalidAmount
	}
	if amount > a.balance {
		return ErrInsufficientFunds
	}
	a.balance -= amount
	return nil
}

func WriteAll(wc io.WriteCloser, lines []string) (err error) {
	defer func() {
		// Keep an earlier write error; otherwise report the close error.
		if cerr := wc.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	for _, line := range lines {
		if _, werr := fmt.Fprintln(wc, line); werr != nil {
			return werr
		}
	}
	return nil
}

func WriteFile(path string, lines []string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := f.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	for _, line := range lines {
		if _, werr := fmt.Fprintln(f, line); werr != nil {
			return werr
		}
	}
	return nil
}
