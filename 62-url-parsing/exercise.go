//go:build !solution

// Package urlparsing contains exercises for https://gobyexample.com/url-parsing
//
// Run the tests for this section with:
//
//	go test ./62-url-parsing/
package urlparsing

// All exercises start the same way: url.Parse(rawURL) from the standard
// "net/url" package turns the string into a *url.URL, whose fields and
// methods give access to the parts. The example URL of the gobyexample page
// contains all of them:
//
//	postgres://user:pass@host.com:5432/path?k=v#f
//
// Every function returns the error from url.Parse (and zero values for the
// other results) when rawURL cannot be parsed.

// DBConfig is the result of ParseDBConfig.
type DBConfig struct {
	Driver   string            // the URL scheme, e.g. "postgres"
	User     string            // user name, "" if the URL has none
	Password string            // password, "" if the URL has none
	Host     string            // host name without the port
	Port     int               // port number
	Database string            // the URL path without its leading "/"
	Options  map[string]string // query parameters (first value of each key)
}

// Scheme returns the scheme of the URL - the part before "://".
//
// Example: for "postgres://user:pass@host.com:5432/path?k=v#f" it returns
// "postgres".
func Scheme(rawURL string) (string, error) {
	// TODO: parse the URL and return its Scheme field.
	return "", nil
}

// Credentials returns the user name and the password from the URL. Both
// are empty when the URL has no "user:pass@" part; the password is empty
// when only a user name is given.
//
// Example: for "postgres://user:pass@host.com:5432/path" it returns "user"
// and "pass".
//
// The User field holds both; see its Username and Password methods.
func Credentials(rawURL string) (user, password string, err error) {
	// TODO: parse the URL and take the credentials from u.User.
	return "", "", nil
}

// HostPort returns the host name and the port of the URL as two separate
// strings. When the URL names no port, port is "" and host is the complete
// host part.
//
// Examples:
//
//	"postgres://user:pass@host.com:5432/path" -> "host.com", "5432"
//	"https://gobyexample.com/url-parsing"     -> "gobyexample.com", ""
//	"http://[::1]:8080/"                      -> "::1", "8080"
//
// The Host field contains both ("host.com:5432"); net.SplitHostPort from
// the "net" package takes that apart. Mind that SplitHostPort returns an
// error when there is no port to split off - u.Port() tells you up front.
func HostPort(rawURL string) (host, port string, err error) {
	// TODO: parse the URL and split its Host field.
	return "", "", nil
}

// PathAndFragment returns the path of the URL and its fragment (the part
// after "#", without the "#").
//
// Example: for "postgres://user:pass@host.com:5432/path?k=v#f" it returns
// "/path" and "f".
func PathAndFragment(rawURL string) (path, fragment string, err error) {
	// TODO: parse the URL and return the two fields.
	return "", "", nil
}

// QueryValues returns all values that the query string of the URL (the part
// between "?" and "#") has for key, in order. A key that does not occur
// gives an empty result.
//
// Examples for "https://example.com/search?q=go&tag=a&tag=b":
//
//	key "q"    -> [go]
//	key "tag"  -> [a b]
//	key "page" -> []
//
// The RawQuery field holds the query as one string ("q=go&tag=a&tag=b");
// url.ParseQuery turns that into a map from each key to a slice of its
// values, decoding escapes such as %20 on the way.
func QueryValues(rawURL, key string) ([]string, error) {
	// TODO: parse the URL, parse its query, look up key.
	return nil, nil
}

// ParseDBConfig takes a database connection URL such as
//
//	postgres://admin:s3cret@db.example.com:5432/shop?sslmode=require&timeout=10
//
// apart into a DBConfig. For the URL above the result is
//
//	DBConfig{
//		Driver:   "postgres",
//		User:     "admin",
//		Password: "s3cret",
//		Host:     "db.example.com",
//		Port:     5432,
//		Database: "shop",
//		Options:  map[string]string{"sslmode": "require", "timeout": "10"},
//	}
//
// The port is required here: ParseDBConfig returns an error when the URL
// cannot be parsed or has no port. Options is an empty, non-nil map when
// there are no query parameters; when a key occurs more than once its first
// value counts.
//
// Hint: strconv.Atoi (previous section) converts the port.
func ParseDBConfig(rawURL string) (DBConfig, error) {
	// TODO: parse the URL and fill in a DBConfig.
	return DBConfig{}, nil
}
