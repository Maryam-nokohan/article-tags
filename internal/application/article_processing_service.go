package application

import (
	"context"
	"encoding/json"
	"time"

	"github.com/maryam-nokohan/go-article/internal/domain"
	"github.com/maryam-nokohan/go-article/internal/ports"
)

type ProcessingService struct {
	storage   ports.ObjectStorage
	repo      ports.ArticleRepository
	extractor ports.TagExtractor
}

func NewProcessingService(
	storage ports.ObjectStorage,
	repo ports.ArticleRepository,
	extractor ports.TagExtractor,
) *ProcessingService {
	return &ProcessingService{
		storage:   storage,
		repo:      repo,
		extractor: extractor,
	}
}

func (s *ProcessingService) HandleArticleCreated(data []byte) error {
	var evt domain.ArticleCreatedEvent
	if err := json.Unmarshal(data, &evt); err != nil {
		return err
	}

	raw, err := s.storage.Download(evt.ObjectKey)
	if err != nil {
		return err
	}

	var article domain.Article
	if err := json.Unmarshal(raw, &article); err != nil {
		return err
	}

	article.Tags = s.extractor.Extract(article.Body, -1)
	article.Created_at = time.Now()

	return s.repo.Save(context.Background(), &article)
}
func (s *ProcessingService) GetTopTags(ctx context.Context, topN int64) ([]domain.Tag, error) {
	if topN <= 0 {
		topN = 10
	}

	tags, err := s.repo.GetTopTags(ctx, topN)
	if err != nil {
		return nil, err
	}

	return tags, nil
}
