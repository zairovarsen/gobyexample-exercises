package urlparsing

import (
	"maps"
	"slices"
	"testing"
)

// The example from the gobyexample page.
const pgURL = "postgres://user:pass@host.com:5432/path?k=v#f"

// Strings that url.Parse rejects.
var badURLs = []string{
	"://missing-scheme",
	"http://host:port/",
	"http://[::1",
	"http://exa mple.com/",
}

func TestScheme(t *testing.T) {
	tests := []struct {
		rawURL string
		want   string
	}{
		{pgURL, "postgres"},
		{"https://gobyexample.com/url-parsing", "https"},
		{"mailto:gopher@example.com", "mailto"},
		{"HTTP://EXAMPLE.COM", "http"},
		{"/just/a/path", ""},
	}
	for _, tt := range tests {
		got, err := Scheme(tt.rawURL)
		if err != nil {
			t.Errorf("Scheme(%q) returned error %v, want nil", tt.rawURL, err)
		}
		if got != tt.want {
			t.Errorf("Scheme(%q) = %q, want %q", tt.rawURL, got, tt.want)
		}
	}
	for _, rawURL := range badURLs {
		if _, err := Scheme(rawURL); err == nil {
			t.Errorf("Scheme(%q) returned a nil error, want the error from url.Parse", rawURL)
		}
	}
}

func TestCredentials(t *testing.T) {
	tests := []struct {
		rawURL             string
		wantUser, wantPass string
	}{
		{pgURL, "user", "pass"},
		{"ftp://anonymous@files.example.com/pub", "anonymous", ""},
		{"https://gobyexample.com/url-parsing", "", ""},
		{"https://ann:p%40ss%3Aword@example.com", "ann", "p@ss:word"},
	}
	for _, tt := range tests {
		user, pass, err := Credentials(tt.rawURL)
		if err != nil {
			t.Errorf("Credentials(%q) returned error %v, want nil", tt.rawURL, err)
		}
		if user != tt.wantUser || pass != tt.wantPass {
			t.Errorf("Credentials(%q) = %q, %q; want %q, %q", tt.rawURL, user, pass, tt.wantUser, tt.wantPass)
		}
	}
	for _, rawURL := range badURLs {
		if _, _, err := Credentials(rawURL); err == nil {
			t.Errorf("Credentials(%q) returned a nil error, want the error from url.Parse", rawURL)
		}
	}
}

func TestHostPort(t *testing.T) {
	tests := []struct {
		rawURL             string
		wantHost, wantPort string
	}{
		{pgURL, "host.com", "5432"},
		{"https://gobyexample.com/url-parsing", "gobyexample.com", ""},
		{"http://localhost:8080", "localhost", "8080"},
		{"http://127.0.0.1:80/index.html", "127.0.0.1", "80"},
		{"http://[::1]:8080/", "::1", "8080"},
	}
	for _, tt := range tests {
		host, port, err := HostPort(tt.rawURL)
		if err != nil {
			t.Errorf("HostPort(%q) returned error %v, want nil", tt.rawURL, err)
		}
		if host != tt.wantHost || port != tt.wantPort {
			t.Errorf("HostPort(%q) = %q, %q; want %q, %q", tt.rawURL, host, port, tt.wantHost, tt.wantPort)
		}
	}
	for _, rawURL := range badURLs {
		if _, _, err := HostPort(rawURL); err == nil {
			t.Errorf("HostPort(%q) returned a nil error, want the error from url.Parse", rawURL)
		}
	}
}

func TestPathAndFragment(t *testing.T) {
	tests := []struct {
		rawURL                 string
		wantPath, wantFragment string
	}{
		{pgURL, "/path", "f"},
		{"https://gobyexample.com/url-parsing", "/url-parsing", ""},
		{"https://go.dev/doc/effective_go#names", "/doc/effective_go", "names"},
		{"https://example.com/a%20b/c?x=1#top", "/a b/c", "top"},
		{"https://example.com", "", ""},
	}
	for _, tt := range tests {
		path, fragment, err := PathAndFragment(tt.rawURL)
		if err != nil {
			t.Errorf("PathAndFragment(%q) returned error %v, want nil", tt.rawURL, err)
		}
		if path != tt.wantPath || fragment != tt.wantFragment {
			t.Errorf("PathAndFragment(%q) = %q, %q; want %q, %q", tt.rawURL, path, fragment, tt.wantPath, tt.wantFragment)
		}
	}
	for _, rawURL := range badURLs {
		if _, _, err := PathAndFragment(rawURL); err == nil {
			t.Errorf("PathAndFragment(%q) returned a nil error, want the error from url.Parse", rawURL)
		}
	}
}

