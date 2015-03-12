package main

import (
//	"net/http"

	mux "github.com/julienschmidt/httprouter"
)

func NewRouter() *mux.Router {

	router := mux.New()

	for _, route := range routes {
		
                /* TODO: Implement logging 
		var handler mux.Handle

		handler = route.Handle
		handler = Logger(handler, route.Name)
                */

		router.Handle(route.Method, route.Pattern, route.Handle)
	}

	return router
}
