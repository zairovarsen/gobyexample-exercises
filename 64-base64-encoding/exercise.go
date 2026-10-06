//go:build !solution

// Package base64encoding contains exercises for https://gobyexample.com/base64-encoding
//
// Run the tests for this section with:
//
//	go test ./64-base64-encoding/
package base64encoding

// All exercises use the standard "encoding/base64" package. It offers
// ready-made encodings as package variables; the two used here are
// base64.StdEncoding (the alphabet ends in "+" and "/") and
// base64.URLEncoding (which uses "-" and "_" instead, so that the result is
// safe in URLs and file names). Both have the methods EncodeToString, which
// takes a []byte, and DecodeString, which returns a []byte and an error.

// Encode returns s in standard base64.
//
// Example: Encode("abc123!?$*&()'-=@~") returns "YWJjMTIzIT8kKiYoKSctPUB+".
func Encode(s string) string {
	// TODO: encode the bytes of s with the standard encoding.
	return ""
}

// Decode is the reverse of Encode: it decodes standard base64 and returns
// the result as a string. When encoded is not valid standard base64 it
// returns an empty string and the error from the base64 package.
//
// Example: Decode("YWJjMTIzIT8kKiYoKSctPUB+") returns "abc123!?$*&()'-=@~".
func Decode(encoded string) (string, error) {
	// TODO: decode with the standard encoding.
	return "", nil
}

// EncodeURL returns s in the URL-safe variant of base64.
//
// Example: EncodeURL("abc123!?$*&()'-=@~") returns
// "YWJjMTIzIT8kKiYoKSctPUB-" (compare the last character with Encode).
func EncodeURL(s string) string {
	// TODO: encode with the URL encoding.
	return ""
}

// DecodeURL decodes URL-safe base64 and returns the result as a string, or
// an empty string and an error when encoded is not valid URL-safe base64.
func DecodeURL(encoded string) (string, error) {
	// TODO: decode with the URL encoding.
	return "", nil
}

// ToURLSafe converts a standard base64 string into the URL-safe variant by
// decoding it and encoding the bytes again. Invalid input gives an empty
// string and the decoding error.
//
// Example: ToURLSafe("Pj4+Pz8/") returns "Pj4-Pz8_".
func ToURLSafe(std string) (string, error) {
	// TODO: decode with one encoding, encode with the other.
	return "", nil
}

// BasicAuth builds the value of an HTTP "Authorization" header for basic
// authentication: the word "Basic", a space, and the standard base64
// encoding of "<user>:<password>".
//
// Example: BasicAuth("aladdin", "open sesame") returns
// "Basic YWxhZGRpbjpvcGVuIHNlc2FtZQ==".
func BasicAuth(user, password string) string {
	// TODO: join user and password with ":", encode, add the prefix.
	return ""
}

// ParseBasicAuth is the reverse of BasicAuth: it takes the header value
// apart again and returns the user name and the password.
//
// It returns an error (and empty strings) when header does not start with
// "Basic ", when the rest is not valid standard base64, or when the decoded
// text contains no ":". The user name ends at the first ":"; a password may
// contain further colons.
//
// Example: ParseBasicAuth("Basic YWxhZGRpbjpvcGVuIHNlc2FtZQ==") returns
// "aladdin", "open sesame", nil.
//
// Hint: strings.CutPrefix and strings.Cut make the string handling short.
func ParseBasicAuth(header string) (user, password string, err error) {
	// TODO: strip the prefix, decode, split at the first ":".
	return "", "", nil
}
