//go:build solution

package linefilters

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func Upper(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if _, err := fmt.Fprintln(w, strings.ToUpper(scanner.Text())); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func NumberLines(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for n := 1; scanner.Scan(); n++ {
		if _, err := fmt.Fprintf(w, "%d: %s\n", n, scanner.Text()); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func Grep(r io.Reader, w io.Writer, substr string) (int, error) {
	matches := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, substr) {
			continue
		}
		if _, err := fmt.Fprintln(w, line); err != nil {
			return matches, err
		}
		matches++
	}
	return matches, scanner.Err()
}

func Uniq(r io.Reader, w io.Writer) error {
	// A separate flag (rather than comparing with "") keeps a leading empty
	// line from being mistaken for a repeat.
	var prev string
	first := true
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if !first && line == prev {
			continue
		}
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
		prev, first = line, false
	}
	return scanner.Err()
}

func Filter(r io.Reader, w io.Writer, steps ...func(line string) (string, bool)) error {
	scanner := bufio.NewScanner(r)
lines:
	for scanner.Scan() {
		line := scanner.Text()
		for _, step := range steps {
			var keep bool
			if line, keep = step(line); !keep {
				continue lines
			}
		}
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return scanner.Err()
}
