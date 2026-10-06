//go:build solution

package writingfiles

import (
	"bufio"
	"os"
	"strings"
)

func WriteAll(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

func WriteParts(path string, data []byte, text string) (int, error) {
	f, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n1, err := f.Write(data)
	if err != nil {
		return n1, err
	}
	n2, err := f.WriteString(text)
	if err != nil {
		return n1 + n2, err
	}
	if err := f.Sync(); err != nil {
		return n1 + n2, err
	}
	return n1 + n2, nil
}

func WriteLines(path string, lines []string) (int, error) {
	f, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	total := 0
	for _, line := range lines {
		n, err := w.WriteString(line + "\n")
		total += n
		if err != nil {
			return total, err
		}
	}
	if err := w.Flush(); err != nil {
		return total, err
	}
	return total, nil
}

func CopyUpper(src, dst string) (int, error) {
	in, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer out.Close()

	w := bufio.NewWriter(out)
	lines := 0
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		if _, err := w.WriteString(strings.ToUpper(scanner.Text()) + "\n"); err != nil {
			return lines, err
		}
		lines++
	}
	if err := scanner.Err(); err != nil {
		return lines, err
	}
	if err := w.Flush(); err != nil {
		return lines, err
	}
	return lines, nil
}
