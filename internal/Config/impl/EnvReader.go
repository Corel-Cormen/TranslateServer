package ConfigCore

import (
	"bytes"
	"fmt"
	"strings"

	"TranslateServer/internal/Config/api"
	"TranslateServer/internal/OsPlatform/api"
)

type EnvReader struct {
	envFilePath string
	osInterface OsPlatformApi.OsInterface
}

func NewEnvReader(envFilePath string, osInterface OsPlatformApi.OsInterface) ConfigApi.EnvReaderInterface {
	return &EnvReader{
		envFilePath: envFilePath,
		osInterface: osInterface,
	}
}

func (e *EnvReader) loadEnv() error {
	errStatus := error(nil)

	fileContent, err := e.osInterface.ReadFile(e.envFilePath)
	if err != nil {
		errStatus = fmt.Errorf("failed to read env file: %w", err)
	}

	if errStatus == nil {
		lineCount := bytes.Count(fileContent, []byte{'\n'})
		envParts := strings.SplitN(string(fileContent), "\n", lineCount)

		for _, line := range envParts {
			line = strings.TrimSpace(line)

			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				if err := e.osInterface.SetEnv(name, value); err != nil {
					errStatus = fmt.Errorf("failed to set env variable %s: %w", name, err)
				}
			} else {
				errStatus = fmt.Errorf("invalid env line: %s", line)
			}

			if errStatus != nil {
				break
			}
		}
	}

	return errStatus
}

func (e *EnvReader) LoadFileEnv() error {
	errStatus := error(nil)

	if e.osInterface.FileExist(e.envFilePath) {
		fmt.Println("Config script detected | Scrap data ENV from:", e.envFilePath)
		if err := e.loadEnv(); err != nil {
			errStatus = fmt.Errorf("failed to load environment variables: %w", err)
		}
	}

	return errStatus
}

func (e *EnvReader) Read(env string) (string, error) {
	value, ok := e.osInterface.LookupEnv(env)
	if !ok {
		return "", fmt.Errorf("environment variable %s not found", env)
	}

	if value[0] == '"' && value[len(value)-1] == '"' {
		value = value[1 : len(value)-1]
	}

	return value, nil
}
