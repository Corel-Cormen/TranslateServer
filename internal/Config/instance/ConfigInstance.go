package ConfigInstance

import (
	"TranslateServer/internal/Config/api"
	"TranslateServer/internal/Config/impl"
	"TranslateServer/internal/OsPlatform"
	"sync"
)

var (
	configInstance           ConfigApi.ConfigInterface
	configReaderInstance     ConfigApi.ConfigReaderInterface
	envReaderInstance        ConfigApi.EnvReaderInterface
	onceConfigInstance       sync.Once
	onceConfigReaderInstance sync.Once
	onceEnvReaderInstance    sync.Once
)

func getEnvReader() ConfigApi.EnvReaderInterface {
	onceEnvReaderInstance.Do(func() {
		envFilePath := ".env"
		envReaderInstance = ConfigCore.NewEnvReader(envFilePath, OsPlatform.GetOsInstance())
	})
	return envReaderInstance
}

func getConfigReader() ConfigApi.ConfigReaderInterface {
	onceConfigReaderInstance.Do(func() {
		configReaderInstance = ConfigCore.NewConfigEnvReader(getEnvReader())
	})
	return configReaderInstance
}

func GetConfigInstance() ConfigApi.ConfigInterface {
	onceConfigInstance.Do(func() {
		configReader := getConfigReader()
		configInstance = ConfigCore.NewConfigSysEnv(configReader)
	})
	return configInstance
}
