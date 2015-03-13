package main

import (
	"encoding/json"
	"fmt"
	"net/http"
//	"html"
	"io"
	"io/ioutil"
	"strconv"
	"log"
	"time"
	
	mux "github.com/julienschmidt/httprouter"
)

func Logger(r *http.Request) {
	
	start:= time.Now()
	
	log.Printf(
		"%s\t%s\t%q\t%s",
		r.Method,
		r.RequestURI,
		r.Header,
		time.Since(start),
	)
}

func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	
	Logger(r)

	fmt.Fprintf(w, "<h1>Hello, welcome to my blog</h1>")
	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	//fmt.Fprintf(w, "Hello, %s\n", p.ByName("anything"))
}

func PostIndex(w http.ResponseWriter, r *http.Request, _ mux.Params) {

	Logger(r)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}
}

func PostShow(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	
	Logger(r)
	
	//fmt.Fprintf(w, "Todo show: %s\n", ps.ByName("todoId"))

	for _, todo := range posts {

		id, err := strconv.Atoi(ps.ByName("todoId"))
		if err != nil {
			panic(err)
		}
		
		if todo.Id == id {
			if err = json.NewEncoder(w).Encode(todo); err != nil {
				panic(err)
			}
			return
		}
	}

}

func PostCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	
	Logger(r)
	
	var post Post
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Save JSON to Todo struct
	if err := json.Unmarshal(body, &post); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := CreatePost(post)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

func PostDownload(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	
	Logger(r)
	
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}

	/* TODO: Write to file and download to user
	if enc, err := json.NewEncoder(w)
	if err != nil {
		panic(err)
	}
        */
	
	b, err := json.Marshal(posts) 
	if err != nil {
		panic(err)
	}
	
	if err = ioutil.WriteFile("todos.json", b, 0777); err != nil {
		panic(err)
	}
	
	fmt.Println(b)
}
