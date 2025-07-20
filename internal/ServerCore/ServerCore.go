package Server_Core

import (
	"net/http"
	"TranslateServer/internal/ServerCore/interfaces"
	"TranslateServer/internal/ServerCore/source"
)

func NewServer() Server_Interface.ServerInterface {
	mux := http.NewServeMux()
	mux.Handle("/", &Server_Source.EchoHandler{})

	return &Server_Source.HTTPServer{
		Address: "127.0.0.1",
		Port:    5000,
		Handler: mux,
	}
}
