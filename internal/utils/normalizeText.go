package utils

import (
	"log"
	"os"
	"strings"
	"unicode"
)

var stopWords map[string]bool

func LoadStopWords() {
	data, err := os.ReadFile("stop_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	stopWords = map[string]bool{}
	for _, w := range strings.Split(string(data), "\n") {
		stopWords[strings.TrimSpace(w)] = true
	}
}

func isStopWord(word string) bool {
	_, ok := stopWords[word]
	return ok
}

func NormilizeText(text string) []string {
	text = strings.ToLower(text)

	var builder strings.Builder
	for _, ch := range text {
		if !unicode.IsPunct(ch) {
			builder.WriteRune(ch)
		}
	}

	words := strings.Fields(builder.String())
	return words
}
