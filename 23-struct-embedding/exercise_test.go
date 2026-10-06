package structembedding

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

// The tests cannot write c.Num or c.Describe() directly: that would not
// compile before Container embeds Base. They look fields and methods up at
// run time instead, which finds promoted ones too.

// requireEmbedded fails the test unless the struct v embeds a type called name.
func requireEmbedded(t *testing.T, v any, name string) {
	t.Helper()
	f, ok := reflect.TypeOf(v).FieldByName(name)
	if !ok || !f.Anonymous {
		t.Fatalf("%T does not embed %s - add it to the struct as a field without a name", v, name)
	}
}

// intField returns the value of the (possibly promoted) int field name of v.
func intField(t *testing.T, v any, name string) int {
	t.Helper()
	f := reflect.ValueOf(v).FieldByName(name)
	if !f.IsValid() {
		t.Fatalf("%T has no field %s (not even a promoted one)", v, name)
	}
	return int(f.Int())
}

type fake string

func (f fake) Describe() string { return "fake " + string(f) }

func TestBaseDescribe(t *testing.T) {
	tests := []struct {
		base Base
		want string
	}{
		{Base{Num: 1}, "base with num=1"},
		{Base{Num: -20}, "base with num=-20"},
		{Base{}, "base with num=0"},
	}
	for _, tt := range tests {
		if got := tt.base.Describe(); got != tt.want {
			t.Errorf("%#v.Describe() = %q, want %q", tt.base, got, tt.want)
		}
	}
}

func TestNewContainer(t *testing.T) {
	c := NewContainer(1, "some name")
	requireEmbedded(t, c, "Base")

	if got := intField(t, c, "Num"); got != 1 {
		t.Errorf(`NewContainer(1, "some name"): c.Num = %d, want 1`, got)
	}
	if c.Str != "some name" {
		t.Errorf(`NewContainer(1, "some name"): c.Str = %q, want "some name"`, c.Str)
	}

	// The method of Base is promoted to Container.
	d, ok := any(c).(Describer)
	if !ok {
		t.Fatal("Container has no Describe method - it should be promoted from the embedded Base")
	}
	if got, want := d.Describe(), "base with num=1"; got != want {
		t.Errorf(`NewContainer(1, "some name"): c.Describe() = %q, want %q`, got, want)
	}
}

func TestAsDescriber(t *testing.T) {
	d := AsDescriber(NewContainer(7, "seven"))
	if d == nil {
		t.Fatal(`AsDescriber(NewContainer(7, "seven")) = nil, want the container`)
	}
	if got, want := d.Describe(), "base with num=7"; got != want {
		t.Errorf("AsDescriber(...).Describe() = %q, want %q", got, want)
	}
	if c, ok := any(d).(Container); !ok || c.Str != "seven" {
		t.Errorf("AsDescriber returned %#v, want the Container it was given", d)
	}
}

func TestDescribeAll(t *testing.T) {
	got := DescribeAll(Base{Num: 1}, fake("thing"), Base{Num: 3})
	want := []string{"base with num=1", "fake thing", "base with num=3"}
	if !slices.Equal(got, want) {
		t.Errorf(`DescribeAll(Base{1}, fake("thing"), Base{3}) = %q, want %q`, got, want)
	}

	if got := DescribeAll(); len(got) != 0 {
		t.Errorf("DescribeAll() = %q, want an empty result", got)
	}
}

func TestPoint(t *testing.T) {
	p := Point{X: 1, Y: 2}
	if got, want := p.String(), "(1,2)"; got != want {
		t.Errorf("Point{1, 2}.String() = %q, want %q", got, want)
	}

	p.Move(2, -5)
	if want := (Point{X: 3, Y: -3}); p != want {
		t.Errorf("Point{1, 2} after Move(2, -5) = %#v, want %#v", p, want)
	}
	if got, want := fmt.Sprint(p), "(3,-3)"; got != want {
		t.Errorf("fmt.Sprint(Point{3, -3}) = %q, want %q", got, want)
	}
}

func TestCircle(t *testing.T) {
	c := NewCircle(1, 2, 3)
	requireEmbedded(t, c, "Point")

	if x, y := intField(t, c, "X"), intField(t, c, "Y"); x != 1 || y != 2 || c.Radius != 3 {
		t.Errorf("NewCircle(1, 2, 3): X, Y, Radius = %d, %d, %d, want 1, 2, 3", x, y, c.Radius)
	}

	// Move is promoted from *Point, so it is callable on an addressable Circle.
	mover, ok := any(&c).(interface{ Move(dx, dy int) })
	if !ok {
		t.Fatal("*Circle has no Move method - it should be promoted from the embedded Point")
	}
	mover.Move(10, 20)
	if x, y := intField(t, c, "X"), intField(t, c, "Y"); x != 11 || y != 22 {
		t.Errorf("NewCircle(1, 2, 3) after Move(10, 20): X, Y = %d, %d, want 11, 22", x, y)
	}

	// Circle's own String method wins over the one promoted from Point.
	if got, want := fmt.Sprint(c), "circle at (11,22) r=3"; got != want {
		t.Errorf("fmt.Sprint(circle) = %q, want %q", got, want)
	}
	if got, want := c.String(), "circle at (11,22) r=3"; got != want {
		t.Errorf("circle.String() = %q, want %q", got, want)
	}
}
