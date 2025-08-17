package OsPlatformImpl

import (
	"os"
	"os/exec"

	"TranslateServer/internal/OsPlatform/api"
)

type OsFacade struct{}

func (f *OsFacade) FileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil && !os.IsNotExist(err) && !os.IsPermission(err)
}

func (f *OsFacade) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func (f *OsFacade) SetEnv(name, value string) error {
	return os.Setenv(name, value)
}

func (f *OsFacade) LookupEnv(env string) (string, bool) {
	return os.LookupEnv(env)
}

func (f *OsFacade) AsyncCommand(name string, args ...string) (OsPlatformApi.ProcessProp, error) {
	processProp := OsPlatformApi.ProcessProp{}
	c := exec.Command(name, args...)
	stdin, _ := c.StdinPipe()
	stdout, _ := c.StdoutPipe()
	stderr, _ := c.StderrPipe()

	err := c.Start()
	if err == nil {
		processProp.Pid = c.Process.Pid
		processProp.In = stdin
		processProp.Out = stdout
		processProp.Err = stderr

		go func() {
			_ = c.Wait()
		}()
	}

	return processProp, err
}

func (f *OsFacade) GetProcess(pid int) (OsPlatformApi.ProcessInterface, error) {
	proc, err := os.FindProcess(pid)
	if err == nil {
		return NewProcessFacade(proc), nil
	}
	return nil, err
}
