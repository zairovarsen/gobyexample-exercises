package helloworld

import "testing"

func TestGreeting(t *testing.T) {
	if got, want := Greeting(), "hello world"; got != want {
		t.Errorf("Greeting() = %q, want %q", got, want)
	}
}

func TestGreetingFor(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"gopher", "hello gopher"},
		{"Go by Example", "hello Go by Example"},
		{"", "hello world"},
	}
	for _, tt := range tests {
		if got := GreetingFor(tt.name); got != tt.want {
			t.Errorf("GreetingFor(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}

// An Example function is a test too: `go test` runs it and compares what it
// printed to standard output with the "Output:" comment at the end.
func ExamplePrintGreeting() {
	PrintGreeting()
	// Output:
	// hello world
}
