package structs

import (
	"reflect"
	"testing"
)

// personField returns the named field of the Person that p points to. The
// test looks the field up at run time (with reflection) rather than writing
// p.Name, so that this file compiles before the fields of Person exist.
func personField(t *testing.T, p *Person, name string, kind reflect.Kind) reflect.Value {
	t.Helper()
	if p == nil {
		t.Fatal("got a nil *Person")
	}
	f := reflect.ValueOf(p).Elem().FieldByName(name)
	if !f.IsValid() {
		t.Fatalf("Person has no field called %s - add it to the struct declaration", name)
	}
	if f.Kind() != kind {
		t.Fatalf("field %s of Person has type %s, want %s", name, f.Type(), kind)
	}
	return f
}

func TestPersonFields(t *testing.T) {
	p := &Person{}
	if got := personField(t, p, "Name", reflect.String).String(); got != "" {
		t.Errorf("Name of a zero Person = %q, want %q", got, "")
	}
	if got := personField(t, p, "Age", reflect.Int).Int(); got != 0 {
		t.Errorf("Age of a zero Person = %d, want 0", got)
	}
	if n := reflect.TypeFor[Person]().NumField(); n != 2 {
		t.Errorf("Person has %d fields, want exactly 2 (Name and Age)", n)
	}
}

func TestNewPerson(t *testing.T) {
	for _, name := range []string{"Alice", "Bob"} {
		p := NewPerson(name)
		if p == nil {
			t.Fatalf("NewPerson(%q) = nil, want a pointer to a Person", name)
		}
		if got := personField(t, p, "Name", reflect.String).String(); got != name {
			t.Errorf("NewPerson(%q) has Name %q, want %q", name, got, name)
		}
		if got := personField(t, p, "Age", reflect.Int).Int(); got != 42 {
			t.Errorf("NewPerson(%q) has Age %d, want 42", name, got)
		}
	}

	if a, b := NewPerson("Jon"), NewPerson("Jon"); a == b {
		t.Error("two calls of NewPerson returned the same pointer, want two different Persons")
	}
}

func TestTranslate(t *testing.T) {
	tests := []struct {
		p      Point
		dx, dy int
		want   Point
	}{
		{Point{1, 2}, 10, 20, Point{11, 22}},
		{Point{0, 0}, -3, 4, Point{-3, 4}},
		{Point{5, 5}, 0, -5, Point{5, 0}},
	}
	for _, tt := range tests {
		// Structs with comparable fields can be compared with == and !=.
		if got := Translate(tt.p, tt.dx, tt.dy); got != tt.want {
			t.Errorf("Translate(%+v, %d, %d) = %+v, want %+v", tt.p, tt.dx, tt.dy, got, tt.want)
		}
	}
}

func TestMoveTo(t *testing.T) {
	p := Point{X: 1, Y: 2}
	MoveTo(&p, 7, -8)
	if want := (Point{X: 7, Y: -8}); p != want {
		t.Errorf("p := Point{1, 2}; MoveTo(&p, 7, -8) left p = %+v, want %+v", p, want)
	}

	sp := &Point{}
	MoveTo(sp, 3, 0)
	if want := (Point{X: 3, Y: 0}); *sp != want {
		t.Errorf("sp := &Point{}; MoveTo(sp, 3, 0) left *sp = %+v, want %+v", *sp, want)
	}
}

func TestRexTheDog(t *testing.T) {
	dog := RexTheDog()
	if dog.Name != "Rex" || !dog.IsGood {
		t.Errorf("RexTheDog() = %+v, want {Name:Rex IsGood:true}", dog)
	}
}

func TestBoundingBox(t *testing.T) {
	tests := []struct {
		points []Point
		want   Rect
	}{
		{
			[]Point{{1, 5}, {4, 2}, {3, 9}},
			Rect{Min: Point{1, 2}, Max: Point{4, 9}},
		},
		{
			[]Point{{-2, -3}},
			Rect{Min: Point{-2, -3}, Max: Point{-2, -3}},
		},
		{
			[]Point{{5, 5}, {-1, 7}, {2, -6}, {0, 0}},
			Rect{Min: Point{-1, -6}, Max: Point{5, 7}},
		},
		{
			// All coordinates are positive: the box must not start at (0, 0).
			[]Point{{3, 4}, {6, 8}},
			Rect{Min: Point{3, 4}, Max: Point{6, 8}},
		},
		{nil, Rect{}},
	}
	for _, tt := range tests {
		if got := BoundingBox(tt.points); got != tt.want {
			t.Errorf("BoundingBox(%v) = %+v, want %+v", tt.points, got, tt.want)
		}
	}
}
