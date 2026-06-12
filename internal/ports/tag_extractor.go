package ports

import "github.com/maryam-nokohan/go-article/internal/domain"

type TagExtractor interface {
	Extract(text string, topN int64) []domain.Tag
	IsStopWord(word string) bool
}
