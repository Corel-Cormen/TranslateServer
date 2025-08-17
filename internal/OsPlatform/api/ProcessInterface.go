package OsPlatformApi

type ProcessInterface interface {
	Signal(signal int) error
}
