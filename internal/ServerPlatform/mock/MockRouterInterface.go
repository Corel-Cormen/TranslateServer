package MockServerInterface

import (
	"github.com/stretchr/testify/mock"
	"TranslateServer/internal/ServerPlatform/api"
)

type MockRouterInterface struct {
    mock.Mock
}

func (m *MockRouterInterface) GET(path string, handler func(ServerCoreApi.HandlerInterface)) {
    m.Called(path, handler)
}

func (m *MockRouterInterface) Run(addr ...string) error {
	args := m.Called(addr)
	return args.Error(0)
}
