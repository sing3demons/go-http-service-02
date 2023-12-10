package gin

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
	"github.com/sing3demons/go-http-service/router"
)

type ginRouter struct {
	*gin.Engine
}

func NewMicroservice() router.IMicroservice {
	r := gin.Default()
	return &ginRouter{r}
}

func (r *ginRouter) GET(path string, handlers ...router.HandlerFunc) {
	var h []gin.HandlerFunc
	for index := range handlers {
		if index == 0 {
			h = append(h, func(ctx *gin.Context) {
				handlers[0](NewMyContext(ctx))
			})
		} else {
			h = append(h, func(ctx *gin.Context) {
				handlers[index](NewMyContext(ctx))
				ctx.Next()
			})
		}
	}
	r.Engine.GET(path, h...)
	// r.Engine.GET(path, func(ctx *gin.Context) {
	// 	h(NewMyContext(ctx))
	// })
}

func (r *ginRouter) POST(path string, handlers ...router.HandlerFunc) {
	var h []gin.HandlerFunc
	for index := range handlers {
		if index == 0 {
			h = append(h, func(ctx *gin.Context) {
				handlers[0](NewMyContext(ctx))
			})
		} else {
			h = append(h, func(ctx *gin.Context) {
				handlers[index](NewMyContext(ctx))
				ctx.Next()
			})
		}
	}
	r.Engine.POST(path, h...)
	// r.Engine.POST(path, func(ctx *gin.Context) {
	// 	h(NewMyContext(ctx))
	// })
}

func (r *ginRouter) PUT(path string, handlers ...router.HandlerFunc) {
	var h []gin.HandlerFunc
	for index := range handlers {
		if index == 0 {
			h = append(h, func(ctx *gin.Context) {
				handlers[0](NewMyContext(ctx))
			})
		} else {
			h = append(h, func(ctx *gin.Context) {
				handlers[index](NewMyContext(ctx))
				ctx.Next()
			})
		}
	}
	r.Engine.PUT(path, h...)
	// r.Engine.PUT(path, func(ctx *gin.Context) {
	// 	h(NewMyContext(ctx))
	// })
}

func (r *ginRouter) PATCH(path string, handlers ...router.HandlerFunc) {
	var h []gin.HandlerFunc
	for index := range handlers {
		if index == 0 {
			h = append(h, func(ctx *gin.Context) {
				handlers[0](NewMyContext(ctx))
			})
		} else {
			h = append(h, func(ctx *gin.Context) {
				handlers[index](NewMyContext(ctx))
				ctx.Next()
			})
		}
	}
	r.Engine.PATCH(path, h...)
	// r.Engine.PATCH(path, func(ctx *gin.Context) {
	// 	h(NewMyContext(ctx))
	// })
}

func (r *ginRouter) DELETE(path string, handlers ...router.HandlerFunc) {
	var h []gin.HandlerFunc
	for index := range handlers {
		if index == 0 {
			h = append(h, func(ctx *gin.Context) {
				handlers[0](NewMyContext(ctx))
			})
		} else {
			h = append(h, func(ctx *gin.Context) {
				handlers[index](NewMyContext(ctx))
				ctx.Next()
			})
		}
	}
	r.Engine.DELETE(path, h...)
	// r.Engine.DELETE(path, func(ctx *gin.Context) {
	// 	h(NewMyContext(ctx))
	// })
}

func (r *ginRouter) Use(mwf ...router.HandlerFunc) {
	var h []gin.HandlerFunc
	for index := range mwf {
		h = append(h, func(ctx *gin.Context) {
			mwf[index](NewMyContext(ctx))
			ctx.Next()
		})
	}
	r.Engine.Use(h...)
}

func (r *ginRouter) StartHttp() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	sv := &http.Server{
		Addr:    ":" + PORT,
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
