package jsonx

import (
	"bytes"
	"encoding/json"
	"errors"
	"math"
	"reflect"
	"strings"
	"testing"
)

// brokenWriter is an io.Writer whose every Write fails.
type brokenWriter struct{}

func (brokenWriter) Write([]byte) (int, error) { return 0, errors.New("broken writer") }

func TestToJSON(t *testing.T) {
	tests := []struct {
		v    any
		want string
	}{
		{true, `true`},
		{1, `1`},
		{2.34, `2.34`},
		{"gopher", `"gopher"`},
		{[]string{"apple", "peach", "pear"}, `["apple","peach","pear"]`},
		{map[string]int{"apple": 5, "lettuce": 7}, `{"apple":5,"lettuce":7}`},
		{nil, `null`},
	}
	for _, tt := range tests {
		got, err := ToJSON(tt.v)
		if err != nil {
			t.Errorf("ToJSON(%#v) returned error %v, want nil", tt.v, err)
		}
		if got != tt.want {
			t.Errorf("ToJSON(%#v) = %s, want %s", tt.v, got, tt.want)
		}
	}

	got, err := ToJSON(make(chan int))
	if err == nil {
		t.Errorf("ToJSON(a channel) returned a nil error, want the error from json.Marshal")
	}
	if got != "" {
		t.Errorf("ToJSON(a channel) = %q together with an error, want an empty string", got)
	}
}

// TestUserTags does not call any of your functions: it checks the struct
// tags on User and Address by encoding a User with json.Marshal.
func TestUserTags(t *testing.T) {
	tests := []struct {
		user User
		want string
	}{
		{
			User{
				Name:    "Ann",
				Age:     30,
				Email:   "ann@example.com",
				Roles:   []string{"admin"},
				Address: Address{City: "Oslo", Country: "NO"},
			},
			`{"name":"Ann","age":30,"email":"ann@example.com","roles":["admin"],"address":{"city":"Oslo","country":"NO"}}`,
		},
		{
			User{
				Name:    "Bob",
				Age:     12,
				Roles:   []string{"guest", "kid"},
				Address: Address{City: "Lima", Country: "PE"},
			},
			`{"name":"Bob","age":12,"roles":["guest","kid"],"address":{"city":"Lima","country":"PE"}}`,
		},
	}
	for _, tt := range tests {
		data, err := json.Marshal(tt.user)
		if err != nil {
			t.Fatalf("json.Marshal(%+v) failed: %v", tt.user, err)
		}
		if got := string(data); got != tt.want {
			t.Errorf("json.Marshal(%+v) =\n  %s, want\n  %s\ncheck the struct tags", tt.user, got, tt.want)
		}
	}
}

func TestDecodeUser(t *testing.T) {
	data := `{
		"name": "Ann",
		"age": 30,
		"email": "ann@example.com",
		"roles": ["admin", "dev"],
		"address": {"city": "Oslo", "country": "NO"},
		"unknown": "keys are ignored"
	}`
	want := User{
		Name:    "Ann",
		Age:     30,
		Email:   "ann@example.com",
		Roles:   []string{"admin", "dev"},
		Address: Address{City: "Oslo", Country: "NO"},
	}
	got, err := DecodeUser(data)
	if err != nil {
		t.Errorf("DecodeUser(%s) returned error %v, want nil", data, err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("DecodeUser(%s) =\n  %+v, want\n  %+v", data, got, want)
	}

	got, err = DecodeUser(`{"name":"Bob"}`)
	if err != nil {
		t.Errorf(`DecodeUser({"name":"Bob"}) returned error %v, want nil`, err)
	}
	if want := (User{Name: "Bob"}); !reflect.DeepEqual(got, want) {
		t.Errorf(`DecodeUser({"name":"Bob"}) = %+v, want %+v`, got, want)
	}
}

func TestDecodeUserErrors(t *testing.T) {
	for _, data := range []string{
		`{"name":"Ann"`,
		`not json`,
		`{"name":"Ann","age":"thirty"}`,
		`["a","list"]`,
		``,
	} {
		if _, err := DecodeUser(data); err == nil {
			t.Errorf("DecodeUser(%q) returned a nil error, want an error", data)
		}
	}
}

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		data string
		want float64
	}{
		{`{"a":1.5,"b":"text","c":2,"d":[1,2]}`, 3.5},
		{`{"num":6.13,"strs":["a","b"]}`, 6.13},
		{`{"x":10,"y":-4,"nested":{"z":100},"flag":true,"nothing":null,"s":"7"}`, 6},
		{`{}`, 0},
	}
	for _, tt := range tests {
		got, err := SumNumbers(tt.data)
		if err != nil {
			t.Errorf("SumNumbers(%s) returned error %v, want nil", tt.data, err)
		}
		if math.Abs(got-tt.want) > 1e-9 {
			t.Errorf("SumNumbers(%s) = %v, want %v", tt.data, got, tt.want)
		}
	}

	for _, data := range []string{`{"a":1`, `[1,2,3]`, `oops`} {
		if _, err := SumNumbers(data); err == nil {
			t.Errorf("SumNumbers(%q) returned a nil error, want an error", data)
		}
	}
}

