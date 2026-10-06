//go:build solution

package readingfiles

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadAll(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ReadFirst(path string, n int) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := make([]byte, n)
	read, err := f.Read(buf)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return buf[:read], nil
}

func ReadSection(path string, offset int64, n int) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if _, err := f.Seek(offset, io.SeekStart); err != nil {
		return nil, err
	}
	buf := make([]byte, n)
	if _, err := io.ReadAtLeast(f, buf, n); err != nil {
		return nil, err
	}
	return buf, nil
}

func ReadTail(path string, n int64) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Seeking to the end returns the offset of the end: the file size.
	size, err := f.Seek(0, io.SeekEnd)
	if err != nil {
		return nil, err
	}
	n = min(n, size)
	if _, err := f.Seek(-n, io.SeekEnd); err != nil {
		return nil, err
	}
	buf := make([]byte, n)
	if _, err := io.ReadAtLeast(f, buf, len(buf)); err != nil {
		return nil, err
	}
	return buf, nil
}

func PeekAndReadLine(path string, n int) (peeked, line string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	head, err := r.Peek(n)
	if err != nil {
		return "", "", err
	}
	peeked = string(head)

	// io.EOF only means that the last line had no line ending.
	line, err = r.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", "", err
	}
	return peeked, strings.TrimRight(line, "\r\n"), nil
}

func FindLines(path, substr string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var found []string
	scanner := bufio.NewScanner(f)
	for n := 1; scanner.Scan(); n++ {
		if line := scanner.Text(); strings.Contains(line, substr) {
			found = append(found, fmt.Sprintf("%d: %s", n, line))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return found, nil
}
