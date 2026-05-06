package main

import (
	"log"
	"net/http"
	"os"

	"github.com/alvin/mvps/internal/cases"
	"github.com/alvin/mvps/internal/platform/httpserver"
)

func main() {
	repo := cases.NewInMemoryRepository()
	service := cases.NewService(repo)
	handler := httpserver.NewHandler(service, repo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
