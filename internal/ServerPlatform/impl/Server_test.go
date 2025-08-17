package ServerCore

import (
	"testing"

	"TranslateServer/internal/ServerPlatform/mock"
	"TranslateServer/internal/Supervisor/mock"
	"TranslateServer/internal/Translator/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestServer_CorrectStart(t *testing.T) {
	expectedAddr := "127.0.0.1:8080"

	mockTranslator := new(MockTranslator.MockTranslatorInterface)
	mockRouter := new(MockServerInterface.MockRouterInterface)
	supervisorMock := new(MockSupervisorApi.MockSupervisorApi)

	mockRouter.On("GET", "/echo", mock.AnythingOfType("func(ServerCoreApi.HandlerInterface)")).Return()
	mockRouter.On("GET", "/metric", mock.AnythingOfType("func(ServerCoreApi.HandlerInterface)")).Return()
	mockRouter.On("POST", "/translate", mock.AnythingOfType("func(ServerCoreApi.HandlerInterface)")).Return()
	mockRouter.On("Run", []string{expectedAddr}).Return(nil)

	server := NewServer("127.0.0.1", 8080, mockRouter, mockTranslator, supervisorMock)

	err := server.Start()

	assert.NoError(t, err)
	mockRouter.AssertExpectations(t)
}
