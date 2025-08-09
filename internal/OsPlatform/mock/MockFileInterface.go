package MockOsPlatformApi

import (
	"github.com/stretchr/testify/mock"
)

type MockFileInterface struct {
	mock.Mock
}

func (m *MockFileInterface) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockFileInterface) Read() ([]byte, error) {
	args := m.Called()
	return args.Get(0).([]byte), args.Error(1)
}
