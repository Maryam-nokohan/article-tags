package domain

import (
	"context"

)

type ArticleRepository interface {
	Save(ctx context.Context , article *Article)error
	GetTopTags(ctx context.Context , limit int64) ([]Tag , error)
}