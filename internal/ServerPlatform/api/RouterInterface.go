package ServerCoreApi

type RoutherInterface interface {
	GET(path string, callback func(handler HandlerInterface))
	Run(addr ...string) error
}
