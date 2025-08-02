package ConfigCore

import (
	"errors"
	"fmt"
	"TranslateServer/internal/Config/api"
)

type ConfigEnvReader struct {
	envReader ConfigApi.EnvReaderInterface
}

func NewConfigEnvReader(envReader ConfigApi.EnvReaderInterface) ConfigApi.ConfigReaderInterface {
	return &ConfigEnvReader{
		envReader: envReader,
	}
}

func (c *ConfigEnvReader) Read() (ConfigApi.ConfigData, error) {

	if err := c.envReader.LoadFileEnv(); err != nil {
		return ConfigApi.ConfigData{}, fmt.Errorf("failed to load environment variables: %w", err)
	}

	marianPath, err := c.envReader.Read("MARIAN_INSTALL_PATH")
	if err != nil {
		return ConfigApi.ConfigData{}, fmt.Errorf("failed to read MARIAN_INSTALL_PATH: %w", err)
	}
	fmt.Println("Marian install path:", marianPath)

	vocabPath, err := c.envReader.Read("VOCAB_PATH")
	if err != nil {
		return ConfigApi.ConfigData{}, fmt.Errorf("failed to read VOCAB_PATH: %w", err)
	}
	fmt.Println("Vocab path:", vocabPath)

	if marianPath == "" || vocabPath == "" {
		return ConfigApi.ConfigData{}, errors.New("environment variables not set")
	}

	return ConfigApi.ConfigData{
		MarianInstallPath: marianPath,
		VocabPath:         vocabPath,
	}, nil
}
