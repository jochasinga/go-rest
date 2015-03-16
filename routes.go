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
	/*n
	Route{
		"PostDownload",
		"GET",
		"/posts.json",
		PostDownload,
	},
	*/
}
