package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html"

	mux "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintln(w, "Todo Index!")

	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	fmt.Fprintf(w, "Todo show: %s\n", ps.ByName("todoId"))
}
