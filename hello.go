package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Ola, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
