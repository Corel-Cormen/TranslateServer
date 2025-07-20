package ServerPlatform

import (
	"TranslateServer/internal/ServerPlatform/impl"
)

func Run() {
	router := ServerCore.NewGinRouter()
	server := ServerCore.NewServer("127.0.0.1", 5000, router)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
