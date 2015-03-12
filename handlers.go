package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"html"
	"io"
	"io/ioutil"
	
	"log"
	"time"
	
	mux "github.com/julienschmidt/httprouter"
)

func Logger(r *http.Request) {
	
	start:= time.Now()
	
	log.Printf(
		"%s\t%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		//name,
		time.Since(start),
	)
}

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	
	Logger(r)
	
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ mux.Params) {

	Logger(r)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoShow(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	
	Logger(r)
	
	fmt.Fprintf(w, "Todo show: %s\n", ps.ByName("todoId"))
}

func TodoCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	
	Logger(r)
	
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Save JSON to Todo struct
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func TodoDownload(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	
	Logger(r)
	
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}

	/* TODO: Write to file and download to user
	if enc, err := json.NewEncoder(w)
	if err != nil {
		panic(err)
	}
        */
	
	b, err := json.Marshal(todos) 
	if err != nil {
		panic(err)
	}
	
	fmt.Println(b)
}
