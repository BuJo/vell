package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

// NewRouter instanciates a HTTP Handler for Vell endpoints.
func NewRouter() http.Handler {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := withLogging(route.HandlerFunc, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
