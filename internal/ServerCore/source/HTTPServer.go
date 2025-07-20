package Server_Source

import (
	"fmt"
	"net/http"
	"TranslateServer/internal/ServerCore/interfaces"
)

type HTTPServer struct {
	Address string
	Port    int
	Handler Server_Interface.HandlerInterface
}

func (s *HTTPServer) Start() error {
	address := fmt.Sprintf("%s:%d", s.Address, s.Port)
	fmt.Println("Starting server at:", address)
	return http.ListenAndServe(address, s.Handler)
}
