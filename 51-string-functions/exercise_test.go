package stringfunctions

import "testing"

func TestFileKind(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{".gitignore", "hidden"},
		{".hidden_test.go", "hidden"},
		{"types_gen.go", "generated"},
		{"types_gen_test.go", "test"},
		{"api_gen.pb.go", "generated"},
		{"main_test.go", "test"},
		{"main.go", "source"},
		{"test.go", "source"},
		{"README.md", "other"},
		{"main.go.bak", "other"},
		{"go", "other"},
	}
	for _, tt := range tests {
		if got := FileKind(tt.name); got != tt.want {
			t.Errorf("FileKind(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}

func TestSlug(t *testing.T) {
	tests := []struct {
		title string
		want  string
	}{
		{"Go by Example", "go-by-example"},
		{"String Functions", "string-functions"},
		{"already-a-slug", "already-a-slug"},
		{"Two  Spaces", "two--spaces"},
		{"SHA256 Hashes", "sha256-hashes"},
	}
	for _, tt := range tests {
		if got := Slug(tt.title); got != tt.want {
			t.Errorf("Slug(%q) = %q, want %q", tt.title, got, tt.want)
		}
	}
}

func TestCountWord(t *testing.T) {
	tests := []struct {
		text, word string
		want       int
	}{
		{"Go go GO gopher", "go", 4},
		{"The cat sat on the mat. THE END", "the", 3},
		{"The cat sat on the mat. THE END", "The", 3},
		{"test", "t", 2},
		{"cheese", "e", 3},
		{"nothing here", "xyz", 0},
	}
	for _, tt := range tests {
		if got := CountWord(tt.text, tt.word); got != tt.want {
			t.Errorf("CountWord(%q, %q) = %d, want %d", tt.text, tt.word, got, tt.want)
		}
	}
}

func TestUnderline(t *testing.T) {
	tests := []struct {
		title string
		want  string
	}{
		{"Hello", "Hello\n====="},
		{"Go", "Go\n=="},
		{"String Functions", "String Functions\n================"},
	}
	for _, tt := range tests {
		if got := Underline(tt.title); got != tt.want {
			t.Errorf("Underline(%q) = %q, want %q", tt.title, got, tt.want)
		}
	}
}

func TestDomain(t *testing.T) {
	tests := []struct {
		address string
		want    string
	}{
		{"gopher@example.com", "example.com"},
		{"a@b", "b"},
		{"first@second@third", "second@third"},
		{"trailing@", ""},
		{"no-at-sign", ""},
		{"@handle", "handle"},
	}
	for _, tt := range tests {
		if got := Domain(tt.address); got != tt.want {
			t.Errorf("Domain(%q) = %q, want %q", tt.address, got, tt.want)
		}
	}
}

func TestCleanList(t *testing.T) {
	tests := []struct {
		list string
		want string
	}{
		{"  Go, rust ,, ZIG ", "go,rust,zig"},
		{"a,b,c", "a,b,c"},
		{"Single", "single"},
		{"Hello World , two words", "hello world,two words"},
		{" , ,", ""},
		{"x,\tY\t,\nz\n", "x,y,z"},
	}
	for _, tt := range tests {
		if got := CleanList(tt.list); got != tt.want {
			t.Errorf("CleanList(%q) = %q, want %q", tt.list, got, tt.want)
		}
	}
}
