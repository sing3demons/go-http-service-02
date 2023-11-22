package main

import (
	"log"
	"net/http"
)

func Home(c IContext) {
	c.JSON(200, map[string]any{
		"msg":  "hello world",
		"name": c.Query("name"),
	})
}

func main() {
	r := NewMicroservice()
	r.Use(loggingMiddleware)

	r.GET("/{id}", func(c IContext) {
		c.JSON(http.StatusOK, map[string]any{
			"id": c.Param("id"),
		})
	})

	r.GET("/", Home)

	r.POST("/post", func(c IContext) {
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

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
