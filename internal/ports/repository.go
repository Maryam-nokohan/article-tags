package domain

import (
	"context"

	"github.com/maryam-nokohan/go-article/internal/domain"
)

type Tag struct {
	Word string
	Freq int64
}

type ArticleRepository interface {
	Save(ctx context.Context , article *domain.Article)error
	GetTopTags(ctx context.Context , limit int64) ([]Tag , error)
}