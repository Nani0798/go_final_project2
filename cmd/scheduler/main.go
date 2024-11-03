package main

import (
	"log"
	"net/http"

	"github.com/Nani0798/go_final_project/internal/config"
	"github.com/Nani0798/go_final_project/pkg/router"
)

func main() {
	config.MustLoad()

	port := ":" + config.Port

	router := router.SetupRouter()

	log.Printf("Server is running at %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	log.Fatalf("server stopped")
}
