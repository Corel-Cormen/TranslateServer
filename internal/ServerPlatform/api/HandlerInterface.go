package ServerCoreApi

type HandlerInterface interface {
	TextCallback(code int, obj interface{})
	JsonCallback(code int, obj interface{})
	BindJSON(obj interface{}) error
}
