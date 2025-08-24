package MockOsPlatformApi

import (
	"bytes"

	"github.com/stretchr/testify/mock"
)

type MockWriteCloser struct {
	mock.Mock
	Buf *bytes.Buffer
}

func (m *MockWriteCloser) Write(p []byte) (int, error) {
	args := m.Called(p)
	n, _ := m.Buf.Write(p)
	return n, args.Error(1)
}

func (m *MockWriteCloser) Close() error {
	args := m.Called()
	return args.Error(0)
}
