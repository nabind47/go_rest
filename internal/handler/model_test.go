package handler_test

import (
	"net/url"
	"testing"
	"time"

	"github.com/nabind47/go_rest47/internal/handler"
	"github.com/nabind47/go_rest47/internal/store"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewsPostRequestBody_Validate(t *testing.T) {
	type expectations struct {
		err  string
		news store.News
	}

	testCases := []struct {
		name         string
		body         handler.NewsPostRequestBody
		expectations expectations
	}{
		{
			name: "All fields empty - expect error",
			body: handler.NewsPostRequestBody{},
			expectations: expectations{
				err: "invalid payload: one or more required fields are empty",
			},
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
			expectations: expectations{
				err: `invalid created_at format: parsing time "InvalidDate"`,
			},
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
			expectations: expectations{
				err: `invalid source URL format: parse "invalid_url"`,
			},
		},
		// {
		// 	name: "Valid payload - no error expected",
		// 	body: handler.NewsPostRequestBody{
		// 		Author:    "Author",
		// 		Title:     "Title",
		// 		Summary:   "Summary",
		// 		CreatedAt: "2024-12-14T15:04:05Z",
		// 		Content:   "Content",
		// 		Source:    "http://example.com",
		// 		Tags:      []string{"tag1", "tag2"},
		// 	},
		// 	expectations: expectations{
		// 		news: store.News{
		// 			ID:      uuid.New(),
		// 			Author:  "Author",
		// 			Title:   "Title",
		// 			Summary: "Summary",
		// 			Content: "Content",
		// 			Tags:    []string{"tag1", "tag2"},
		// 		},
		// 	},
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			news, err := tc.body.Validate()

			if tc.expectations.err != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectations.err)
			} else {
				assert.NoError(t, err)

				parsedTime, parsedErr := time.Parse(time.RFC3339, tc.body.CreatedAt)
				require.NoError(t, parsedErr)

				parsedURL, parsedURLErr := url.ParseRequestURI(tc.body.Source)
				require.NoError(t, parsedURLErr)

				tc.expectations.news.CreatedAt = parsedTime
				tc.expectations.news.Source = parsedURL
				assert.Contains(t, tc.expectations.news, news)
			}
		})
	}
}
