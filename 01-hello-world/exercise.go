//go:build !solution

// Package helloworld contains exercises for https://gobyexample.com/hello-world
//
// Run the tests for this section with:
//
//	go test ./01-hello-world/
package helloworld

import "fmt"

// Greeting returns the classic first-program message: "hello world".
func Greeting() string {
	return "hello world"
}

// GreetingFor returns "hello <name>", e.g. GreetingFor("gopher") returns
// "hello gopher". When name is empty it falls back to "hello world".
func GreetingFor(name string) string {
	if name != "" {
		return "hello " + name 
	}
	return "hello world"
}

// PrintGreeting prints "hello world" followed by a newline to standard output.
//
// Hint: you need to import the "fmt" package.
func PrintGreeting() {
	// TODO: print the message.
	    fmt.Println("hello world")
}
