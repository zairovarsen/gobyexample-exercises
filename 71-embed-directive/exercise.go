//go:build !solution

// Package embeddirective contains exercises for https://gobyexample.com/embed-directive
//
// Run the tests for this section with:
//
//	go test ./71-embed-directive/
//
// The files to embed live in the assets folder next to this file:
//
//	assets/motd.txt
//	assets/config/dev.env
//	assets/config/prod.env
//	assets/config/notes.txt
package embeddirective

import "embed"

// TODO: add a //go:embed directive directly above this variable so that it
// holds the content of assets/motd.txt.
var motd string

// TODO: embed the same file again, this time as raw bytes.
var motdBytes []byte

// TODO: embed every .env file in assets/config (and only the .env files - not
// notes.txt) into this file system, using a wildcard.
var configs embed.FS

// ConfigNames returns the names (without directories) of the files embedded
// under assets/config, in alphabetical order: ["dev.env", "prod.env"].
//
// Hint: embed.FS has a ReadDir method. Paths inside an embed.FS always use
// forward slashes, on every operating system.
func ConfigNames() ([]string, error) {
	// TODO: list the embedded directory.
	return nil, nil
}

// ConfigValue looks up key in the embedded file assets/config/<env>.env.
// These files consist of KEY=VALUE lines, so ConfigValue("dev", "PORT")
// returns "8080".
//
// It returns a non-nil error when there is no embedded file for env, or when
// the file has no line for key.
func ConfigValue(env, key string) (string, error) {
	// TODO: read the file from configs and search its lines.
	return "", nil
}
