package ServerCore

import (
	"net/http"
	"testing"

	"TranslateServer/internal/ServerPlatform/mock"
)

func TestEchoHandler_HandleEchoMessage(t *testing.T) {
	echoHandler := &EchoHandler{}
	mockHandler := new(MockServerInterface.MockServerInterface)

	mockHandler.On("Callback", http.StatusOK, "echo").Return()

	echoHandler.Handle(mockHandler)

	mockHandler.AssertExpectations(t)
}
