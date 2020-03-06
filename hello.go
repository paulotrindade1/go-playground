package main

import "fmt"

const portuguese = "Portuguese"
const french = "French"
const englishHelloPrefix = "Hello, "
const portugueseHelloPrefix = "Olá, "
const frenchHelloPrefix = "Bonjour, " 

func Hello(name string, language string) string{
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix 

	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portugueseHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main () {
	fmt.Println(Hello("Paulo", "Portuguese"))
}
