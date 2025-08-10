package MockServerInterface

import (
	"TranslateServer/internal/ServerPlatform/api"
	"github.com/stretchr/testify/mock"
)

type MockRouterInterface struct {
	mock.Mock
}

func (m *MockRouterInterface) GET(path string, handler func(ServerCoreApi.HandlerInterface)) {
	m.Called(path, handler)
}

func (m *MockRouterInterface) POST(path string, handler func(ServerCoreApi.HandlerInterface)) {
	m.Called(path, handler)
}

func (m *MockRouterInterface) Run(addr ...string) error {
	args := m.Called(addr)
	return args.Error(0)
}
