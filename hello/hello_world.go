package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, lang string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) string {
    prefix := englishHelloPrefix

    switch lang {
    case "French":
        prefix = frenchHelloPrefix
    case "Spanish":
        prefix = spanishHelloPrefix 
    }

    return prefix
}

func main() {
	fmt.Println(Hello("Shinto", "English"))
}
