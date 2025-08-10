package ServerCoreApi

type RouterInterface interface {
	GET(path string, callback func(handler HandlerInterface))
	POST(path string, handler func(c HandlerInterface))
	Run(addr ...string) error
}
