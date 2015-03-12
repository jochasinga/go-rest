package main

import (
	
       "net/http"
	   
	   mux "github.com/julienschmidt/httprouter"
)

func NewRouter() *mux.Router {
	
	router := mux.New()
	
	for _, route := range routes {
		
		var handler http.Handler
	
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	
	return router
}






