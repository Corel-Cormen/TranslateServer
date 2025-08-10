package ServerPlatformInstance

import (
	"sync"

	"TranslateServer/internal/ServerPlatform/api"
	"TranslateServer/internal/ServerPlatform/impl"
)

var (
	serverInstance ServerCoreApi.ServerInterface
	routerInstance ServerCoreApi.RouterInterface
	onceServer     sync.Once
	onceRouter     sync.Once
)

func getRouter() ServerCoreApi.RouterInterface {
	onceRouter.Do(func() {
		routerInstance = ServerCore.NewGinRouter()
	})
	return routerInstance
}

func GetServer() ServerCoreApi.ServerInterface {
	onceServer.Do(func() {
		addr := "0.0.0.0"
		port := 5000
		serverInstance = ServerCore.NewServer(addr, port, getRouter())
	})
	return serverInstance
}
