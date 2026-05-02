package main

import "fmt"

func hello(name string, language string) string {
	const prefixMessage = "Hello, "
	const portuguesePrefixMessage = "Olá, "
	const spanishPrefixMessage = "Hola, "

	if name == "" {
		name = "world"
	}

	if language == "portuguese" {
		return portuguesePrefixMessage + name
	}

	if language == "spanish" {
		return spanishPrefixMessage + name
	}

	return prefixMessage + name
}

func main() {
	fmt.Println(hello("madruga665", "portuguese"))
}
