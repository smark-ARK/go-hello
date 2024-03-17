package main

import (
	"flag"
	"fmt"
)

type language string

var phraseBook = map[language]string{
	"en": "Hello World!",
	"ar": "مرحبًا بالعالم!",
	"fr": "Bonjour, le monde !",
	"es": "¡Hola, mundo!",
}

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The language in which user expects the greeting")
	flag.Parse()
	greeting := greet(language(lang))
	fmt.Println(greeting)
}

func greet(l language) string {
	greeting, ok := phraseBook[l]
	if !ok {
		return fmt.Sprintf("Language %v Not Sypported", l)
	}
	return greeting

}
