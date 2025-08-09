package ConfigApi

type EnvReaderInterface interface {
	LoadFileEnv() error
	Read(env string) (string, error)
}
