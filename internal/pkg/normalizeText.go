package pkg

import (
	"log"
	"os"
	"strings"
	"unicode"
)

func LoadStopWords() map[string]bool {
	data, err := os.ReadFile("internal/utils/stopwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	stopWords := map[string]bool{}
	for _, w := range strings.Split(string(data), "\n") {
		stopWords[strings.TrimSpace(w)] = true
	}
	return stopWords
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
