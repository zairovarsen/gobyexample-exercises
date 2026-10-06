package filepaths

import (
	"path/filepath"
	"testing"
)

// The test inputs are written with forward slashes and converted with
// filepath.FromSlash, and the expected values are built with filepath.Join,
// so the same tests hold on Windows (`a\b`) and on Linux/macOS ("a/b").
var native = filepath.FromSlash

func TestConfigPath(t *testing.T) {
	tests := []struct {
		dir, name string
		want      string
	}{
		{"app", "db", filepath.Join("app", "config", "db.json")},
		{native("srv/app"), "cache", filepath.Join("srv", "app", "config", "cache.json")},
		{native("dir1//"), "db", filepath.Join("dir1", "config", "db.json")},
		{native("dir1/../dir1"), "db", filepath.Join("dir1", "config", "db.json")},
	}
	for _, tt := range tests {
		if got := ConfigPath(tt.dir, tt.name); got != tt.want {
			t.Errorf("ConfigPath(%q, %q) = %q, want %q", tt.dir, tt.name, got, tt.want)
		}
	}
}

func TestDirAndBase(t *testing.T) {
	tests := []struct {
		p        string
		wantDir  string
		wantBase string
	}{
		{native("dir1/dir2/filename"), filepath.Join("dir1", "dir2"), "filename"},
		{native("dir/config.json"), "dir", "config.json"},
		{"filename", ".", "filename"},
		{native("dir/sub/"), filepath.Join("dir", "sub"), "sub"},
	}
	for _, tt := range tests {
		dir, base := DirAndBase(tt.p)
		if dir != tt.wantDir || base != tt.wantBase {
			t.Errorf("DirAndBase(%q) = %q, %q, want %q, %q", tt.p, dir, base, tt.wantDir, tt.wantBase)
		}
	}
}

func TestStem(t *testing.T) {
	tests := []struct {
		p    string
		want string
	}{
		{"config.json", "config"},
		{native("dir/config.json"), "config"},
		{native("a/b/archive.tar.gz"), "archive.tar"},
		{native("dir/README"), "README"},
		{native("release.d/notes"), "notes"},
	}
	for _, tt := range tests {
		if got := Stem(tt.p); got != tt.want {
			t.Errorf("Stem(%q) = %q, want %q", tt.p, got, tt.want)
		}
	}
}

func TestReplaceExt(t *testing.T) {
	tests := []struct {
		p, ext string
		want   string
	}{
		{"report.txt", ".md", "report.md"},
		{native("docs/report.txt"), ".md", filepath.Join("docs", "report.md")},
		{native("docs/notes"), ".md", filepath.Join("docs", "notes.md")},
		{native("a/archive.tar.gz"), ".zip", filepath.Join("a", "archive.tar.zip")},
		{native("release.d/notes"), ".txt", filepath.Join("release.d", "notes.txt")},
		{native("docs/report.txt"), "", filepath.Join("docs", "report")},
	}
	for _, tt := range tests {
		if got := ReplaceExt(tt.p, tt.ext); got != tt.want {
			t.Errorf("ReplaceExt(%q, %q) = %q, want %q", tt.p, tt.ext, got, tt.want)
		}
	}
}

func TestAbsolute(t *testing.T) {
	// A temporary directory is an absolute path on every operating system
	// ("/tmp/..." or `C:\Users\...`), so no OS-specific literal is needed.
	abs := t.TempDir()
	if !filepath.IsAbs(abs) {
		t.Fatalf("test setup: %q is not absolute", abs)
	}
	other := filepath.Join(abs, "other", "file.txt")

	tests := []struct {
		base, p string
		want    string
	}{
		{abs, native("data/x.txt"), filepath.Join(abs, "data", "x.txt")},
		{abs, "x.txt", filepath.Join(abs, "x.txt")},
		{abs, native("../sibling"), filepath.Join(filepath.Dir(abs), "sibling")},
		{abs, other, other},
		{"ignored", other, other},
		{"rel", native("data/x.txt"), filepath.Join("rel", "data", "x.txt")},
	}
	for _, tt := range tests {
		if got := Absolute(tt.base, tt.p); got != tt.want {
			t.Errorf("Absolute(%q, %q) = %q, want %q", tt.base, tt.p, got, tt.want)
		}
	}
}

func TestIsInside(t *testing.T) {
	abs := t.TempDir()
	tests := []struct {
		base, target string
		want         bool
	}{
		{native("a/b"), native("a/b/t/file"), true},
		{native("a/b"), native("a/c/t/file"), false},
		{native("a/b"), native("a/b"), true},
		{native("a/b"), "a", false},
		{native("a/b"), native("a/bc"), false},
		{native("a/b"), native("a/b/..hidden"), true},
		{native("a/b"), native("a/b/../b/file"), true},
		{native("a/b"), native("a/b/../../x"), false},
		{abs, filepath.Join(abs, "sub", "file.txt"), true},
		{filepath.Join(abs, "sub"), filepath.Join(abs, "file.txt"), false},
		{abs, native("relative/path"), false},
	}
	for _, tt := range tests {
		if got := IsInside(tt.base, tt.target); got != tt.want {
			t.Errorf("IsInside(%q, %q) = %v, want %v", tt.base, tt.target, got, tt.want)
		}
	}
}
