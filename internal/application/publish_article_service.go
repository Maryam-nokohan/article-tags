package application

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/maryam-nokohan/go-article/internal/domain"
	"github.com/maryam-nokohan/go-article/internal/ports"
	"github.com/google/uuid"
)

type IngestionService struct {
    broker  ports.MessageBroker
    storage ports.ObjectStorage
}

func NewIngestionService(b ports.MessageBroker, s ports.ObjectStorage) *IngestionService {
    return &IngestionService{
        broker:  b,
        storage: s,
    }
}

func (s *IngestionService) AcceptArticle(article *domain.Article) error {
    article.Created_at = time.Now()
    article.ID = uuid.New().String()

    raw, err := json.Marshal(article)
    if err != nil {
        return err
    }

    key := fmt.Sprintf("articles/%s.json", article.ID)

    _, err = s.storage.Upload(key, raw)
    if err != nil {
        return err
    }

    evt := domain.ArticleCreatedEvent{
        ArticleID: article.ID,
        Title:     article.Title,
        ObjectKey: key,
        CreatedAt: article.Created_at.Format(time.RFC3339),
    }

    payload, err := json.Marshal(evt)
    if err != nil {
        return err
    }

    return s.broker.Publish("article.created", payload)
}
