package ports

import (
	"context"

	"github.com/maryam-nokohan/go-article/internal/domain"
)

type ArticleRepository interface {
	Save(ctx context.Context, article *domain.Article) error
	GetTopTags(ctx context.Context, limit int64) ([]domain.Tag, error)
}
