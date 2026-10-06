//go:build solution

package commandlineflags

import (
	"flag"
	"fmt"
	"io"
)

type Config struct {
	Word string
	Numb int
	Fork bool
	SVar string
	Tail []string
}

func ParseWord(args []string) (string, error) {
	fs := flag.NewFlagSet("word", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	wordPtr := fs.String("word", "foo", "a string")

	if err := fs.Parse(args); err != nil {
		return "", err
	}
	return *wordPtr, nil
}

func ParseNumbFork(args []string) (numb int, fork bool, err error) {
	fs := flag.NewFlagSet("numbfork", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	numbPtr := fs.Int("numb", 42, "an int")
	forkPtr := fs.Bool("fork", false, "a bool")

	if err := fs.Parse(args); err != nil {
		return 0, false, err
	}
	return *numbPtr, *forkPtr, nil
}

func Parse(args []string, output io.Writer) (Config, error) {
	var cfg Config

	fs := flag.NewFlagSet("command-line-flags", flag.ContinueOnError)
	fs.SetOutput(output)
	wordPtr := fs.String("word", "foo", "a string")
	numbPtr := fs.Int("numb", 42, "an int")
	forkPtr := fs.Bool("fork", false, "a bool")
	fs.StringVar(&cfg.SVar, "svar", "bar", "a string var")

	if err := fs.Parse(args); err != nil {
		return Config{}, err
	}
	cfg.Word = *wordPtr
	cfg.Numb = *numbPtr
	cfg.Fork = *forkPtr
	cfg.Tail = fs.Args()
	return cfg, nil
}

func ParseServer(args []string) (string, error) {
	fs := flag.NewFlagSet("server", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	host := fs.String("host", "localhost", "host name to listen on")
	port := fs.Int("port", 8080, "port to listen on")
	tls := fs.Bool("tls", false, "serve https instead of http")

	if err := fs.Parse(args); err != nil {
		return "", err
	}
	if fs.NArg() > 0 {
		return "", fmt.Errorf("unexpected arguments: %q", fs.Args())
	}
	if *port < 1 || *port > 65535 {
		return "", fmt.Errorf("port %d is out of range 1-65535", *port)
	}

	scheme := "http"
	if *tls {
		scheme = "https"
	}
	return fmt.Sprintf("%s://%s:%d", scheme, *host, *port), nil
}
