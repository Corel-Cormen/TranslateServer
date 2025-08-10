package Translator

import (
	"TranslateServer/internal/Config/instance"
	"TranslateServer/internal/Translator/instance"
)

func RunTranslator() {
	translatorInstance := TranslatorInstance.GetTranslatorInstance()

	cfg, err := ConfigInstance.GetConfigInstance().Get()
	if err != nil {
		panic("Failed to get config: " + err.Error())
	}

	err = translatorInstance.Configure(cfg)
	if err != nil {
		panic("Failed Translator configure: " + err.Error())
	}

	err = translatorInstance.Run()
	if err != nil {
		panic("Failed run Translate instance: " + err.Error())
	}
}
