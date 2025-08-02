package OsPlatformImpl

import (
	"os"
	"os/exec"
	"TranslateServer/internal/OsPlatform/api"
)

type OsFasade struct{}

func (f *OsFasade) FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil && !os.IsNotExist(err) && !os.IsPermission(err)
}

func (f *OsFasade) OpenFile(path string) (OsPlatformApi.FileInterface, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return NewFileFasade(file), nil
}

func (f *OsFasade) ExeScript(path string) ([]byte, error) {
	cmd := exec.Command("bash", path)
	return cmd.Output()
}

func (f *OsFasade) SetEnv(name, value string) error {
	return os.Setenv(name, value)
}

func (f *OsFasade) LookupEnv(env string) (string, bool) {
	return os.LookupEnv(env)
}
