package ServerPlatformInstance

import (
	"sync"
	"TranslateServer/internal/ServerPlatform/impl"
	"TranslateServer/internal/ServerPlatform/api"
)

var (
	serverInstance ServerCoreApi.ServerInterface
	routerInstance ServerCoreApi.RoutherInterface
	onceServer sync.Once
	onceRouter sync.Once
)

func getRouter() ServerCoreApi.RoutherInterface {
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
