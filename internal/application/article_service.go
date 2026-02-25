package application

import (
	"context"
	"log"
	"time"

	"github.com/maryam-nokohan/go-article/internal/domain"
	"github.com/maryam-nokohan/go-article/internal/ports"
)

type ArticleService struct {
	repo      ports.ArticleRepository
	extractor ports.TagExtractor
}

func NewArticleService(r ports.ArticleRepository, e ports.TagExtractor) *ArticleService {
	return &ArticleService{
		repo:      r,
		extractor: e,
	}
}

func (s *ArticleService) ProcessArticle(articles *domain.Article) error {
	log.Println("In ArticleService : Processing article with title:", articles.Title)
	tagWords := s.extractor.Extract(articles.Body, -1)
	articles.Tags = tagWords
	articles.Created_at = time.Now()
	err := s.repo.Save(context.Background(), articles)
	return err
}

func (s *ArticleService) GetTopTags(ctx context.Context, limit int64) ([]domain.Tag, error) {
	log.Println("In ArticleService : Getting top tags with limit:", limit)
	tags, err := s.repo.GetTopTags(ctx, limit)
	if err != nil {
		return nil, err
	}
	return tags, nil
}
