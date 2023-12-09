package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/sing3demons/go-http-service/router/ctx"
)

type myContext struct {
	*gin.Context
}

func NewMyContext(ctx *gin.Context) ctx.IContext {
	return &myContext{ctx}
}

func (c *myContext) JSON(code int, obj any) {
	c.Context.JSON(code, obj)
}

func (c *myContext) RequestURI() string {
	return c.Context.Request.RequestURI
}

func (c *myContext) Bind(code int, obj any) {
	c.Context.ShouldBind(obj)
}
func (c *myContext) BodyParser(obj any) error {
	return c.Context.ShouldBind(obj)
}
func (c *myContext) ReadBody() ([]byte, error) {
	return c.Context.GetRawData()
}
func (c *myContext) Param(key string) string {
	return c.Context.Param(key)
}
func (c *myContext) Query(key string) string {
	return c.Context.Query(key)
}

func (c *myContext) Next() {
	c.Context.Next()
}
