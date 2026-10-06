//go:build !solution

// Package jsonx contains exercises for https://gobyexample.com/json
//
// Run the tests for this section with:
//
//	go test ./55-json/
package jsonx

import "io"

// All exercises use the standard "encoding/json" package.

// Address is part of a User.
//
// TODO: add struct tags so that the fields are called "city" and "country"
// in JSON.
type Address struct {
	City    string
	Country string
}

// User is the record that the exercises encode and decode. Only exported
// fields take part in JSON, and without tags they would appear under their
// Go names ("Name", "Age", ...).
//
// TODO: add struct tags (like `json:"name"`) so that a User looks like this
// in JSON:
//
//	{"name":"Ann","age":30,"email":"ann@example.com","roles":["admin"],"address":{"city":"Oslo","country":"NO"}}
//
// The "email" key must be left out completely when Email is the empty
// string - that is what the ",omitempty" tag option does.
type User struct {
	Name    string
	Age     int
	Email   string
	Roles   []string
	Address Address
}

// ToJSON encodes v with json.Marshal and returns the JSON as a string. When
// v cannot be encoded (a channel or a function, for example) it returns an
// empty string and the error.
//
// Examples: ToJSON(true) -> "true", ToJSON("gopher") -> `"gopher"`,
// ToJSON([]string{"a", "b"}) -> `["a","b"]`,
// ToJSON(map[string]int{"apple": 5}) -> `{"apple":5}`.
func ToJSON(v any) (string, error) {
	// TODO: marshal v; Marshal returns a []byte.
	return "", nil
}

// DecodeUser decodes the JSON object in data into a User with
// json.Unmarshal. It returns the error from Unmarshal when data is not
// valid JSON or a value has the wrong type (e.g. a string for "age").
//
// Remember that Unmarshal needs a pointer to the value it should fill.
func DecodeUser(data string) (User, error) {
	// TODO: unmarshal into a User.
	return User{}, nil
}

// SumNumbers decodes data, a JSON object with arbitrary keys and values, and
// returns the sum of all values that are numbers. Values of other kinds
// (strings, booleans, arrays, ...) are ignored. Invalid JSON gives an error.
//
// Example: SumNumbers(`{"a":1.5,"b":"text","c":2,"d":[1,2]}`) returns 3.5.
//
// When you do not know the shape of the data up front, decode into a
// map[string]any. JSON numbers arrive in it as float64; use a type
// assertion or a type switch to pick them out.
func SumNumbers(data string) (float64, error) {
	// TODO: unmarshal into map[string]any and add up the float64 values.
	return 0, nil
}

// LookupCity digs the city out of nested JSON such as
//
//	{"name":"Ann","address":{"city":"Oslo","zip":"0150"}}
//
// without declaring a struct for it: decode into a map[string]any, then the
// value under "address" is itself a map[string]any, and the value under
// "city" in there is a string. For the document above it returns "Oslo".
//
// LookupCity must return an error (and never panic) when data is invalid
// JSON, when "address" is missing or not an object, or when "city" is
// missing or not a string. Hint: use the two-value form of the type
// assertion, v, ok := x.(T).
func LookupCity(data string) (string, error) {
	// TODO: unmarshal into a map and walk down with checked type assertions.
	return "", nil
}

// WriteJSON encodes v as JSON straight to w, followed by a newline, and
// returns any error from encoding or writing. A real program would pass
// os.Stdout or an http.ResponseWriter.
//
// Example: WriteJSON(w, map[string]int{"apple": 5}) writes `{"apple":5}`
// and "\n".
//
// Use json.NewEncoder(w); its Encode method adds the newline itself.
func WriteJSON(w io.Writer, v any) error {
	// TODO: encode v to w.
	return nil
}

// FilterAdults reads a stream of JSON-encoded users from r - several
// objects one after another, usually one per line - and writes those with
// an Age of 18 or more to w, each as one line of JSON, in the order they
// were read. It returns how many users it wrote.
//
// When the input contains something that cannot be decoded into a User,
// FilterAdults stops and returns the error (together with the number of
// users written so far). Reaching the end of the input is not an error.
//
// Use a json.Decoder for reading: its Decode method reads the next value
// from the stream and returns io.EOF when nothing is left. Use a
// json.Encoder for writing.
func FilterAdults(r io.Reader, w io.Writer) (int, error) {
	// TODO: decode users in a loop until io.EOF, encode the adults.
	return 0, nil
}
