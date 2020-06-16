package main

import "github.com/julienschmidt/httprouter"

func routerFactory() *httprouter.Router {
	r := httprouter.New()
	r.GET("/route/:key", getRoute)
	r.POST("/route", postRoute)
	r.GET("/statusReport", getStatusReport)

	return r
}
