package router

import (
	"github.com/sing3demons/go-http-service/router/ctx"
)

type IMicroservice interface {
	GET(path string, h ...HandlerFunc)
	POST(path string, h ...HandlerFunc)
	PUT(path string, h ...HandlerFunc)
	PATCH(path string, h ...HandlerFunc)
	DELETE(path string, h ...HandlerFunc)
	Use(mwf ...HandlerFunc)
	StartHttp()
}

type HandlerFunc func(c ctx.IContext)
