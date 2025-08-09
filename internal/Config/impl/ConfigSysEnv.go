package ConfigCore

import (
	"errors"

	"TranslateServer/internal/Config/api"
)

type ConfigSysEnv struct {
	configReader ConfigApi.ConfigReaderInterface
	configData   ConfigApi.ConfigData
	isInit       bool
}

func NewConfigSysEnv(cfgReader ConfigApi.ConfigReaderInterface) ConfigApi.ConfigInterface {
	return &ConfigSysEnv{
		configReader: cfgReader,
		configData:   ConfigApi.ConfigData{},
		isInit:       false,
	}
}

func (c *ConfigSysEnv) Init() error {
	configData, err := c.configReader.Read()
	if err == nil {
		c.configData = configData
		c.isInit = true
	}
	return err
}

func (c *ConfigSysEnv) Get() (ConfigApi.ConfigData, error) {
	if !c.isInit {
		return ConfigApi.ConfigData{}, errors.New("ConfigSysEnv not initialized")
	}
	return c.configData, nil
}
