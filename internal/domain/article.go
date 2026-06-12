package domain

import "time"

type Article struct {
	ID         string `bson:"_id,omitempty"`
	Title      string `bson:"title"`
	Body       string `bson:"body"`
	Tags       []Tag `bson:"tags"`
	Created_at time.Time `bson:"created_at"`
}

type Tag struct {
	Word string
	Freq int64
}
