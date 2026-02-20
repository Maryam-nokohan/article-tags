package domain

import "context"

type Tag struct {
	Tag string
	Freq int64
}

type ArticleRepository interface {
	Save(ctx context.Context , article *Article)error
	GetTopTags(ctx context.Context , limit int64) ([]Tag , error)
}