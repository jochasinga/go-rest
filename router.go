package main

import (
//	"net/http"

	mux "github.com/julienschmidt/httprouter"
)

func NewRouter() *mux.Router {

	router := mux.New()

	for _, route := range routes {
		
		//var handler mux.Handle

		//handler = Logger(router, route.Method, route.Name, route.Pattern, route.Handle)
		Logger(router, route.Method, route.Name, route.Pattern, route.Handle)
		//router.Handle(route.Method, route.Pattern, route.Handle)
	}

	return router
}
