//go:build !solution

// Package commandlineflags contains exercises for https://gobyexample.com/command-line-flags
//
// Run the tests for this section with:
//
//	go test ./74-command-line-flags/
package commandlineflags

import "io"

// How this section maps to the page:
//
// The page declares its flags with the package-level functions flag.String,
// flag.Int, flag.Bool and flag.StringVar and then calls flag.Parse(). Those
// functions work on one global flag set (flag.CommandLine) that parses
// os.Args[1:] and EXITS the program when the arguments are wrong - impossible
// to test.
//
// The exercises use the same API one level down: create your own set with
//
//	fs := flag.NewFlagSet("name", flag.ContinueOnError)
//
// and call the methods of the same names on it: fs.String, fs.Int, fs.Bool,
// fs.StringVar, fs.Parse(args), fs.Args(). With flag.ContinueOnError a wrong
// argument makes fs.Parse return an error instead of exiting.
//
// In all exercises args are the arguments WITHOUT the program name; a real
// program would pass os.Args[1:].
//
// A FlagSet prints error and usage messages to os.Stderr. To keep the test
// output clean, redirect them with fs.SetOutput(io.Discard) unless the
// exercise says otherwise.

// Config holds the values of the flags declared on the page.
type Config struct {
	Word string   // -word, default "foo"
	Numb int      // -numb, default 42
	Fork bool     // -fork, default false
	SVar string   // -svar, default "bar"
	Tail []string // the positional arguments that follow the flags
}

// ParseWord declares a single string flag called "word" with the default
// value "foo" and the usage text "a string", parses args and returns the
// value of the flag.
//
// Example: ParseWord([]string{"-word=opt"}) is "opt", ParseWord(nil) is
// "foo". An unknown flag such as -wat makes it return "" and the error from
// fs.Parse.
//
// Note that fs.String returns a pointer; it only points to the parsed value
// after fs.Parse has run.
func ParseWord(args []string) (string, error) {
	// TODO: new flag set, declare the flag, parse, dereference.
	return "", nil
}

// ParseNumbFork declares an int flag "numb" (default 42, usage "an int") and
// a bool flag "fork" (default false, usage "a bool"), parses args and returns
// both values.
//
// Example: ParseNumbFork([]string{"-numb=7", "-fork"}) is 7, true.
// A value that is not a number (-numb=seven) is an error.
func ParseNumbFork(args []string) (numb int, fork bool, err error) {
	// TODO: new flag set, declare both flags, parse.
	return 0, false, nil
}

// Parse is the complete program from the page: it declares the flags
//
//	-word  string  default "foo"  usage "a string"
//	-numb  int     default 42     usage "an int"
//	-fork  bool    default false  usage "a bool"
//	-svar  string  default "bar"  usage "a string var"
//
// parses args and returns the values in a Config. Tail gets the positional
// arguments that are left over after the flags (fs.Args()). Flags must come
// before positional arguments: parsing stops at the first argument that is
// not a flag, so in "-word=opt a1 -numb=7" the tail is ["a1", "-numb=7"] and
// numb keeps its default.
//
// Declare -svar with fs.StringVar so that it writes straight into the SVar
// field of your Config (the page uses flag.StringVar with an existing
// variable).
//
// Send the flag set's messages to output with fs.SetOutput(output): when
// parsing fails, or when the automatic -h / -help flag is given (fs.Parse
// then returns flag.ErrHelp), the error message and the list of flags are
// written there. On an error return the zero Config and the error.
func Parse(args []string, output io.Writer) (Config, error) {
	// TODO: declare the four flags, parse, fill in the Config.
	return Config{}, nil
}

// ParseServer parses the flags of an imaginary web server
//
//	-host  string  default "localhost"
//	-port  int     default 8080
//	-tls   bool    default false
//
// and returns the URL the server would be reachable at: "http://host:port",
// or "https://host:port" when -tls is given.
//
// Example: ParseServer([]string{"-port=9000", "-tls"}) is
// "https://localhost:9000".
//
// Parsing alone is not enough here - validate the result as well. Return ""
// and an error when
//   - fs.Parse fails,
//   - the port is not in the range 1 to 65535,
//   - positional arguments are left over (the server takes none).
func ParseServer(args []string) (string, error) {
	// TODO: declare the flags, parse, validate, format the URL.
	return "", nil
}
