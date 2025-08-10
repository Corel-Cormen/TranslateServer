package MockServerInterface

import (
	"github.com/stretchr/testify/mock"
)

type MockServerInterface struct {
	mock.Mock
}

func (m *MockServerInterface) TextCallback(code int, obj interface{}) {
	m.Called(code, obj)
}

func (m *MockServerInterface) JsonCallback(code int, obj interface{}) {
	m.Called(code, obj)
}

func (m *MockServerInterface) BindJSON(obj interface{}) error {
	args := m.Called(obj)
	return args.Error(0)
}
