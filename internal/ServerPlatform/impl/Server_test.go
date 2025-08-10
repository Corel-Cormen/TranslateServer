package ServerCore

import (
	"testing"

	"TranslateServer/internal/ServerPlatform/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServer_CorrectStart(t *testing.T) {
	expectedAddr := "127.0.0.1:8080"

	mockRouter := new(MockServerInterface.MockRouterInterface)
	mockRouter.On("GET", "/echo", mock.AnythingOfType("func(ServerCoreApi.HandlerInterface)")).Return()
	mockRouter.On("Run", []string{expectedAddr}).Return(nil)

	server := NewServer("127.0.0.1", 8080, mockRouter)

	err := server.Start()

	assert.NoError(t, err)
	mockRouter.AssertExpectations(t)
}
