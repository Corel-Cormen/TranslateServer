package OsPlatformApi

type OsInterface interface {
	FileExist(path string) bool
	OpenFile(path string) (FileInterface, error)
	ExeScript(path string) ([]byte, error)
	SetEnv(name, value string) error
	LookupEnv(env string) (string, bool)
}
