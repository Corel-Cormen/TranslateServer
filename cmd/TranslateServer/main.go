package main

import (
	"TranslateServer/internal/ServerCore"
)

func main() {
	server := Server_Core.NewServer()
	if err := server.Start(); err != nil {
		panic(err)
	}
}
