package xmlx

import (
	"maps"
	"slices"
	"testing"
)

const coffeeXML = `<plant id="27">
  <name>Coffee</name>
  <origin>Ethiopia</origin>
  <origin>Brazil</origin>
</plant>`

const gardenXML = `<garden name="Kew">
  <plants>
    <plant id="27">
      <name>Coffee</name>
      <origin>Ethiopia</origin>
      <origin>Brazil</origin>
    </plant>
    <plant id="81">
      <name>Tomato</name>
      <origin>Mexico</origin>
      <origin>Brazil</origin>
    </plant>
  </plants>
</garden>`

func coffee() Plant {
	return Plant{ID: 27, Name: "Coffee", Origin: []string{"Ethiopia", "Brazil"}}
}

func tomato() Plant {
	return Plant{ID: 81, Name: "Tomato", Origin: []string{"Mexico", "Brazil"}}
}

func TestEncodePlant(t *testing.T) {
	got, err := EncodePlant(coffee())
	if err != nil {
		t.Fatalf("EncodePlant(coffee) returned error %v, want nil", err)
	}
	if got != coffeeXML {
		t.Errorf("EncodePlant(coffee) =\n%s\nwant\n%s", got, coffeeXML)
	}

	got, err = EncodePlant(Plant{ID: 5, Name: "Moss"})
	if err != nil {
		t.Fatalf("EncodePlant(moss) returned error %v, want nil", err)
	}
	if want := "<plant id=\"5\">\n  <name>Moss</name>\n</plant>"; got != want {
		t.Errorf("EncodePlant(moss) =\n%s\nwant\n%s", got, want)
	}
}

func TestEncodeDocument(t *testing.T) {
	got, err := EncodeDocument(coffee())
	if err != nil {
		t.Fatalf("EncodeDocument(coffee) returned error %v, want nil", err)
	}
	want := `<?xml version="1.0" encoding="UTF-8"?>` + "\n" + coffeeXML
	if got != want {
		t.Errorf("EncodeDocument(coffee) =\n%s\nwant\n%s", got, want)
	}
}

func TestDecodePlant(t *testing.T) {
	tests := []struct {
		data string
		want Plant
	}{
		{coffeeXML, coffee()},
		{`<plant id="81"><name>Tomato</name><origin>Mexico</origin><origin>Brazil</origin></plant>`, tomato()},
		{`<?xml version="1.0" encoding="UTF-8"?>` + "\n" + `<plant id="5"><name>Moss</name></plant>`, Plant{ID: 5, Name: "Moss"}},
	}
	for _, tt := range tests {
		got, err := DecodePlant(tt.data)
		if err != nil {
			t.Errorf("DecodePlant(%s) returned error %v, want nil", tt.data, err)
			continue
		}
		// XMLName is filled in by the decoder; compare the other fields.
		if got.ID != tt.want.ID || got.Name != tt.want.Name || !slices.Equal(got.Origin, tt.want.Origin) {
			t.Errorf("DecodePlant(%s) =\n  {ID:%d Name:%q Origin:%q}, want\n  {ID:%d Name:%q Origin:%q}",
				tt.data, got.ID, got.Name, got.Origin, tt.want.ID, tt.want.Name, tt.want.Origin)
		}
	}
}

func TestDecodePlantErrors(t *testing.T) {
	tests := []struct {
		why  string
		data string
	}{
		{"the <name> element is never closed", `<plant id="27"><name>Coffee</plant>`},
		{"the id attribute is not a number", `<plant id="many"><name>Coffee</name></plant>`},
		{"the root element is <tree>, not <plant>", `<tree id="27"><name>Oak</name></tree>`},
		{"it is not XML", `plant: coffee`},
	}
	for _, tt := range tests {
		if _, err := DecodePlant(tt.data); err == nil {
			t.Errorf("DecodePlant(%s) returned a nil error, want an error because %s", tt.data, tt.why)
		}
	}
}

func TestEncodeGarden(t *testing.T) {
	g := Garden{Name: "Kew", Plants: []Plant{coffee(), tomato()}}
	got, err := EncodeGarden(g)
	if err != nil {
		t.Fatalf("EncodeGarden(kew) returned error %v, want nil", err)
	}
	if got != gardenXML {
		t.Errorf("EncodeGarden(kew) =\n%s\nwant\n%s", got, gardenXML)
	}
}

func TestOriginCounts(t *testing.T) {
	want := map[string]int{"Brazil": 2, "Ethiopia": 1, "Mexico": 1}
	got, err := OriginCounts(gardenXML)
	if err != nil {
		t.Fatalf("OriginCounts(kew garden) returned error %v, want nil", err)
	}
	if !maps.Equal(got, want) {
		t.Errorf("OriginCounts(kew garden) = %v, want %v", got, want)
	}

	oneLine := `<garden name="Tiny"><plants><plant id="1"><name>Fern</name><origin>Peru</origin></plant></plants></garden>`
	want = map[string]int{"Peru": 1}
	got, err = OriginCounts(oneLine)
	if err != nil {
		t.Fatalf("OriginCounts(%s) returned error %v, want nil", oneLine, err)
	}
	if !maps.Equal(got, want) {
		t.Errorf("OriginCounts(%s) = %v, want %v", oneLine, got, want)
	}

	broken := `<garden name="Kew"><plants><plant id="27"></garden>`
	got, err = OriginCounts(broken)
	if err == nil {
		t.Errorf("OriginCounts(%s) returned a nil error, want an error for malformed XML", broken)
	}
	if got != nil {
		t.Errorf("OriginCounts(%s) = %v together with an error, want a nil map", broken, got)
	}
}
