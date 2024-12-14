package store

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type News struct {
	ID        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	Content   string    `json:"content"`
	Source    *url.URL  `json:"source"`
	CreatedAt time.Time `json:"created_at"`
	Tags      []string  `json:"tags"`
}
