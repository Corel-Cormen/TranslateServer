package ServerCore

import (
	"fmt"

	"TranslateServer/internal/ServerPlatform/api"
	"TranslateServer/internal/Translator/api"
)

type Server struct {
	address          string
	port             int
	router           ServerCoreApi.RouterInterface
	translateService TranslatorApi.TranslatorInterface
}

func NewServer(address string, port int, router ServerCoreApi.RouterInterface, translateService TranslatorApi.TranslatorInterface) ServerCoreApi.ServerInterface {
	return &Server{
		address:          address,
		port:             port,
		router:           router,
		translateService: translateService,
	}
}

func wrapHandler(handler ServerCoreApi.RouterHandlerInterface) func(ServerCoreApi.HandlerInterface) {
	return func(h ServerCoreApi.HandlerInterface) {
		handler.Handle(h)
	}
}

func (s *Server) Start() error {
	s.router.GET("/echo", wrapHandler(&EchoHandler{}))
	s.router.POST("/translate", wrapHandler(&TranslateHandler{Translator: s.translateService}))

	serverAddr := fmt.Sprintf("%s:%d", s.address, s.port)
	return s.router.Run(serverAddr)
}
