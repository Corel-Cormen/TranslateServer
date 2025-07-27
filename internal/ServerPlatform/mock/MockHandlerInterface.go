package MockServerInterface

import (
	"github.com/stretchr/testify/mock"
)

type MockServerInterface struct {
    mock.Mock
}

func (m *MockServerInterface) Callback(code int, obj interface{}) {
    m.Called(code, obj)
}
