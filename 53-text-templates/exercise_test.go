package texttemplates

import (
	"bytes"
	"errors"
	"testing"
)

// brokenWriter is an io.Writer whose every Write fails.
type brokenWriter struct{}

var errBroken = errors.New("broken writer")

func (brokenWriter) Write([]byte) (int, error) { return 0, errBroken }

func TestRenderValue(t *testing.T) {
	tests := []struct {
		value any
		want  string
	}{
		{42, "Value is 42\n"},
		{"some text", "Value is some text\n"},
		{[]string{"Go", "Rust", "C++", "C#"}, "Value is [Go Rust C++ C#]\n"},
		{true, "Value is true\n"},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		if err := RenderValue(&buf, tt.value); err != nil {
			t.Errorf("RenderValue(w, %#v) returned error %v, want nil", tt.value, err)
		}
		if got := buf.String(); got != tt.want {
			t.Errorf("RenderValue(w, %#v) wrote %q, want %q", tt.value, got, tt.want)
		}
	}

	if err := RenderValue(brokenWriter{}, 1); err == nil {
		t.Errorf("RenderValue(a failing writer, 1) returned nil, want the error from Execute")
	}
}

func TestRender(t *testing.T) {
	tests := []struct {
		text string
		data any
		want string
	}{
		{"Hi {{.Name}}!", Person{Name: "Ann"}, "Hi Ann!"},
		{"Hi {{.name}}!", map[string]string{"name": "Ann"}, "Hi Ann!"},
		{"{{.Age}}/{{.Name}}/{{.Age}}", Person{"Bo", 7}, "7/Bo/7"},
		{"{{.}} + {{.}}", 1, "1 + 1"},
		{"no actions at all", nil, "no actions at all"},
	}
	for _, tt := range tests {
		got, err := Render(tt.text, tt.data)
		if err != nil {
			t.Errorf("Render(%q, %#v) returned error %v, want nil", tt.text, tt.data, err)
		}
		if got != tt.want {
			t.Errorf("Render(%q, %#v) = %q, want %q", tt.text, tt.data, got, tt.want)
		}
	}
}

func TestRenderErrors(t *testing.T) {
	tests := []struct {
		why  string
		text string
		data any
	}{
		{"the action is not closed", "Hi {{.Name", Person{Name: "Ann"}},
		{"{{end}} has nothing to end", "Hi {{end}}", Person{Name: "Ann"}},
		{"Person has no field Email", "Hi {{.Email}}", Person{Name: "Ann"}},
	}
	for _, tt := range tests {
		got, err := Render(tt.text, tt.data)
		if err == nil {
			t.Errorf("Render(%q, %#v) returned a nil error, want an error because %s", tt.text, tt.data, tt.why)
		}
		if got != "" {
			t.Errorf("Render(%q, %#v) = %q together with an error, want an empty string", tt.text, tt.data, got)
		}
	}
}

func TestNameTag(t *testing.T) {
	tests := []struct {
		p    Person
		want string
	}{
		{Person{"Jane Doe", 30}, "Hello, my name is Jane Doe and I am 30 years old."},
		{Person{"Mickey Mouse", 96}, "Hello, my name is Mickey Mouse and I am 96 years old."},
	}
	for _, tt := range tests {
		got, err := NameTag(tt.p)
		if err != nil {
			t.Errorf("NameTag(%+v) returned error %v, want nil", tt.p, err)
		}
		if got != tt.want {
			t.Errorf("NameTag(%+v) = %q, want %q", tt.p, got, tt.want)
		}
	}
}

func TestYesNo(t *testing.T) {
	var nilPerson *Person
	tests := []struct {
		v    any
		want string
	}{
		{"not empty", "yes"},
		{"", "no"},
		{true, "yes"},
		{false, "no"},
		{7, "yes"},
		{0, "no"},
		{[]string{"x"}, "yes"},
		{[]string{}, "no"},
		{map[string]int{}, "no"},
		{&Person{}, "yes"},
		{nilPerson, "no"},
	}
	for _, tt := range tests {
		got, err := YesNo(tt.v)
		if err != nil {
			t.Errorf("YesNo(%#v) returned error %v, want nil", tt.v, err)
		}
		if got != tt.want {
			t.Errorf("YesNo(%#v) = %q, want %q", tt.v, got, tt.want)
		}
	}
}

func TestBulletList(t *testing.T) {
	tests := []struct {
		items []string
		want  string
	}{
		{[]string{"Go", "Rust"}, "- Go\n- Rust\n"},
		{[]string{"Go", "Rust", "C++", "C#"}, "- Go\n- Rust\n- C++\n- C#\n"},
		{[]string{"only one"}, "- only one\n"},
		{nil, ""},
	}
	for _, tt := range tests {
		got, err := BulletList(tt.items)
		if err != nil {
			t.Errorf("BulletList(%q) returned error %v, want nil", tt.items, err)
		}
		if got != tt.want {
			t.Errorf("BulletList(%q) = %q, want %q", tt.items, got, tt.want)
		}
	}
}

func TestWriteOrder(t *testing.T) {
	tests := []struct {
		order Order
		want  string
	}{
		{
			Order{"Jane", []OrderItem{{"apple", 2}, {"pear", 1}}, true},
			"Order for Jane\n- 2 x apple\n- 1 x pear\nStatus: paid\n",
		},
		{
			Order{"Bob", []OrderItem{{"coffee", 12}}, false},
			"Order for Bob\n- 12 x coffee\nStatus: unpaid\n",
		},
		{
			Order{Customer: "Nobody"},
			"Order for Nobody\nStatus: unpaid\n",
		},
	}
	for _, tt := range tests {
		var buf bytes.Buffer
		if err := WriteOrder(&buf, tt.order); err != nil {
			t.Errorf("WriteOrder(w, %+v) returned error %v, want nil", tt.order, err)
		}
		if got := buf.String(); got != tt.want {
			t.Errorf("WriteOrder(w, %+v) wrote\n%q, want\n%q", tt.order, got, tt.want)
		}
	}

	if err := WriteOrder(brokenWriter{}, tests[0].order); err == nil {
		t.Errorf("WriteOrder(a failing writer, ...) returned nil, want the error from Execute")
	}
}
