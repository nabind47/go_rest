package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/nabind47/go_rest47/internal/logger"
)

type NewsStorer interface {
	Create(NewsPostRequestBody) (NewsPostRequestBody, error)
	FindAll() ([]NewsPostRequestBody, error)
	FindByID(uuid.UUID) (NewsPostRequestBody, error)

	UpdateByID(uuid.UUID) error
	DeleteByID(uuid.UUID) error
}

func PostNews(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())

		var body NewsPostRequestBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			logger.Error("failed to decode request body", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := body.Validate(); err != nil {
			logger.Error("invalid payload", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if _, err := ns.Create(body); err != nil {
			logger.Error("failed to create news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
func GetNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
func GetNewsById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
func UpdateNewsById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
func DeleteNewsById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}
