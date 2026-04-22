package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"
	italian = "Italian"

	englishGreetingPrefix = "Hello, "
	spanishGreetingPrefix = "Hola, "
	frenchGreetingPrefix = "Bonjour, "
	italianGreetingPrefix = "Ciao, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishGreetingPrefix
	case french:
		prefix = frenchGreetingPrefix
	case italian:
		prefix = italianGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}