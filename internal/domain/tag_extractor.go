package domain

type TagExtractor interface {
	Extract(text string, topN int64) []Tag
	IsStopWord(word string) bool
}
