package domain

import "time"

type Article struct {
	ID         string
	Title      string
	Body       string
	Tags       []string
	Created_at time.Time
}
