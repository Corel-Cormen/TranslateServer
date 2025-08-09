package MockConfigApi

import (
	"github.com/stretchr/testify/mock"
)

type MockEnvReaderInterface struct {
	mock.Mock
}

func (m *MockEnvReaderInterface) LoadFileEnv() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockEnvReaderInterface) Read(env string) (string, error) {
	args := m.Called(env)
	return args.String(0), args.Error(1)
}
