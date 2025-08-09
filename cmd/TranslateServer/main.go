package main

import (
	"TranslateServer/internal/ServerPlatform"
	"TranslateServer/internal/Config"
)

func main() {
	Config.InitializeConfig()
	ServerPlatform.Run()
}
