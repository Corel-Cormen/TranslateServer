package Config

import (
	"TranslateServer/internal/Config/instance"
)

func InitializeConfig() {
	configInstance := ConfigInstance.GetConfigInstance()
	if err := configInstance.Init(); err != nil {
		panic("failed to initialize config: " + err.Error())
	}
}
