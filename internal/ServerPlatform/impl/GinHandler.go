package ServerCore

import (
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	*gin.Context
}

func (c *GinContext) Callback(code int, obj interface{}) {
	c.Header("Content-Type", "application/text")
	c.String(code, obj.(string))
}
