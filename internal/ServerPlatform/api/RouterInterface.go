package ServerCoreApi

type RouterInterface interface {
	GET(path string, callback func(handler HandlerInterface))
	Run(addr ...string) error
}
