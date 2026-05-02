package main

import "fmt"

func hello(name string) string {
	const prefixMessage = "Hello, "

	if name == "" {
		name = "world"
	}

	return prefixMessage + name
}

func main() {
	fmt.Println(hello("madruga665"))
}
