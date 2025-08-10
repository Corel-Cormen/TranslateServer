package main

import (
	"TranslateServer/internal/Config"
	"TranslateServer/internal/ServerPlatform"
	"TranslateServer/internal/Translator"
)

func main() {
	Config.InitializeConfig()
	Translator.RunTranslator()
	ServerPlatform.Run()
}
