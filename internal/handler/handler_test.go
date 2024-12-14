package handler_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/nabind47/go_rest47/internal/handler"
)

func Test_PostNews(t *testing.T) {
	validJSON := `{
		"author": "Author", 
		"title": "Title", 
		"summary": "Summary", 
		"content": "Content", 
		"source": "http://example.com", 
		"created_at": "2021-01-01T00:00:00Z", 
		"tags": ["tag1"]
	}`
	missingCreatedAtJSON := `{
		"author": "Author", 
		"title": "Title", 
		"summary": "Summary", 
		"content": "Content", 
		"source": "http://example.com", 
		"tags": ["tag1"]
	}`

	testCases := []struct {
		name           string
		store          handler.NewsStorer
		body           io.Reader
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Invalid JSON payload",
			store:          &mockNewsStorer{},
			body:           strings.NewReader(`{`),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "Missing required field (created_at)",
			store:          &mockNewsStorer{},
			body:           strings.NewReader(missingCreatedAtJSON),
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "Syncing error (store failure)",
			store: &mockNewsStorer{
				err: true,
			},
			body:           strings.NewReader(validJSON),
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "Success",
			store:          &mockNewsStorer{},
			body:           strings.NewReader(validJSON),
			expectedStatus: http.StatusCreated,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// ARRANGE
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", tc.body)

			// ACT
			handler.PostNews(tc.store)(w, r)

			// ASSERT
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
func Test_GetNews(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// ARRANGE
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// ACT
			handler.GetNews()(w, r)

			// ASSERT
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
func Test_GetNewsById(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// ARRANGE
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// ACT
			handler.GetNewsById()(w, r)

			// ASSERT
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
func Test_UpdateNewsById(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// ARRANGE
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// ACT
			handler.UpdateNewsById()(w, r)

			// ASSERT
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}
func Test_DeleteNewsById(t *testing.T) {
	testCases := []struct {
		name           string
		expectedStatus int
	}{
		{
			name:           "not implemented",
			expectedStatus: http.StatusNotImplemented,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// ARRANGE
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, "/", nil)

			// ACT
			handler.DeleteNewsById()(w, r)

			// ASSERT
			if w.Result().StatusCode != tc.expectedStatus {
				t.Errorf("expected status %d, got %d", tc.expectedStatus, w.Result().StatusCode)
			}
		})
	}
}

type mockNewsStorer struct {
	err bool
}

func (m mockNewsStorer) Create(_ handler.NewsPostRequestBody) (news handler.NewsPostRequestBody, err error) {
	if m.err {
		return news, errors.New("some error")
	}
	return news, nil
}
func (m mockNewsStorer) FindByID(_ uuid.UUID) (news handler.NewsPostRequestBody, err error) {
	if m.err {
		return news, errors.New("some error")
	}
	return news, nil
}
func (m mockNewsStorer) FindAll() (news []handler.NewsPostRequestBody, err error) {
	if m.err {
		return news, errors.New("some error")
	}
	return news, nil
}
func (m mockNewsStorer) UpdateByID(_ uuid.UUID) error {
	if m.err {
		return errors.New("some error")
	}
	return nil
}
func (m mockNewsStorer) DeleteByID(_ uuid.UUID) error {
	if m.err {
		return errors.New("some error")
	}
	return nil
}
