package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Q           map[string]string
	handlerFunc func(http.ResponseWriter, *http.Request)
)

func RegisterHandler(
	router *mux.Router,
	method string,
	path string,
	query Q,
	handler handlerFunc,
	middlewares ...mux.MiddlewareFunc, // optional middlewares
) *mux.Route {

	route := router.Path(path)

	if method != "" && method != "*" {
		route = route.Methods(method)
	}

	for key, value := range query {
		route.Queries(key, value)
	}

	// specify middlewares (if any) on a dedicated sub-Router
	if len(middlewares) > 0 {
		router = route.Subrouter()
		router.Use(middlewares...)
		route = router.NewRoute()
	}

	if handler != nil {
		route = route.HandlerFunc(handler)
	}

	return route

}
