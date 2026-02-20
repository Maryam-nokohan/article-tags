package application

import (
	"context"

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

func (s *ArticleService) ProcessArticle(ctx context.Context , articles []domain.Article) (error){

	return  nil
}

func (s *ArticleService) GetTopTags(ctx context.Context , limit int64)([]domain.Tag , error){

	
	return  nil , nil
}
