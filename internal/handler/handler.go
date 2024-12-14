package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/nabind47/go_rest47/internal/logger"
	"github.com/nabind47/go_rest47/internal/store"
)

type NewsStorer interface {
	Create(store.News) (store.News, error)
	FindAll() ([]store.News, error)
	FindByID(uuid.UUID) (store.News, error)

	UpdateByID(store.News) error
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

		n, err := body.Validate()
		if err != nil {
			logger.Error("invalid payload", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if _, err := ns.Create(n); err != nil {
			logger.Error("failed to create news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func GetNews(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		news, err := ns.FindAll()

		if err != nil {
			logger.Error("failed to get news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		data := AllNewsResponse{News: news}
		if err := json.NewEncoder(w).Encode(data); err != nil {
			logger.Error("failed to encode response", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func GetNewsById(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		id := r.PathValue("id")
		newsUUID, err := uuid.Parse(id)

		if err != nil {
			logger.Error("failed to parse uuid", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		news, err := ns.FindByID(newsUUID)
		if err != nil {
			logger.Error("failed to get news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if err := json.NewEncoder(w).Encode(news); err != nil {
			logger.Error("failed to encode response", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}

func UpdateNewsById(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())

		var body NewsPostRequestBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			logger.Error("failed to decode request body", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		n, err := body.Validate()
		if err != nil {
			logger.Error("invalid payload", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if err := ns.UpdateByID(n); err != nil {
			logger.Error("failed to update news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}
}
func DeleteNewsById(ns NewsStorer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := logger.FromContext(r.Context())
		id := r.PathValue("id")
		newsUUID, err := uuid.Parse(id)

		if err != nil {
			logger.Error("failed to parse uuid", "error", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if err := ns.DeleteByID(newsUUID); err != nil {
			logger.Error("failed to delete news", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusNoContent)

	}
}
