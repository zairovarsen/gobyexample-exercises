//go:build solution

package xmlx

import "encoding/xml"

type Plant struct {
	XMLName xml.Name `xml:"plant"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

type Garden struct {
	XMLName xml.Name `xml:"garden"`
	Name    string   `xml:"name,attr"`
	Plants  []Plant  `xml:"plants>plant"`
}

func EncodePlant(p Plant) (string, error) {
	out, err := xml.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func EncodeDocument(p Plant) (string, error) {
	body, err := EncodePlant(p)
	if err != nil {
		return "", err
	}
	return xml.Header + body, nil
}

func DecodePlant(data string) (Plant, error) {
	var p Plant
	if err := xml.Unmarshal([]byte(data), &p); err != nil {
		return Plant{}, err
	}
	return p, nil
}

func EncodeGarden(g Garden) (string, error) {
	out, err := xml.MarshalIndent(g, "", "  ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func OriginCounts(data string) (map[string]int, error) {
	var g Garden
	if err := xml.Unmarshal([]byte(data), &g); err != nil {
		return nil, err
	}
	counts := make(map[string]int)
	for _, plant := range g.Plants {
		for _, origin := range plant.Origin {
			counts[origin]++
		}
	}
	return counts, nil
}
