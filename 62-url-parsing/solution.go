//go:build solution

package urlparsing

import (
	"net"
	"net/url"
	"strconv"
	"strings"
)

type DBConfig struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Options  map[string]string
}

func Scheme(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return u.Scheme, nil
}

func Credentials(rawURL string) (user, password string, err error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", "", err
	}
	// Username and Password are safe to call on a nil *Userinfo.
	password, _ = u.User.Password()
	return u.User.Username(), password, nil
}

func HostPort(rawURL string) (host, port string, err error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", "", err
	}
	if u.Port() == "" {
		return u.Host, "", nil
	}
	return net.SplitHostPort(u.Host)
}

func PathAndFragment(rawURL string) (path, fragment string, err error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", "", err
	}
	return u.Path, u.Fragment, nil
}

func QueryValues(rawURL, key string) ([]string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return nil, err
	}
	return m[key], nil
}

func ParseDBConfig(rawURL string) (DBConfig, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return DBConfig{}, err
	}
	host, portText, err := net.SplitHostPort(u.Host)
	if err != nil {
		return DBConfig{}, err
	}
	port, err := strconv.Atoi(portText)
	if err != nil {
		return DBConfig{}, err
	}
	query, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return DBConfig{}, err
	}
	options := make(map[string]string, len(query))
	for key, values := range query {
		options[key] = values[0]
	}
	password, _ := u.User.Password()
	return DBConfig{
		Driver:   u.Scheme,
		User:     u.User.Username(),
		Password: password,
		Host:     host,
		Port:     port,
		Database: strings.TrimPrefix(u.Path, "/"),
		Options:  options,
	}, nil
}
