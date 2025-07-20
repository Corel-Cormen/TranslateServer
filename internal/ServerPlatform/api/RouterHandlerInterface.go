package ServerCoreApi

type RouterHandlerInterface interface {
	Handle(handler HandlerInterface)
}
