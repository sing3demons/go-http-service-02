package ms

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type IRouter interface {
	GET(path string, h handlerFunc)

	StartHttp()
}

type ginRouter struct {
	*gin.Engine
}

func NewRouter() IRouter {
	r := gin.Default()
	return &ginRouter{r}
}

type handlerFunc func(c IMyContext)

func (r *ginRouter) GET(path string, h handlerFunc) {
	r.Engine.GET(path, func(ctx *gin.Context) {
		h(NewMyContext(ctx))
	})
}

func (r *ginRouter) StartHttp() {
	port := "8080"
	sv := &http.Server{
		Addr:    ":" + port,
		Handler: r.Engine,
	}

	go func() {
		fmt.Printf("http listen: %s\n", sv.Addr)

		if err := sv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server listen err: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := sv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown: ", err)
	}
	fmt.Println("server exited")
}
