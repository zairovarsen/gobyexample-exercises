//go:build solution

package filepaths

import (
	"path/filepath"
	"strings"
)

func ConfigPath(dir, name string) string {
	return filepath.Join(dir, "config", name+".json")
}

func DirAndBase(p string) (dir, base string) {
	return filepath.Dir(p), filepath.Base(p)
}

func Stem(p string) string {
	filename := filepath.Base(p)
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}

func ReplaceExt(p, ext string) string {
	return strings.TrimSuffix(p, filepath.Ext(p)) + ext
}

func Absolute(base, p string) string {
	if filepath.IsAbs(p) {
		return filepath.Clean(p)
	}
	return filepath.Join(base, p)
}

func IsInside(base, target string) bool {
	rel, err := filepath.Rel(base, target)
	if err != nil {
		return false
	}
	return rel != ".." && !strings.HasPrefix(rel, ".."+string(filepath.Separator))
}
