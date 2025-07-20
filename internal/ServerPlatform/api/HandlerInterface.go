package ServerCoreApi

type HandlerInterface interface {
	Callback(code int, obj interface{})
}
