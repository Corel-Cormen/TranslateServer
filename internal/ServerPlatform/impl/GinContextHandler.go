package ServerCore

import (
	"github.com/gin-gonic/gin"
)

type GinContextHandler struct {
	*gin.Context
}

func (c *GinContextHandler) TextCallback(code int, obj interface{}) {
	c.Header("Content-Type", "text/plain")
	c.String(code, obj.(string))
}

func (c *GinContextHandler) JsonCallback(code int, obj interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(code, obj)
}

func (c *GinContextHandler) BindJSON(obj interface{}) error {
	return c.Context.ShouldBindJSON(obj)
}
