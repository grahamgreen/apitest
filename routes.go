package main

//Copyright 2015 Graham Green

import "net/http"

//Route the route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes slice of Routes
type Routes []Route

//router.HandleFunc("/", Index)
//router.HandleFunc("/add/account/", AddAccount)
//router.HandleFunc("/add/action/", AddAction)

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AddAccount",
		"POST",
		"/add/accout/",
		AddAccount,
	},
	Route{
		"AddAction",
		"POST",
		"/add/action/",
		AddAccount,
	},
	Route{
		"Donate",
		"POST",
		"/donate/",
		Donate,
	},
}
