//go:build !solution

// Package xmlx contains exercises for https://gobyexample.com/xml
//
// Run the tests for this section with:
//
//	go test ./56-xml/
package xmlx

import "encoding/xml"

// Plant is the type from the gobyexample page. Encoded as XML it has to
// look like this:
//
//	<plant id="27">
//	  <name>Coffee</name>
//	  <origin>Ethiopia</origin>
//	  <origin>Brazil</origin>
//	</plant>
//
// TODO: add struct tags to all four fields to get there.
//
//   - The tag on the special XMLName field sets the name of the element that
//     stands for the struct itself: `xml:"plant"`.
//   - `xml:"id,attr"` makes a field an attribute of that element instead of
//     a nested element.
//   - `xml:"name"` gives a nested element its name; a slice produces one
//     element per item.
type Plant struct {
	XMLName xml.Name
	ID      int
	Name    string
	Origin  []string
}

// Garden holds several plants. Encoded as XML it has to look like this:
//
//	<garden name="Kew">
//	  <plants>
//	    <plant id="27">
//	      ...
//	    </plant>
//	    <plant id="81">
//	      ...
//	    </plant>
//	  </plants>
//	</garden>
//
// TODO: add struct tags. The plants are nested inside a <plants> element
// that has no struct of its own: a tag of the form `xml:"parent>child"`
// creates such in-between elements.
type Garden struct {
	XMLName xml.Name
	Name    string
	Plants  []Plant
}

// EncodePlant returns p as human-readable XML, exactly as shown in the
// comment on Plant: nested elements are indented by two spaces per level
// and there is no trailing newline.
//
// Use xml.MarshalIndent(p, "", "  "); the second argument is a prefix for
// every line (none here) and the third the indentation.
func EncodePlant(p Plant) (string, error) {
	// TODO: marshal with indentation and convert the bytes to a string.
	return "", nil
}

// EncodeDocument returns p as a complete XML document: the standard XML
// declaration line followed by the output of EncodePlant.
//
//	<?xml version="1.0" encoding="UTF-8"?>
//	<plant id="27">
//	  ...
//	</plant>
//
// Hint: the declaration, including its newline, is the constant xml.Header.
func EncodeDocument(p Plant) (string, error) {
	// TODO: put the header in front of the encoded plant.
	return "", nil
}

// DecodePlant parses the XML in data into a Plant with xml.Unmarshal. It
// returns the error from Unmarshal when data is malformed or when its root
// element is not <plant>.
func DecodePlant(data string) (Plant, error) {
	// TODO: unmarshal into a Plant.
	return Plant{}, nil
}

// EncodeGarden returns g as indented XML in the same style as EncodePlant
// (two spaces per level, no trailing newline), in the shape shown in the
// comment on Garden.
func EncodeGarden(g Garden) (string, error) {
	// TODO: marshal with indentation.
	return "", nil
}

// OriginCounts parses data, a <garden> document, and counts for every
// origin how many plants of the garden list it.
//
// Example: for a garden with Coffee (origins Ethiopia, Brazil) and Tomato
// (origins Mexico, Brazil) it returns
// map[Brazil:2 Ethiopia:1 Mexico:1].
//
// Malformed XML gives a nil map and the error from xml.Unmarshal.
func OriginCounts(data string) (map[string]int, error) {
	// TODO: unmarshal into a Garden, then count with a map.
	return nil, nil
}
