//go:build !solution

// Package sha256hashes contains exercises for https://gobyexample.com/sha256-hashes
//
// Run the tests for this section with:
//
//	go test ./63-sha256-hashes/
package sha256hashes

import "io"

// All exercises use the standard "crypto/sha256" package. A SHA-256 hash is
// 32 bytes long and is normally shown as 64 hexadecimal digits; fmt's %x
// verb (or the "encoding/hex" package) produces that form.

// HashString returns the SHA-256 hash of s as a lower-case hexadecimal
// string.
//
// Example: HashString("sha256 this string") returns
// "1af1dfa857bf1d8814fe1af8983c18080019922e557f15a8a0d3db739d77aacb".
//
// Follow the gobyexample page: create a hash with sha256.New(), Write the
// bytes of s to it, get the result with Sum(nil) and format it with %x.
func HashString(s string) string {
	// TODO: hash s and return the hex digest.
	return ""
}

// HashParts returns the hex-encoded SHA-256 hash of all parts taken
// together, as if they were one string: HashParts("hello", " ", "world")
// equals HashString("hello world").
//
// Do not concatenate the parts. A hash accepts its input in pieces: call
// Write once per part and Sum at the end.
func HashParts(parts ...string) string {
	// TODO: write every part to one hash.
	return ""
}

// HashReader returns the hex-encoded SHA-256 hash of everything that can be
// read from r, without loading it all into memory first - this is how a
// large file or a download is hashed (a real program would pass an *os.File
// or a response body). When reading fails it returns an empty string and
// the error.
//
// Hint: a hash is an io.Writer, so io.Copy(h, r) feeds it from a reader.
func HashReader(r io.Reader) (string, error) {
	// TODO: copy r into a hash and return the hex digest.
	return "", nil
}

// Verify reports whether the SHA-256 hash of data equals wantHex, a hash in
// hexadecimal notation as published next to a download. Upper-case and
// lower-case hex digits both have to be accepted.
//
// Hint: for a []byte that is already in memory, sha256.Sum256(data) is a
// shortcut that returns the hash as a [32]byte array.
func Verify(data []byte, wantHex string) bool {
	// TODO: hash data and compare with wantHex, ignoring case.
	return false
}

// ChecksumLines produces the kind of listing that the sha256sum command
// prints. files maps file names to file contents; the result has one line
// per file, sorted by file name, in the form
//
//	<64 hex digits><two spaces><file name>
//
// Example: for {"b.txt": "bravo", "a.txt": "alpha"} it returns
//
//	8ed3f6ad685b959ead7022518e1af76cd816f8e8ec7ccdda1ed4018e8f2223f8  a.txt
//	f144a6907dc4284d1f9fe6a7d9b9ff53c02c1d07ba68f24d413d7ff7f757a782  b.txt
//
// as a slice of two strings (without newlines).
func ChecksumLines(files map[string][]byte) []string {
	// TODO: sort the names, then build one line per file.
	return nil
}
