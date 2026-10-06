//go:build solution

package base64encoding

import (
	b64 "encoding/base64"
	"errors"
	"strings"
)

func Encode(s string) string {
	return b64.StdEncoding.EncodeToString([]byte(s))
}

func Decode(encoded string) (string, error) {
	data, err := b64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func EncodeURL(s string) string {
	return b64.URLEncoding.EncodeToString([]byte(s))
}

func DecodeURL(encoded string) (string, error) {
	data, err := b64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ToURLSafe(std string) (string, error) {
	data, err := b64.StdEncoding.DecodeString(std)
	if err != nil {
		return "", err
	}
	return b64.URLEncoding.EncodeToString(data), nil
}

func BasicAuth(user, password string) string {
	return "Basic " + Encode(user+":"+password)
}

func ParseBasicAuth(header string) (user, password string, err error) {
	encoded, ok := strings.CutPrefix(header, "Basic ")
	if !ok {
		return "", "", errors.New(`header does not start with "Basic "`)
	}
	decoded, err := Decode(encoded)
	if err != nil {
		return "", "", err
	}
	user, password, ok = strings.Cut(decoded, ":")
	if !ok {
		return "", "", errors.New(`credentials contain no ":"`)
	}
	return user, password, nil
}
