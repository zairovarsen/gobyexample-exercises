//go:build !solution

// Package environmentvariables contains exercises for https://gobyexample.com/environment-variables
//
// Run the tests for this section with:
//
//	go test ./76-environment-variables/
package environmentvariables

// Config is the configuration that LoadConfig reads from the environment.
type Config struct {
	Host  string
	Port  int
	Debug bool
}

// GetOrDefault returns the value of the environment variable key, or def when
// the variable is not set or set to the empty string.
//
// Example: with FOO=1 in the environment, GetOrDefault("FOO", "x") is "1" and
// GetOrDefault("BAR", "x") is "x".
//
// Hint: os.Getenv returns "" for a variable that does not exist.
func GetOrDefault(key, def string) string {
	// TODO: read the variable, fall back to def.
	return ""
}

// Status describes the state of the environment variable key in one of three
// ways (shown for the key "FOO"):
//
//	"FOO is not set"  the variable does not exist
//	"FOO is empty"    the variable exists, its value is ""
//	"FOO=1"           the variable exists with the value "1"
//
// os.Getenv cannot tell the first two apart - both give "". Use
// os.LookupEnv, which also reports whether the variable exists.
func Status(key string) string {
	// TODO: look the variable up and describe what you found.
	return ""
}

// SetAll sets every key/value pair of vars as an environment variable of the
// current process with os.Setenv. It stops at the first failure and returns
// that error (os.Setenv rejects an empty key, for example).
func SetAll(vars map[string]string) error {
	// TODO: loop over the map and set each variable.
	return nil
}

// EnvMap returns all environment variables whose name starts with prefix as a
// map from the full name to the value.
//
// Example: with GBE_HOST=localhost and GBE_PORT=80 in the environment,
// EnvMap("GBE_") is {"GBE_HOST": "localhost", "GBE_PORT": "80"}.
//
// os.Environ lists the whole environment as "KEY=value" strings. Split each
// one with strings.SplitN(e, "=", 2) like the page does: the limit of 2
// matters because a VALUE may contain "=" characters as well.
func EnvMap(prefix string) map[string]string {
	// TODO: go through os.Environ, split the entries, keep the matching ones.
	return nil
}

// LoadConfig builds a Config from three environment variables whose names
// start with prefix. For the prefix "APP_" these are:
//
//	APP_HOST   the host name               default "localhost"
//	APP_PORT   a whole number              default 8080
//	APP_DEBUG  a boolean such as "true",   default false
//	           "false", "1" or "0"
//
// A variable that is not set or empty gets its default. When APP_PORT is not
// a number or APP_DEBUG is not a boolean, LoadConfig returns the zero Config
// and an error that mentions the name of the offending variable.
//
// Hint: environment variables are always strings; strconv.Atoi and
// strconv.ParseBool convert them.
func LoadConfig(prefix string) (Config, error) {
	// TODO: read the three variables, apply defaults, convert, report errors.
	return Config{}, nil
}
