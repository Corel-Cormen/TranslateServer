package ServerCore

import (
	"fmt"

	"TranslateServer/internal/ServerPlatform/api"
	"TranslateServer/internal/Supervisor/api"
	"TranslateServer/internal/Translator/api"
)

type Server struct {
	address             string
	port                int
	router              ServerCoreApi.RouterInterface
	translateService    TranslatorApi.TranslatorInterface
	supervisorInterface SupervisorApi.SupervisorInterface
}

func NewServer(address string, port int, router ServerCoreApi.RouterInterface,
	translateService TranslatorApi.TranslatorInterface, supervisorInterface SupervisorApi.SupervisorInterface) ServerCoreApi.ServerInterface {
	return &Server{
		address:             address,
		port:                port,
		router:              router,
		translateService:    translateService,
		supervisorInterface: supervisorInterface,
	}
}

func wrapHandler(handler ServerCoreApi.RouterHandlerInterface) func(ServerCoreApi.HandlerInterface) {
	return func(h ServerCoreApi.HandlerInterface) {
		handler.Handle(h)
	}
}

func (s *Server) Start() error {
	s.router.GET("/echo", wrapHandler(&EchoHandler{}))
	s.router.GET("/metric", wrapHandler(&MetricHandler{supervisorInterface: s.supervisorInterface}))
	s.router.POST("/translate", wrapHandler(&TranslateHandler{Translator: s.translateService}))

	serverAddr := fmt.Sprintf("%s:%d", s.address, s.port)
	return s.router.Run(serverAddr)
}
