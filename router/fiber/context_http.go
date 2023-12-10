package fiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sing3demons/go-http-service/router/ctx"
)

type myContext struct {
	*fiber.Ctx
}

func NewMyContext(ctx *fiber.Ctx) ctx.IContext {
	return &myContext{ctx}
}

func (c *myContext) JSON(code int, obj any) {
	c.Ctx.Status(code).JSON(obj)
}

func (c *myContext) RequestURI() string {
	return string(c.Ctx.Request().RequestURI())
}

func (c *myContext) Bind(code int, obj any) {
	c.Ctx.BodyParser(obj)
}
func (c *myContext) BodyParser(obj any) error {
	return c.Ctx.BodyParser(obj)
}
func (c *myContext) ReadBody() ([]byte, error) {
	return c.Ctx.Body(), nil
}
func (c *myContext) Param(key string) string {
	return c.Ctx.Params(key)
}
func (c *myContext) Query(key string) string {
	return c.Ctx.Query(key)
}

