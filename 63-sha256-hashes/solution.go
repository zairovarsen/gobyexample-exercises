//go:build solution

package sha256hashes

import (
	"crypto/sha256"
	"fmt"
	"io"
	"slices"
	"strings"
)

func HashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func HashParts(parts ...string) string {
	h := sha256.New()
	for _, part := range parts {
		h.Write([]byte(part))
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func HashReader(r io.Reader) (string, error) {
	h := sha256.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func Verify(data []byte, wantHex string) bool {
	got := fmt.Sprintf("%x", sha256.Sum256(data))
	return strings.EqualFold(got, wantHex)
}

func ChecksumLines(files map[string][]byte) []string {
	names := make([]string, 0, len(files))
	for name := range files {
		names = append(names, name)
	}
	slices.Sort(names)

	lines := make([]string, 0, len(names))
	for _, name := range names {
		lines = append(lines, fmt.Sprintf("%x  %s", sha256.Sum256(files[name]), name))
	}
	return lines
}
