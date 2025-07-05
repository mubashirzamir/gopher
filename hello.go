package main

import "fmt"

const (
	languageEnglish = "English"
	languageFrench  = "French"
	languageSpanish = "Spanish"

	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	return getPrefix(language) + getName(name)
}

func getPrefix(language string) (prefix string) {
	switch language {
	case languageFrench:
		prefix = frenchHelloPrefix
	case languageSpanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func getName(name string) string {
	if name == "" {
		return "World"
	}

	return name
}

func main() {
	fmt.Println(Hello("world", ""))
}
