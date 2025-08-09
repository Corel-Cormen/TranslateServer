package MockOsPlatformApi

import (
	"TranslateServer/internal/OsPlatform/api"
	"github.com/stretchr/testify/mock"
)

type MockOsInterface struct {
	mock.Mock
}

func (m *MockOsInterface) FileExist(path string) bool {
	args := m.Called(path)
	return args.Bool(0)
}

func (m *MockOsInterface) OpenFile(path string) (OsPlatformApi.FileInterface, error) {
	args := m.Called(path)
	return args.Get(0).(OsPlatformApi.FileInterface), args.Error(1)
}

func (m *MockOsInterface) SetEnv(name, value string) error {
	args := m.Called(name, value)
	return args.Error(0)
}

func (m *MockOsInterface) LookupEnv(env string) (string, bool) {
	args := m.Called(env)
	return args.String(0), args.Bool(1)
}
