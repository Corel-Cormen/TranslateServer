package OsPlatformApi

import "io"

type OsInterface interface {
	FileExist(path string) bool
	OpenFile(path string) (FileInterface, error)
	SetEnv(name, value string) error
	LookupEnv(env string) (string, bool)
	AsyncCommand(name string, args ...string) (io.WriteCloser, io.ReadCloser, io.ReadCloser, error)
}
