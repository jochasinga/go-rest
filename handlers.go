package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	mux "github.com/julienschmidt/httprouter"
)

// Logger logs the method, URI, header, and the dispatch time of the request.
func Logger(r *http.Request) {
	start := time.Now()
	log.Printf(
		"%s\t%s\t%q\t%s",
		r.Method,
		r.RequestURI,
		r.Header,
		time.Since(start),
	)
}

// Index handler handles the index at "/" and writes a welcome message.
func Index(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	fmt.Fprintf(w, "<h1 style=\"font-family: Helvetica;\">Hello, welcome to blog service</h1>")
}

// PostIndex handler handles "/posts" and show all the blog posts data as JSON
func PostIndex(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	posts := FindAll()
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		panic(err)
	}
}

// PostShow handler shows the post at "posts/id" as JSON.
func PostShow(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	id, err := strconv.Atoi(ps.ByName("postId"))
	HandleError(err)
	post := FindPost(id)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		panic(err)
	}
}

// PostCreate creates a new post data
func PostCreate(w http.ResponseWriter, r *http.Request, _ mux.Params) {
	Logger(r)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	HandleError(err)
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	// Save JSON to Post struct (should this be a pointer?)
	var post Post
	if err := json.Unmarshal(body, &post); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
	}

	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}

	CreatePost(post)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// PostDelete handler deletes the blog post.
func PostDelete(w http.ResponseWriter, r *http.Request, ps mux.Params) {
	id, err := strconv.Atoi(ps.ByName("postId"))
	HandleError(err)
	DeletePost(id)
}
