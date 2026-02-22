package application

import (
	"context"
	"time"

	"github.com/maryam-nokohan/go-article/internal/domain"
)

type ArticleService struct {
	repo      domain.ArticleRepository
	extractor domain.TagExtractor
}

func NewArticleService(r domain.ArticleRepository, e domain.TagExtractor) *ArticleService {
	return &ArticleService{
		repo:      r,
		extractor: e,
	}
}

func (s *ArticleService) ProcessArticle(articles domain.Article) error {

	tagWords := s.extractor.Extract(articles.Body, -1)
	articles.Tags = tagWords
	articles.Created_at = time.Now()
	err := s.repo.Save(context.Background(), &articles)
	return err
}

func (s *ArticleService) GetTopTags(ctx context.Context, limit int64) ([]domain.Tag, error) {
	tags, err := s.repo.GetTopTags(ctx, limit)
	if err != nil {
		return nil, err
	}
	return tags, nil
}
