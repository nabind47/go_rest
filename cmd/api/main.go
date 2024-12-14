package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/nabind47/go_rest47/internal/router"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
	logger.Info("server is up :8080")

	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("failed to start server", "error", err)
	}
}
