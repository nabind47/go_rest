package handler

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

type NewsPostRequestBody struct {
	Author    string   `json:"author"`
	Title     string   `json:"title"`
	Summary   string   `json:"summary"`
	CreatedAt string   `json:"created_at"`
	Content   string   `json:"content"`
	Source    string   `json:"source"`
	Tags      []string `json:"tags"`
}

func (n NewsPostRequestBody) Validate() (errs error) {
	if n.Author == "" || n.Title == "" || n.Summary == "" || n.CreatedAt == "" || n.Content == "" || n.Source == "" || len(n.Tags) == 0 {
		return fmt.Errorf("invalid payload: one or more required fields are empty")
	}

	if _, err := time.Parse(time.RFC3339, n.CreatedAt); err != nil {
		errs = errors.Join(errs, fmt.Errorf("invalid created_at format: %w", err))
	}

	if _, err := url.ParseRequestURI(n.Source); err != nil {
		errs = errors.Join(errs, fmt.Errorf("invalid source URL: %w", err))
	}

	return errs
}

type AllNewsResponse struct {
	News []NewsPostRequestBody `json:"news"`
}
