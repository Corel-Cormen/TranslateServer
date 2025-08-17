package MockOsPlatformApi

import (
	"github.com/stretchr/testify/mock"
)

type MockProcessInterface struct {
	mock.Mock
}

func (m *MockProcessInterface) Signal(signal int) error {
	args := m.Called(signal)
	return args.Error(0)
}
