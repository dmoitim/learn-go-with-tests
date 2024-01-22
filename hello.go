package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portugueseHelloPrefix = "Olá, "
const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("DevA", ""))
}
