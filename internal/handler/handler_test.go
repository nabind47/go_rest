package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nabind47/go_rest47/internal/handler"
)

func Test_PostNews(t *testing.T) {
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
			handler.PostNews()(w, r)

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
