package routes

import (
	"api/internal/controller"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:             "/users",
		Method:          http.MethodPost,
		Function:        controller.Create,
		IsAuthenticated: false,
	},
	{
		URI:             "/users",
		Method:          http.MethodGet,
		Function:        controller.FindAll,
		IsAuthenticated: false,
	},
	{
		URI:             "/users/{id}",
		Method:          http.MethodGet,
		Function:        controller.FindById,
		IsAuthenticated: false,
	},
	{
		URI:             "/users/{id}",
		Method:          http.MethodPut,
		Function:        controller.Update,
		IsAuthenticated: false,
	},
	{
		URI:             "/users/{id}",
		Method:          http.MethodDelete,
		Function:        controller.Delete,
		IsAuthenticated: false,
	},
}
