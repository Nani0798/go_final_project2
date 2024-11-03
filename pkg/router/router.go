package router

import (
	"log"
	"net/http"

	"github.com/Nani0798/go_final_project/internal/scheduler"
	"github.com/Nani0798/go_final_project/internal/storage/sqlite"

	"github.com/Nani0798/go_final_project/internal/config"
	"github.com/Nani0798/go_final_project/internal/http-server/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() *chi.Mux {

	webDir := config.WebDirPath

	storage, err := sqlite.New(config.DBFilePath)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	planner := scheduler.NewScheduler(storage)

	router := chi.NewRouter()

	router.Use(middleware.URLFormat)

	router.Handle("/*", http.FileServer(http.Dir(webDir)))

	handlers.RegisterRoutes(router, planner)

	return router
}
