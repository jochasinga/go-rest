package main

/*
import (
	"log"
	"net/http"
	"time"
	 mux "github.com/julienschmidt/httprouter"
)

// Wrapper function for server logging
func Logger(router *mux.Router, method, path, name string, handle mux.Handle) mux.Handle {
	
	return router.Handle(method, path, mux.Handle(func(w http.ResponseWriter, r *http.Request) {
		
		start := time.Now()

		router.ServeHTTP(w, r)
		//router.Handle(route.Method, route.Pattern, route.Handle)
	    //router.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			method,
			path,
			name,
			time.Since(start),
		)
	}),
)}
*/

