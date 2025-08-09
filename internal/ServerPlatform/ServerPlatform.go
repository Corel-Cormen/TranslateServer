package ServerPlatform

import (
	"TranslateServer/internal/ServerPlatform/instance"
)

func Run() {
	server := ServerPlatformInstance.GetServer()
	if err := server.Start(); err != nil {
		panic(err)
	}
}