func TestQueryValues(t *testing.T) {
	const search = "https://example.com/search?q=go&tag=a&tag=b&msg=hello%20world%21#results"
	tests := []struct {
		rawURL, key string
		want        []string
	}{
		{pgURL, "k", []string{"v"}},
		{search, "q", []string{"go"}},
		{search, "tag", []string{"a", "b"}},
		{search, "msg", []string{"hello world!"}},
		{search, "page", nil},
		{"https://example.com/no-query", "q", nil},
	}
	for _, tt := range tests {
		got, err := QueryValues(tt.rawURL, tt.key)
		if err != nil {
			t.Errorf("QueryValues(%q, %q) returned error %v, want nil", tt.rawURL, tt.key, err)
		}
		if !slices.Equal(got, tt.want) {
			t.Errorf("QueryValues(%q, %q) = %q, want %q", tt.rawURL, tt.key, got, tt.want)
		}
	}
	for _, rawURL := range badURLs {
		if _, err := QueryValues(rawURL, "k"); err == nil {
			t.Errorf("QueryValues(%q, \"k\") returned a nil error, want the error from url.Parse", rawURL)
		}
	}
}

func TestParseDBConfig(t *testing.T) {
	tests := []struct {
		rawURL string
		want   DBConfig
	}{
		{
			"postgres://admin:s3cret@db.example.com:5432/shop?sslmode=require&timeout=10",
			DBConfig{
				Driver: "postgres", User: "admin", Password: "s3cret",
				Host: "db.example.com", Port: 5432, Database: "shop",
				Options: map[string]string{"sslmode": "require", "timeout": "10"},
			},
		},
		{
			pgURL,
			DBConfig{
				Driver: "postgres", User: "user", Password: "pass",
				Host: "host.com", Port: 5432, Database: "path",
				Options: map[string]string{"k": "v"},
			},
		},
		{
			"mysql://root@localhost:3306/app",
			DBConfig{
				Driver: "mysql", User: "root", Password: "",
				Host: "localhost", Port: 3306, Database: "app",
				Options: map[string]string{},
			},
		},
		{
			"redis://[::1]:6379/0?tag=first&tag=second",
			DBConfig{
				Driver: "redis", Host: "::1", Port: 6379, Database: "0",
				Options: map[string]string{"tag": "first"},
			},
		},
	}
	for _, tt := range tests {
		got, err := ParseDBConfig(tt.rawURL)
		if err != nil {
			t.Errorf("ParseDBConfig(%q) returned error %v, want nil", tt.rawURL, err)
			continue
		}
		if got.Options == nil {
			t.Errorf("ParseDBConfig(%q): Options is nil, want a non-nil map", tt.rawURL)
		}
		if got.Driver != tt.want.Driver || got.User != tt.want.User || got.Password != tt.want.Password ||
			got.Host != tt.want.Host || got.Port != tt.want.Port || got.Database != tt.want.Database ||
			!maps.Equal(got.Options, tt.want.Options) {
			t.Errorf("ParseDBConfig(%q) =\n  %+v, want\n  %+v", tt.rawURL, got, tt.want)
		}
	}
}

func TestParseDBConfigErrors(t *testing.T) {
	tests := []struct {
		why    string
		rawURL string
	}{
		{"there is no port", "postgres://admin:s3cret@db.example.com/shop"},
		{"there is no host and port", "postgres:///shop"},
	}
	for _, rawURL := range badURLs {
		tests = append(tests, struct{ why, rawURL string }{"url.Parse rejects it", rawURL})
	}
	for _, tt := range tests {
		if _, err := ParseDBConfig(tt.rawURL); err == nil {
			t.Errorf("ParseDBConfig(%q) returned a nil error, want an error because %s", tt.rawURL, tt.why)
		}
	}
}
