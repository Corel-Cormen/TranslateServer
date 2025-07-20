package ServerCore

import (
	"TranslateServer/internal/ServerPlatform/api"
	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	engine *gin.Engine
}

func NewGinRouter() *GinRouter {
	return &GinRouter{
		engine: gin.Default(),
	}
}

func (r *GinRouter) Run(addr ...string) error {
	return r.engine.Run(addr...)
}

func (r *GinRouter) GET(path string, handler func(c ServerCoreApi.HandlerInterface)) {
	r.engine.GET(path, func(c *gin.Context) {
		ginContext := &GinContextHandler{Context: c}
		handler(ginContext)
	})
}
