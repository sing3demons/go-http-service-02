package main

import (
	"github.com/gorilla/mux"
	"github.com/sing3demons/go-http-service/router/ctx"
)

type IMicroservice interface {
	GET(path string, h handlerFunc, mwf ...mux.MiddlewareFunc)
	POST(path string, h handlerFunc, mwf ...mux.MiddlewareFunc)
	PUT(path string, h handlerFunc, mwf ...mux.MiddlewareFunc)
	PATCH(path string, h handlerFunc, mwf ...mux.MiddlewareFunc)
	DELETE(path string, h handlerFunc, mwf ...mux.MiddlewareFunc)
	Use(mwf ...mux.MiddlewareFunc)
	Group(path string) *mux.Router
	StartHttp()
}

type handlerFunc func(c ctx.IContext)
