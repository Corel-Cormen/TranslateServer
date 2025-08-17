package OsPlatformApi

import (
	"io"
)

type ProcessProp struct {
	Pid int
	In  io.WriteCloser
	Out io.ReadCloser
	Err io.ReadCloser
}

type OsInterface interface {
	FileExist(path string) bool
	ReadFile(path string) ([]byte, error)
	SetEnv(name string, value string) error
	LookupEnv(env string) (string, bool)
	AsyncCommand(name string, args ...string) (ProcessProp, error)
	GetProcess(pid int) (ProcessInterface, error)
}
