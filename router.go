package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(repo repo) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = handlerWithRepo(repo, route.HandlerFunc)
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func handlerWithRepo(repo repo, handlerFunc func(repo, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handlerFunc(repo, w, r)
	}
}
