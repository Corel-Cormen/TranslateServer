package ServerCore

import (
	"net/http"
	"TranslateServer/internal/ServerPlatform/api"
)

type EchoHandler struct {}

func (h *EchoHandler) Handle(handler ServerCoreApi.HandlerInterface) {
	handler.Callback(http.StatusOK, "echo")
}
