//go:build solution

package commandlinesubcommands

import (
	"errors"
	"flag"
	"fmt"
	"io"
)

var (
	ErrNoSubcommand      = errors.New("expected 'foo' or 'bar' subcommands")
	ErrUnknownSubcommand = errors.New("unknown subcommand")
)

type FooOptions struct {
	Enable bool
	Name   string
	Tail   []string
}

type BarOptions struct {
	Level int
	Tail  []string
}

func ParseFoo(args []string) (FooOptions, error) {
	fooCmd := flag.NewFlagSet("foo", flag.ContinueOnError)
	fooCmd.SetOutput(io.Discard)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	if err := fooCmd.Parse(args); err != nil {
		return FooOptions{}, err
	}
	return FooOptions{Enable: *fooEnable, Name: *fooName, Tail: fooCmd.Args()}, nil
}

func ParseBar(args []string) (BarOptions, error) {
	barCmd := flag.NewFlagSet("bar", flag.ContinueOnError)
	barCmd.SetOutput(io.Discard)
	barLevel := barCmd.Int("level", 0, "level")

	if err := barCmd.Parse(args); err != nil {
		return BarOptions{}, err
	}
	return BarOptions{Level: *barLevel, Tail: barCmd.Args()}, nil
}

func Run(args []string) (string, error) {
	if len(args) < 1 {
		return "", ErrNoSubcommand
	}

	switch args[0] {
	case "foo":
		foo, err := ParseFoo(args[1:])
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("foo: enable=%t name=%s tail=%v", foo.Enable, foo.Name, foo.Tail), nil
	case "bar":
		bar, err := ParseBar(args[1:])
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("bar: level=%d tail=%v", bar.Level, bar.Tail), nil
	default:
		return "", fmt.Errorf("%w: %q", ErrUnknownSubcommand, args[0])
	}
}

func Dispatch(args []string, commands map[string]func(args []string) (string, error)) (string, error) {
	if len(args) < 1 {
		return "", ErrNoSubcommand
	}
	command, ok := commands[args[0]]
	if !ok {
		return "", fmt.Errorf("%w: %q", ErrUnknownSubcommand, args[0])
	}
	return command(args[1:])
}
