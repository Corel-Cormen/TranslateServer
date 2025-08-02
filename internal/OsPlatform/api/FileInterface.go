package OsPlatformApi

type FileInterface interface {
	Close() error
	Read() ([]byte, error)
}
