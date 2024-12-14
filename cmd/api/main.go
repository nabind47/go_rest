package main

import (
	"log"
	"net/http"

	"github.com/nabind47/go_rest47/internal/router"
)

func main() {
	r := router.New()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
