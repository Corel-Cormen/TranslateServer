package ServerCore

import (
	"github.com/gin-gonic/gin"
)

type GinContextHandler struct {
	*gin.Context
}

func (c *GinContextHandler) Callback(code int, obj interface{}) {
	c.Header("Content-Type", "application/text")
	c.String(code, obj.(string))
}