func TestLookupCity(t *testing.T) {
	tests := []struct {
		data string
		want string
	}{
		{`{"name":"Ann","address":{"city":"Oslo","zip":"0150"}}`, "Oslo"},
		{`{"address":{"city":"Lima"}}`, "Lima"},
	}
	for _, tt := range tests {
		got, err := LookupCity(tt.data)
		if err != nil {
			t.Errorf("LookupCity(%s) returned error %v, want nil", tt.data, err)
		}
		if got != tt.want {
			t.Errorf("LookupCity(%s) = %q, want %q", tt.data, got, tt.want)
		}
	}
}

func TestLookupCityErrors(t *testing.T) {
	tests := []struct {
		why  string
		data string
	}{
		{"it is not valid JSON", `{"address":`},
		{"there is no address", `{"name":"Ann"}`},
		{"the address is a string", `{"address":"somewhere"}`},
		{"the address is null", `{"address":null}`},
		{"there is no city", `{"address":{"zip":"0150"}}`},
		{"the city is a number", `{"address":{"city":5}}`},
	}
	for _, tt := range tests {
		got, err := LookupCity(tt.data)
		if err == nil {
			t.Errorf("LookupCity(%s) returned a nil error, want an error because %s", tt.data, tt.why)
		}
		if got != "" {
			t.Errorf("LookupCity(%s) = %q together with an error, want an empty string", tt.data, got)
		}
	}
}

func TestWriteJSON(t *testing.T) {
	tests := []struct {
		v    any
		want string
	}{
		{map[string]int{"apple": 5, "lettuce": 7}, `{"apple":5,"lettuce":7}` + "\n"},
		{[]int{1, 2, 3}, "[1,2,3]\n"},
		{"gopher", `"gopher"` + "\n"},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		if err := WriteJSON(&buf, tt.v); err != nil {
			t.Errorf("WriteJSON(w, %#v) returned error %v, want nil", tt.v, err)
		}
		if got := buf.String(); got != tt.want {
			t.Errorf("WriteJSON(w, %#v) wrote %q, want %q", tt.v, got, tt.want)
		}
	}

	var buf bytes.Buffer
	if err := WriteJSON(&buf, func() {}); err == nil {
		t.Errorf("WriteJSON(w, a function) returned nil, want an encoding error")
	}
	if err := WriteJSON(brokenWriter{}, 1); err == nil {
		t.Errorf("WriteJSON(a failing writer, 1) returned nil, want the write error")
	}
}

func TestFilterAdults(t *testing.T) {
	input := `{"name":"Ann","age":30,"email":"ann@example.com","roles":["admin"],"address":{"city":"Oslo","country":"NO"}}
{"name":"Bob","age":12,"roles":["kid"],"address":{"city":"Lima","country":"PE"}}
{"name":"Cy","age":18,"roles":[],"address":{"city":"Rome","country":"IT"}}

{"name":"Di","age":17,"roles":[],"address":{"city":"Bern","country":"CH"}} {"name":"Ed","age":64,"roles":["dev","ops"],"address":{"city":"Kyiv","country":"UA"}}
`
	want := []User{
		{"Ann", 30, "ann@example.com", []string{"admin"}, Address{"Oslo", "NO"}},
		{"Cy", 18, "", []string{}, Address{"Rome", "IT"}},
		{"Ed", 64, "", []string{"dev", "ops"}, Address{"Kyiv", "UA"}},
	}

	var out bytes.Buffer
	n, err := FilterAdults(strings.NewReader(input), &out)
	if err != nil {
		t.Fatalf("FilterAdults returned error %v, want nil", err)
	}
	if n != len(want) {
		t.Errorf("FilterAdults returned %d, want %d users written", n, len(want))
	}

	// One JSON object per line...
	lines := strings.Split(strings.TrimSuffix(out.String(), "\n"), "\n")
	if len(lines) != len(want) {
		t.Fatalf("FilterAdults wrote %d lines, want %d. Output:\n%s", len(lines), len(want), out.String())
	}
	// ...that decode to the adults, in input order.
	for i, line := range lines {
		var got User
		if err := json.Unmarshal([]byte(line), &got); err != nil {
			t.Fatalf("output line %d (%q) is not valid JSON: %v", i+1, line, err)
		}
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("output line %d decodes to %+v, want %+v", i+1, got, want[i])
		}
	}
	if !strings.Contains(lines[0], `"name":"Ann"`) {
		t.Errorf("output line 1 is %s, want the tagged field names such as \"name\"", lines[0])
	}

	// Reaching the end of the input is not an error, even right away.
	out.Reset()
	n, err = FilterAdults(strings.NewReader(""), &out)
	if n != 0 || err != nil {
		t.Errorf("FilterAdults(empty input) = %d, %v; want 0, nil", n, err)
	}
	if out.Len() != 0 {
		t.Errorf("FilterAdults(empty input) wrote %q, want nothing", out.String())
	}
}

func TestFilterAdultsBadInput(t *testing.T) {
	input := `{"name":"Ann","age":30}
{"name":"Bob","age":"old"}
{"name":"Cy","age":40}
`
	var out bytes.Buffer
	n, err := FilterAdults(strings.NewReader(input), &out)
	if err == nil {
		t.Fatalf("FilterAdults returned a nil error for input with a broken record, want an error")
	}
	if n != 1 {
		t.Errorf("FilterAdults returned %d, want 1 (one user was written before the broken record)", n)
	}
	if got := out.String(); !strings.Contains(got, "Ann") || strings.Contains(got, "Cy") {
		t.Errorf("FilterAdults wrote %q, want only Ann (stop at the broken record)", got)
	}
}
