package ConfigCore

import (
	"fmt"
	"strings"
	"bytes"
	"TranslateServer/internal/Config/api"
	"TranslateServer/internal/OsPlatform/api"
)

const envFilePath = ".env"

type EnvReader struct {
	scriptPath string
	osInterface OsPlatformApi.OsInterface
}

func NewEnvReader(scriptPath string, osInterface OsPlatformApi.OsInterface) ConfigApi.EnvReaderInterface {
	return &EnvReader{
		scriptPath: scriptPath,
		osInterface: osInterface,
	}
}

func (e *EnvReader) writeEnvFromScript() error {
	_, err := e.osInterface.ExeScript(e.scriptPath)
	if err != nil {
		return fmt.Errorf("failed to execute script: %w", err)
	}
	return nil
}

func (e *EnvReader) loadEnv(path string) error {

	errStatus := error(nil)
	file, err := e.osInterface.OpenFile(path)
	if err != nil {
		errStatus = fmt.Errorf("failed to open env file: %w", err)
	}
	defer file.Close()

	if errStatus == nil {
		fileContent, err := file.Read()
		if err != nil {
			errStatus = fmt.Errorf("failed to read env file: %w", err)
		}

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

	if e.osInterface.FileExist(e.scriptPath) {
		fmt.Println("Config script detected | Scrap data ENV from:", e.scriptPath)
		if err := e.writeEnvFromScript(); err != nil {
			errStatus = fmt.Errorf("failed to write environment variables from script: %w", err)
		}
		if errStatus == nil {
			if err := e.loadEnv(envFilePath); err != nil {
				errStatus = fmt.Errorf("failed to load environment variables: %w", err)
			}
		}
	}

	return errStatus
}

func (e *EnvReader) Read(env string) (string, error) {
	value, ok := e.osInterface.LookupEnv(env)
	if !ok {
		return "", fmt.Errorf("environment variable %s not found", env)
	}
	return value, nil
}
