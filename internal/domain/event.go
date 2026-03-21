package domain

type ArticleCreatedEvent struct {
    ArticleID string `json:"article_id"`
    Title     string `json:"title"`
    ObjectKey string `json:"object_key"`
    CreatedAt string `json:"created_at"`
}
