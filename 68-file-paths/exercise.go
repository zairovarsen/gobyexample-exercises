//go:build !solution

// Package filepaths contains exercises for https://gobyexample.com/file-paths
//
// Run the tests for this section with:
//
//	go test ./68-file-paths/
package filepaths

// ConfigPath returns the path of the JSON config file called name inside the
// "config" directory under dir: ConfigPath("app", "db") is "app/config/db.json"
// on Linux/macOS and `app\config\db.json` on Windows.
//
// Always build paths with filepath.Join instead of gluing strings together
// with "/" or "\\": Join uses the separator of the current operating system
// and also cleans the result, so a dir like "app//" or "app/../app" gives the
// same answer as "app".
func ConfigPath(dir, name string) string {
	// TODO: join dir, "config" and the file name.
	return ""
}

// DirAndBase splits a path into the directory part and the last element.
//
// Example: "dir1/dir2/filename" gives "dir1/dir2" and "filename" (with the
// separators of the current OS); a bare "filename" gives "." and "filename".
//
// Hint: package path/filepath has one function for each half.
func DirAndBase(p string) (dir, base string) {
	// TODO: return the directory and the base name of p.
	return "", ""
}

// Stem returns the file name of p without its directory and without its
// extension: "dir/config.json" -> "config", "archive.tar.gz" -> "archive.tar"
// (only the last extension counts), "dir/README" -> "README".
//
// Hint: filepath.Ext finds the extension, strings.TrimSuffix removes it.
func Stem(p string) string {
	// TODO: take the base name and cut off the extension.
	return ""
}

// ReplaceExt returns p with its extension replaced by ext (which includes the
// dot). The directory part stays untouched; a path without an extension
// simply gets ext appended.
//
// Example: ReplaceExt("docs/report.txt", ".md") is "docs/report.md" and
// ReplaceExt("docs/notes", ".md") is "docs/notes.md".
func ReplaceExt(p, ext string) string {
	// TODO: trim the old extension, add the new one.
	return ""
}

// Absolute resolves p against the directory base: when p is already an
// absolute path it is returned (cleaned with filepath.Clean), otherwise it is
// joined onto base.
//
// Example on Linux: Absolute("/srv", "data/x.txt") is "/srv/data/x.txt",
// Absolute("/srv", "/etc/hosts") is "/etc/hosts".
//
// Hint: filepath.IsAbs.
func Absolute(base, p string) string {
	// TODO: check whether p is absolute and act accordingly.
	return ""
}

// IsInside reports whether target is base itself or lies somewhere below it.
//
// Example: with base "a/b", the targets "a/b" and "a/b/t/file" are inside;
// "a/c/t/file", "a" and "a/bc" are not.
//
// Use filepath.Rel to compute the path that leads from base to target: the
// target is outside exactly when that relative path climbs out of base, i.e.
// it is ".." or starts with ".." followed by a separator
// (filepath.Separator). A name such as "..hidden" is an ordinary file name!
// When Rel returns an error (one path is absolute and the other is not),
// target is not inside.
func IsInside(base, target string) bool {
	// TODO: compute the relative path and inspect how it starts.
	return false
}
