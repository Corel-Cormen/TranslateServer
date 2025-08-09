package OsPlatformImpl

import (
	"os"

	"TranslateServer/internal/OsPlatform/api"
)

type OsFacade struct{}

func (f *OsFacade) FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil && !os.IsNotExist(err) && !os.IsPermission(err)
}

func (f *OsFacade) OpenFile(path string) (OsPlatformApi.FileInterface, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return NewFileFacade(file), nil
}

func (f *OsFacade) SetEnv(name, value string) error {
	return os.Setenv(name, value)
}

func (f *OsFacade) LookupEnv(env string) (string, bool) {
	return os.LookupEnv(env)
}
