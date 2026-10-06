package sha256hashes

import (
	"errors"
	"io"
	"slices"
	"strings"
	"testing"
	"testing/iotest"
)

// Known SHA-256 hashes.
const (
	hashEmpty      = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	hashABC        = "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	hashHelloWorld = "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"
	hashPage       = "1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a0d3db739d77aacb" // "sha256 this string"
	hashAlpha      = "8ed3f6ad685b959ead7022518e1af76cd816f8e8ec7ccdda1ed4018e8f2223f8"
	hashBravo      = "f144a6907dc4284d1f9fe6a7d9b9ff53c02c1d07ba68f24d413d7ff7f757a782"
)

func TestHashString(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"sha256 this string", hashPage},
		{"hello world", hashHelloWorld},
		{"abc", hashABC},
		{"", hashEmpty},
	}
	for _, tt := range tests {
		if got := HashString(tt.s); got != tt.want {
			t.Errorf("HashString(%q) =\n  %q, want\n  %q", tt.s, got, tt.want)
		}
	}
}

func TestHashParts(t *testing.T) {
	tests := []struct {
		parts []string
		want  string
	}{
		{[]string{"hello", " ", "world"}, hashHelloWorld},
		{[]string{"hello world"}, hashHelloWorld},
		{[]string{"", "hello wor", "", "ld"}, hashHelloWorld},
		{[]string{"a", "b", "c"}, hashABC},
		{[]string{"sha256", " this ", "string"}, hashPage},
		{nil, hashEmpty},
	}
	for _, tt := range tests {
		if got := HashParts(tt.parts...); got != tt.want {
			t.Errorf("HashParts(%q...) =\n  %q, want\n  %q", tt.parts, got, tt.want)
		}
	}
}

func TestHashReader(t *testing.T) {
	tests := []struct {
		name string
		r    io.Reader
		want string
	}{
		{`a reader with "hello world"`, strings.NewReader("hello world"), hashHelloWorld},
		{`a reader with "abc"`, strings.NewReader("abc"), hashABC},
		{"an empty reader", strings.NewReader(""), hashEmpty},
		// The data arriving in tiny pieces must not change the hash.
		{`a reader that delivers "sha256 this string" one byte at a time`,
			iotest.OneByteReader(strings.NewReader("sha256 this string")), hashPage},
	}
	for _, tt := range tests {
		got, err := HashReader(tt.r)
		if err != nil {
			t.Errorf("HashReader(%s) returned error %v, want nil", tt.name, err)
		}
		if got != tt.want {
			t.Errorf("HashReader(%s) =\n  %q, want\n  %q", tt.name, got, tt.want)
		}
	}

	errRead := errors.New("connection lost")
	failing := io.MultiReader(strings.NewReader("hello "), iotest.ErrReader(errRead))
	got, err := HashReader(failing)
	if !errors.Is(err, errRead) {
		t.Errorf("HashReader(a reader that fails) returned error %v, want %q", err, errRead)
	}
	if got != "" {
		t.Errorf("HashReader(a reader that fails) = %q together with an error, want an empty string", got)
	}
}

func TestVerify(t *testing.T) {
	tests := []struct {
		data    string
		wantHex string
		want    bool
	}{
		{"hello world", hashHelloWorld, true},
		{"hello world", strings.ToUpper(hashHelloWorld), true},
		{"abc", hashABC, true},
		{"", hashEmpty, true},
		{"hello world!", hashHelloWorld, false},
		{"hello world", hashABC, false},
		{"hello world", hashHelloWorld[:63], false},
		{"hello world", "", false},
		{"", "", false},
	}
	for _, tt := range tests {
		if got := Verify([]byte(tt.data), tt.wantHex); got != tt.want {
			t.Errorf("Verify(%q, %q) = %t, want %t", tt.data, tt.wantHex, got, tt.want)
		}
	}
}

func TestChecksumLines(t *testing.T) {
	files := map[string][]byte{
		"b.txt":     []byte("bravo"),
		"a.txt":     []byte("alpha"),
		"empty.dat": nil,
		"Z.txt":     []byte("abc"),
	}
	want := []string{
		hashABC + "  Z.txt",
		hashAlpha + "  a.txt",
		hashBravo + "  b.txt",
		hashEmpty + "  empty.dat",
	}
	// Map iteration order changes from run to run, so try a few times.
	for range 5 {
		got := ChecksumLines(files)
		if !slices.Equal(got, want) {
			t.Fatalf("ChecksumLines(4 files) =\n  %s\nwant\n  %s",
				strings.Join(got, "\n  "), strings.Join(want, "\n  "))
		}
	}

	if got := ChecksumLines(nil); len(got) != 0 {
		t.Errorf("ChecksumLines(nil) = %q, want no lines", got)
	}
}
