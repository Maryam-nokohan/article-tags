package application

import (
	"sort"

	"github.com/maryam-nokohan/go-article/internal/domain"
	"github.com/maryam-nokohan/go-article/internal/utils"
)

type TagExtractorService struct{
	stopWords map[string]bool
}

func NewTagEctractorService() *TagExtractorService{
	s := utils.LoadStopWords()
	return  &TagExtractorService{
		stopWords: s,
	}
}
func(t *TagExtractorService)isStopWord(word string) bool{
	_, ok := t.stopWords[word]
	return  ok
}
func (t *TagExtractorService)Extract(text string , topN int64) []domain.Tag {

	words := utils.NormilizeText(text)

	wordFreq := make(map[string]int)

	for _, word := range words {
		wordFreq[word]++
	}

	type pair struct {
		word  string
		count int
	}

	var pairs []pair
	for k, v := range wordFreq {
		pairs = append(pairs, pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})
	var tags []domain.Tag
	for _, p := range pairs {
		if !t.isStopWord(p.word) {
			tags = append(tags, domain.Tag{
				Word: p.word,
				Freq: int64(p.count),
			})
		}
		if len(tags) == int(topN) {
			break
		}
	}

	return tags

}

