package main

import "fmt"

func hello(name string) string {
	const prefixMessage = "Olá, "

	return prefixMessage + name
}

func main() {
	fmt.Println(hello("madruga665"))
}
