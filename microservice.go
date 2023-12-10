package main

import (
	"github.com/sing3demons/go-http-service/router"
	"github.com/sing3demons/go-http-service/router/fiber"
	"github.com/sing3demons/go-http-service/router/gin"
	"github.com/sing3demons/go-http-service/router/mux"
)

func NewGinRouter() router.IMicroservice {
	return gin.NewMicroservice()
}

func NewMuxRouter() router.IMicroservice {
	return mux.NewMicroservice()
}

func NewFiberRouter() router.IMicroservice {
	return fiber.NewMicroservice()
}
