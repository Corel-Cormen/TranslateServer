package ConfigApi

type ConfigData struct {
	MarianInstallPath string
	VocabPath         string
}

type ConfigInterface interface {
	Init() error
	Get() (ConfigData, error)
}
