package handler

import (
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/nabind47/go_rest47/internal/store"
)

type NewsPostRequestBody struct {
	ID        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary"`
	CreatedAt string    `json:"created_at"`
	Content   string    `json:"content"`
	Source    string    `json:"source"`
	Tags      []string  `json:"tags"`
}

func (n NewsPostRequestBody) Validate() (news store.News, errs error) {
	if n.Author == "" || n.Title == "" || n.Summary == "" || n.CreatedAt == "" || n.Content == "" || n.Source == "" || len(n.Tags) == 0 {
		return store.News{}, fmt.Errorf("invalid payload: one or more required fields are empty")
	}

	t, err := time.Parse(time.RFC3339, n.CreatedAt)
	if err != nil {
		return store.News{}, fmt.Errorf("invalid created_at format: %w", err)
	}

	parsedURL, err := url.ParseRequestURI(n.Source)
	if err != nil {
		return store.News{}, fmt.Errorf("invalid source URL format: %w", err)
	}

	return store.News{
		ID:        n.ID,
		Author:    n.Author,
		Title:     n.Title,
		Summary:   n.Summary,
		CreatedAt: t,
		Source:    parsedURL,
		Content:   n.Content,
		Tags:      n.Tags,
	}, nil
}

type AllNewsResponse struct {
	News []store.News `json:"news"`
}
