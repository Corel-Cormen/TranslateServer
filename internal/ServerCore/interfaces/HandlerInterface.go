package Server_Interface

import "net/http"

type HandlerInterface interface {
	http.Handler
}
