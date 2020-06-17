package main

import (
	"github.com/EmiSan1998/gpsTiming-backend/handlers"
	"github.com/julienschmidt/httprouter"
)

func getRouter() *httprouter.Router {
	r := httprouter.New()
	r.GET("/route/:key", handlers.GetRoute)
	r.POST("/route", handlers.PostRoute)
	r.GET("/statusReport", handlers.GetStatusReport)

	return r
}
