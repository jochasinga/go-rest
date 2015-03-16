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

	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to blog service</h1>")
}

func PostIndex(w http.ResponseWriter, r *http.Request, _ mux.Params) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	
	var posts Posts
	
	posts = FindAll()

	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}

}

func PostShow(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	
	id, err := strconv.Atoi(ps.ByName("postId"))
	
	HandleError(err)

	post := FindPost(id)
	
	if err := json.NewEncoder(w).Encode(post); err != nil {
		panic(err)
	}
}

func PostCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	
	Logger(r)
	
	var post Post
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Save JSON to Post struct
	if err := json.Unmarshal(body, &post); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	CreatePost(post)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func PostDelete(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	
	id, err := strconv.Atoi(ps.ByName("postId"))
	HandleError(err)

	DeletePost(id)
}

