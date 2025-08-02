package ConfigApi

type ConfigReaderInterface interface {
	Read() (ConfigData, error)
}
