package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

func Config(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	return r
}
