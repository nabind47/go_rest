package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/nabind47/go_rest47/internal/logger"
	"github.com/nabind47/go_rest47/internal/router"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	r := router.New(nil)
	wrappedRouter := logger.AddLoggerMiddleware(log, logger.LoggerMiddleware(r))

	log.Info("server is up :8080")

	if err := http.ListenAndServe(":8080", wrappedRouter); err != nil {
		log.Error("failed to start server", "error", err)
	}
}
