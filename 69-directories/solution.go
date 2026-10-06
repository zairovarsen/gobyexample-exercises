//go:build solution

package directories

import (
	"io/fs"
	"os"
	"path/filepath"
)

func MakeDir(parent, name string) (string, error) {
	path := filepath.Join(parent, name)
	if err := os.Mkdir(path, 0755); err != nil {
		return "", err
	}
	return path, nil
}

func Scaffold(root string, files []string) error {
	for _, file := range files {
		path := filepath.Join(root, filepath.FromSlash(file))
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return err
		}
		if err := os.WriteFile(path, []byte(""), 0644); err != nil {
			return err
		}
	}
	return nil
}

func ListDir(dir string) ([]string, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() {
			name += "/"
		}
		names = append(names, name)
	}
	return names, nil
}

func FindFiles(root, ext string) ([]string, error) {
	var found []string
	visit := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != ext {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		found = append(found, filepath.ToSlash(rel))
		return nil
	}
	if err := filepath.WalkDir(root, visit); err != nil {
		return nil, err
	}
	return found, nil
}

func WithScratchDir(parent string, fn func(dir string) error) error {
	dir := filepath.Join(parent, "scratch")
	if err := os.Mkdir(dir, 0755); err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	return fn(dir)
}
