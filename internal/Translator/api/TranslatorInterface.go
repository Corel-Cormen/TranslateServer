package TranslatorApi

import (
	"TranslateServer/internal/Config/api"
)

type TranslatorInterface interface {
	Configure(cfg ConfigApi.ConfigData) error
	Run() error
	Translate(language string, text string) (string, error)
}
