package main

import (
//	"net/http"
	 mux "github.com/julienschmidt/httprouter"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Handle	    mux.Handle
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/posts",
		PostIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/posts/:id",
		PostShow,
	},
	Route{
		"TodoCreate",
		"POST",
		"/posts",
		PostCreate,
	},
	Route{
		"TodoDownload",
		"GET",
		"/posts.json",
		PostDownload,
	},
}
