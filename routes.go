package main

import (
	
       "net/http"
	   
	   mux "github.com/julienschmidt/httprouter"
)

type Route struct {
	Name, Method, Pattern 	string
	HandlerFunc 			http.HandlerFunc
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
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/:todoId",
		TodoShow,
	},
}




