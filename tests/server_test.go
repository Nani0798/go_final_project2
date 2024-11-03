package tests

import (
	"net/http/httptest"

	"github.com/Nani0798/go_final_project/internal/config"
	"github.com/Nani0798/go_final_project/pkg/router"

	_ "github.com/mattn/go-sqlite3"
)

func createTestServer() *httptest.Server {
	config.MustLoad()

	router := router.SetupRouter()

	return httptest.NewServer(router)
}
