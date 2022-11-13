package router

import (
	"api/internal/router/routes"

	"github.com/gorilla/mux"
)

// Generate will return a router with the configured routes
func Generate() *mux.Router {
	return routes.Configure(mux.NewRouter())
}
