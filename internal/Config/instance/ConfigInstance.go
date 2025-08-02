package ConfigInstance

import (
	"sync"
	"TranslateServer/internal/Config/api"
	"TranslateServer/internal/Config/impl"
	"TranslateServer/internal/OsPlatform/instance"
)

var (
	configInstance ConfigApi.ConfigInterface
	configReaderInstance ConfigApi.ConfigReaderInterface
	envReaderInstance ConfigApi.EnvReaderInterface
	onceConfigInstance sync.Once
	onceConfigReaderInstance sync.Once
	onceEnvReaderInstance sync.Once
)

func getEnvReader() ConfigApi.EnvReaderInterface {
	onceEnvReaderInstance.Do(func() {
		scriptPath := "3rdParty/scripts/scrap_system_variables.sh"
		envReaderInstance = ConfigCore.NewEnvReader(scriptPath, OsInstance.GetOsInstance())
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
