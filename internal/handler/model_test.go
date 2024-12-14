package handler_test

import (
	"testing"

	"github.com/nabind47/go_rest47/internal/handler"
)

func TestNewsPostRequestBody_Validate(t *testing.T) {
	testCases := []struct {
		name     string
		body     handler.NewsPostRequestBody
		expected bool
	}{
		{
			name:     "All fields empty - expect error",
			body:     handler.NewsPostRequestBody{},
			expected: true,
		},
		{
			name: "Invalid CreatedAt format - expect error",
			body: handler.NewsPostRequestBody{
				Author:    "Author",
				Title:     "Title",
				Summary:   "Summary",
				CreatedAt: "InvalidDate",
				Content:   "Content",
				Source:    "http://example.com",
				Tags:      []string{"tag1"},
			},
			expected: true,
		},
		{
			name: "Invalid Source URL - expect error",
			body: handler.NewsPostRequestBody{
				Author:    "Author",
				Title:     "Title",
				Summary:   "Summary",
				CreatedAt: "2024-12-14T15:04:05Z",
				Content:   "Content",
				Source:    "invalid_url",
				Tags:      []string{"tag1"},
			},
			expected: true,
		},
		{
			name: "Valid payload - no error expected",
			body: handler.NewsPostRequestBody{
				Author:    "Author",
				Title:     "Title",
				Summary:   "Summary",
				CreatedAt: "2024-12-14T15:04:05Z",
				Content:   "Content",
				Source:    "http://example.com",
				Tags:      []string{"tag1", "tag2"},
			},
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.body.Validate()

			if tc.expected && err == nil {
				t.Fatalf("expected error but got nil")
			}

			if !tc.expected && err != nil {
				t.Fatalf("expected no error but got: %v", err)
			}
		})
	}
}
