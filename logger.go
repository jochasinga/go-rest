package main

/* TODO: Implement logging
import (
	"log"
	"net/http"
	"time"
	 mux "github.com/julienschmidt/httprouter"
)

func Logger(inner mux.Router, name string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)

	})
}
*/
