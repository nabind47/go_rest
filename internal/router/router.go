package router

import (
	"net/http"

	"github.com/nabind47/go_rest47/internal/handler"
)

// NEW CREATES A NEW ROUTER WITH ALL THE HANDLERS CONFIGURED!
func New(ns handler.NewsStorer) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("POST /news", handler.PostNews(ns))
	r.HandleFunc("GET /news", handler.GetNews())

	r.HandleFunc("GET /news/{id}", handler.GetNewsById())
	r.HandleFunc("PUT /news/{id}", handler.UpdateNewsById())
	r.HandleFunc("DELETE /news/{id}", handler.DeleteNewsById())
	return r
}
