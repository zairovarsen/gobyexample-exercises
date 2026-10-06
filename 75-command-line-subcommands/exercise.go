//go:build !solution

// Package commandlinesubcommands contains exercises for https://gobyexample.com/command-line-subcommands
//
// Run the tests for this section with:
//
//	go test ./75-command-line-subcommands/
package commandlinesubcommands

import "errors"

// The program on the page has two subcommands, each with its own flags:
//
//	prog foo -enable -name=joe a1 a2
//	prog bar -level 8 a1
//
// In all exercises args are the arguments WITHOUT the program name; a real
// program would pass os.Args[1:], so the subcommand is args[0]. The page
// prints its results and exits on errors; here the functions return values
// and errors instead.
//
// As in the previous section, create every flag set with
// flag.NewFlagSet(name, flag.ContinueOnError) (the page uses
// flag.ExitOnError) and silence it with fs.SetOutput(io.Discard).

// Errors returned by Run and Dispatch. Compare with errors.Is.
var (
	// ErrNoSubcommand: no arguments were given at all.
	ErrNoSubcommand = errors.New("expected 'foo' or 'bar' subcommands")
	// ErrUnknownSubcommand: the first argument is not a known subcommand.
	ErrUnknownSubcommand = errors.New("unknown subcommand")
)

// FooOptions are the parsed arguments of the "foo" subcommand.
type FooOptions struct {
	Enable bool     // -enable, default false
	Name   string   // -name, default ""
	Tail   []string // positional arguments after the flags
}

// BarOptions are the parsed arguments of the "bar" subcommand.
type BarOptions struct {
	Level int      // -level, default 0
	Tail  []string // positional arguments after the flags
}

// ParseFoo parses the arguments that follow the word "foo" on the command
// line. It uses its own flag set named "foo" with a bool flag "enable"
// (default false) and a string flag "name" (default "").
//
// Example: ParseFoo([]string{"-enable", "-name=joe", "a1", "a2"}) returns
// FooOptions{Enable: true, Name: "joe", Tail: ["a1", "a2"]}.
//
// When parsing fails (e.g. the unknown flag -level) return the zero
// FooOptions and the error.
func ParseFoo(args []string) (FooOptions, error) {
	// TODO: flag set "foo", two flags, parse, fill in the struct.
	return FooOptions{}, nil
}

// ParseBar parses the arguments that follow the word "bar". It uses its own
// flag set named "bar" with an int flag "level" (default 0).
//
// Example: ParseBar([]string{"-level", "8", "a1"}) returns
// BarOptions{Level: 8, Tail: ["a1"]}.
//
// The flags of foo do not exist here: ParseBar([]string{"-enable"}) is an
// error.
func ParseBar(args []string) (BarOptions, error) {
	// TODO: flag set "bar", one flag, parse, fill in the struct.
	return BarOptions{}, nil
}

// Run is the whole program: it looks at args[0] to pick the subcommand, hands
// the remaining arguments to ParseFoo or ParseBar and returns a one-line
// description of the result:
//
//	Run([]string{"foo", "-enable", "-name=joe", "a1", "a2"})
//	    "foo: enable=true name=joe tail=[a1 a2]"
//	Run([]string{"bar", "-level", "8", "a1"})
//	    "bar: level=8 tail=[a1]"
//
// (the tail is formatted with %v, so an empty tail prints as "[]").
//
// Errors, each returned together with "":
//   - no arguments at all: ErrNoSubcommand
//   - args[0] is neither "foo" nor "bar": an error that wraps
//     ErrUnknownSubcommand (use %w) and mentions the offending word
//   - the subcommand's flags do not parse: the error from ParseFoo/ParseBar
//
// Use a switch on the subcommand like the page does.
func Run(args []string) (string, error) {
	// TODO: check for a subcommand, switch on it, format the result.
	return "", nil
}

// Dispatch is the general form of Run for any set of subcommands: commands
// maps a subcommand name to the function that implements it. Dispatch looks
// up args[0], calls the function with the remaining arguments (args[1:]) and
// returns exactly what the function returns.
//
// The errors are the same as for Run: ErrNoSubcommand when args is empty, and
// an error wrapping ErrUnknownSubcommand that mentions the name when it is
// not a key of commands.
func Dispatch(args []string, commands map[string]func(args []string) (string, error)) (string, error) {
	// TODO: check for a subcommand, look it up in the map, call it.
	return "", nil
}
