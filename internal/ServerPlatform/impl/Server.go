package ServerCore

import (
	"fmt"
	"TranslateServer/internal/ServerPlatform/api"
)

type Server struct {
	address string
	port    int
	router 	ServerCoreApi.RoutherInterface
}

func NewServer(address string, port int, router ServerCoreApi.RoutherInterface) ServerCoreApi.ServerInterface {
	return &Server{
		address: address,
		port:    port,
		router:  router,
	}
}

func wrapHandler(handler ServerCoreApi.RouterHandlerInterface) func(ServerCoreApi.HandlerInterface) {
	return func(h ServerCoreApi.HandlerInterface) {
		handler.Handle(h)
	}
}

func (s *Server) Start() error {
	s.router.GET("/echo", wrapHandler(&EchoHandler{}))

	serverAddr := fmt.Sprintf("%s:%d", s.address, s.port)
	return s.router.Run(serverAddr)
}
