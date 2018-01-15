package main

import (
	//	"net/http"
	mux "github.com/julienschmidt/httprouter"
)

// Route represents a URL that serves a specific resource.
type Route struct {
	Name    string
	Method  string
	Pattern string
	Handle  mux.Handle
}

// Routes are a list of Routes for this application.
type Routes []Route

var routes = Routes{
	// It is a good idea to name the fields when declaring a struct object.
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"PostIndex",
		"GET",
		"/posts",
		PostIndex,
	},
	Route{
		"PostShow",
		"GET",
		"/posts/:postId",
		PostShow,
	},
	Route{
		"PostCreate",
		"POST",
		"/posts",
		PostCreate,
	},
	Route{
		"PostDelete",
		"POST",
		"/posts/del/:postId",
		PostDelete,
	},
	/*
		Route{
			"PostDownload",
			"GET",
			"/posts.json",
			PostDownload,
		},
	*/
}
