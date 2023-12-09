package main

import (
	"net/http"

	"github.com/sing3demons/go-http-service/router/ctx"
	"github.com/sing3demons/go-http-service/router/gin"
)

func main() {
	// r := mux.NewMicroservice()
	r := gin.NewMicroservice()

	r.GET("/:id", func(c ctx.IContext) {
		c.JSON(http.StatusOK, map[string]any{
			"id": c.Param("id"),
		})
	})

	r.GET("/", Home)

	r.POST("/post", func(c ctx.IContext) {
		var req struct {
			Name string `json:"name"`
		}
		if err := c.BodyParser(&req); err != nil {
			c.JSON(http.StatusBadRequest, map[string]any{"msg": err.Error()})
			return
		}
		c.JSON(http.StatusOK, req)
	})
	r.StartHttp()
}

func Home(c ctx.IContext) {
	c.JSON(200, map[string]any{
		"msg":  "hello world",
		"name": c.Query("name"),
	})
}
