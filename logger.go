package main


import (
	"log"
	//"net/http"
	"time"
	 mux "github.com/julienschmidt/httprouter"
)

// Wrapper function for server logging
func Logger(router *mux.Router, method, name, path string, handle mux.Handle) mux.Handle {
	
		router.Handle(method, path, handle)
		
		start := time.Now()

		//inner.ServeHTTP(w, r)
		//router.Handle(route.Method, route.Pattern, route.Handle)
	    //router.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			method,
			path,
			name,
			time.Since(start),
		)
		
		return handle
}

