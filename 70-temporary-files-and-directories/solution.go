//go:build solution

package temporaryfilesanddirectories

import (
	"os"
	"path/filepath"
)

func WriteTemp(dir, prefix string, data []byte) (string, error) {
	f, err := os.CreateTemp(dir, prefix)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return "", err
	}
	return f.Name(), nil
}

func CreateTempExt(dir, prefix, ext string) (string, error) {
	f, err := os.CreateTemp(dir, prefix+"*"+ext)
	if err != nil {
		return "", err
	}
	if err := f.Close(); err != nil {
		return "", err
	}
	return f.Name(), nil
}

func Stage(dir, prefix string, files map[string][]byte) (string, error) {
	dname, err := os.MkdirTemp(dir, prefix)
	if err != nil {
		return "", err
	}
	for name, content := range files {
		fname := filepath.Join(dname, name)
		if err := os.WriteFile(fname, content, 0666); err != nil {
			return "", err
		}
	}
	return dname, nil
}

func Scratch(dir, prefix string) (path string, cleanup func() error, err error) {
	dname, err := os.MkdirTemp(dir, prefix)
	if err != nil {
		return "", nil, err
	}
	cleanup = func() error {
		return os.RemoveAll(dname)
	}
	return dname, cleanup, nil
}

func WithTempFile(dir, prefix string, fn func(f *os.File) error) error {
	f, err := os.CreateTemp(dir, prefix)
	if err != nil {
		return err
	}
	// Deferred calls run last-in-first-out: the file is closed first and
	// removed after that.
	defer os.Remove(f.Name())
	defer f.Close()

	return fn(f)
}
