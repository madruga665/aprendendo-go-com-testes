package main

import (
	"fmt"
)

const prefixMessage = "Hello, "
const portuguesePrefixMessage = "Olá, "
const spanishPrefixMessage = "Hola, "

func prefixOfMessage(language string) (prefix string) {

	switch language {
	case "portuguese":
		prefix = portuguesePrefixMessage
	case "spanish":
		prefix = spanishPrefixMessage
	default:
		prefix = prefixMessage
	}

	return
}

func hello(name string, language string) string {

	if name == "" {
		name = "world"
	}

	return prefixOfMessage(language) + name
}

func main() {
	fmt.Println(hello("madruga665", "portuguese"))
}
