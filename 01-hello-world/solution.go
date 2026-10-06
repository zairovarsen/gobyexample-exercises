//go:build solution

package helloworld

import "fmt"

func Greeting() string {
	return "hello world"
}

func GreetingFor(name string) string {
	if name == "" {
		return Greeting()
	}
	return "hello " + name
}

func PrintGreeting() {
	fmt.Println(Greeting())
}
