package fiber

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/sing3demons/go-http-service/router"
)

type fiberRouter struct {
	*fiber.App
}

func NewMicroservice() router.IMicroservice {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(
		logger.Config{
			Format:     "${pid} ${status} - ${method} ${path}\n",
			TimeFormat: "02-Jan-2006",
			TimeZone:   "Asia/Bangkok"},
	))
	app.Use(recover.New())
	return &fiberRouter{app}
}
func (ms *fiberRouter) GET(path string, handlers ...router.HandlerFunc) {
	var h []func(*fiber.Ctx) error
	for index := range handlers {
		if index == 0 {
			h = append(h, func(c *fiber.Ctx) error {
				handlers[0](NewMyContext(c))
				return nil
			})
		} else {
			h = append(h, func(c *fiber.Ctx) error {
				handlers[index](NewMyContext(c))
				return nil
			})
		}
	}
	ms.App.Get(path, h...)
}
func (ms *fiberRouter) POST(path string, handlers ...router.HandlerFunc) {
	var h []func(*fiber.Ctx) error
	for index := range handlers {
		if index == 0 {
			h = append(h, func(c *fiber.Ctx) error {
				handlers[0](NewMyContext(c))
				return nil
			})
		} else {
			h = append(h, func(c *fiber.Ctx) error {
				handlers[index](NewMyContext(c))
				return nil
			})
		}
	}
	ms.App.Post(path, h...)
}

func (ms *fiberRouter) PUT(path string, handlers ...router.HandlerFunc) {
	var h []func(*fiber.Ctx) error
	for index := range handlers {
		if index == 0 {
			h = append(h, func(c *fiber.Ctx) error {
				handlers[0](NewMyContext(c))
				return nil
			})
		} else {
			h = append(h, func(c *fiber.Ctx) error {
				handlers[index](NewMyContext(c))
				return nil
			})
		}
	}
	ms.App.Put(path, h...)
}

func (ms *fiberRouter) PATCH(path string, handlers ...router.HandlerFunc) {
	var h []func(*fiber.Ctx) error
	for index := range handlers {
		if index == 0 {
			h = append(h, func(c *fiber.Ctx) error {
				handlers[0](NewMyContext(c))
				return nil
			})
		} else {
			h = append(h, func(c *fiber.Ctx) error {
				handlers[index](NewMyContext(c))
				return nil
			})
		}
	}
	ms.App.Patch(path, h...)
}
func (ms *fiberRouter) DELETE(path string, handlers ...router.HandlerFunc) {
	var h []func(*fiber.Ctx) error
	for index := range handlers {
		if index == 0 {
			h = append(h, func(c *fiber.Ctx) error {
				handlers[0](NewMyContext(c))
				return nil
			})
		} else {
			h = append(h, func(c *fiber.Ctx) error {
				handlers[index](NewMyContext(c))
				return nil
			})
		}
	}
	ms.App.Delete(path, h...)
}
func (ms *fiberRouter) Use(mwf ...router.HandlerFunc) {
	var h []func(*fiber.Ctx) error
	for index := range mwf {
		h = append(h, func(c *fiber.Ctx) error {
			mwf[index](NewMyContext(c))
			return c.Next()
		})
	}

	for index := range h {
		ms.App.Use(h[index])
	}

	// ms.App.Use(func(c *fiber.Ctx) error {
	// 	mwf[0](NewMyContext(c))
	// 	return c.Next()
	// })
}
func (ms *fiberRouter) StartHttp() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	//Graceful Shutdown
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := ms.App.Listen(":" + PORT); err != nil {
			fmt.Println("shutting down the server")
		}
	}()

	<-ctx.Done()
	stop()

	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	if err := ms.App.Shutdown(); err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}
}
