//go:build solution

package jsonx

import (
	"encoding/json"
	"errors"
	"io"
)

type Address struct {
	City    string `json:"city"`
	Country string `json:"country"`
}

type User struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email,omitempty"`
	Roles   []string `json:"roles"`
	Address Address  `json:"address"`
}

func ToJSON(v any) (string, error) {
	data, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func DecodeUser(data string) (User, error) {
	var u User
	if err := json.Unmarshal([]byte(data), &u); err != nil {
		return User{}, err
	}
	return u, nil
}

func SumNumbers(data string) (float64, error) {
	var values map[string]any
	if err := json.Unmarshal([]byte(data), &values); err != nil {
		return 0, err
	}
	sum := 0.0
	for _, v := range values {
		if n, ok := v.(float64); ok {
			sum += n
		}
	}
	return sum, nil
}

func LookupCity(data string) (string, error) {
	var doc map[string]any
	if err := json.Unmarshal([]byte(data), &doc); err != nil {
		return "", err
	}
	address, ok := doc["address"].(map[string]any)
	if !ok {
		return "", errors.New("address is missing or not an object")
	}
	city, ok := address["city"].(string)
	if !ok {
		return "", errors.New("city is missing or not a string")
	}
	return city, nil
}

func WriteJSON(w io.Writer, v any) error {
	return json.NewEncoder(w).Encode(v)
}

func FilterAdults(r io.Reader, w io.Writer) (int, error) {
	dec := json.NewDecoder(r)
	enc := json.NewEncoder(w)
	written := 0
	for {
		var u User
		err := dec.Decode(&u)
		if err == io.EOF {
			return written, nil
		}
		if err != nil {
			return written, err
		}
		if u.Age < 18 {
			continue
		}
		if err := enc.Encode(u); err != nil {
			return written, err
		}
		written++
	}
}
