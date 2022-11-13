package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents all routes for api
type Route struct {
	URI             string
	Method          string
	Function        func(http.ResponseWriter, *http.Request)
	IsAuthenticated bool
}

// Configure inserts all routes in the router
func Configure(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
