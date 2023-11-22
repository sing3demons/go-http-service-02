package ms

import "github.com/gin-gonic/gin"

type IMyContext interface {
	JSON(code int, obj any)
}

type myContext struct {
	*gin.Context
}

func NewMyContext(ctx *gin.Context) IMyContext {
	return &myContext{ctx}
}

func (c *myContext) JSON(code int, obj any) {
	c.Context.JSON(code, obj)
}

func (c *myContext) Bind(code int, obj any) {
	c.Context.ShouldBind(obj)
}
