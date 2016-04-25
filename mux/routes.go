package mux

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetHighFives",
		"GET",
		"/highFive",
		GetHighFives,
	},
	Route{
		"GetHighFive",
		"GET",
		"/highFive/{highFiveId}",
		GetHighFive,
	},
	Route{
		"PostHighFive",
		"POST",
		"/highFive",
		PostHighFive,
	},
}
