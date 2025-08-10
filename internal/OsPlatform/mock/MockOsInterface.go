package MockOsPlatformApi

import (
	"TranslateServer/internal/OsPlatform/api"
	"github.com/stretchr/testify/mock"
	"io"
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

func stringSliceToInterfaceSlice(ss []string) []interface{} {
	out := make([]interface{}, len(ss))
	for i, s := range ss {
		out[i] = s
	}
	return out
}

func (m *MockOsInterface) AsyncCommand(name string, args ...string) (io.WriteCloser, io.ReadCloser, io.ReadCloser, error) {
	callArgs := append([]interface{}{name}, stringSliceToInterfaceSlice(args)...)
	mockArgs := m.Called(callArgs...)

	var stdin io.WriteCloser
	var stdout io.ReadCloser
	var stderr io.ReadCloser

	if mockArgs.Get(0) != nil {
		stdin = mockArgs.Get(0).(io.WriteCloser)
	}
	if mockArgs.Get(1) != nil {
		stdout = mockArgs.Get(1).(io.ReadCloser)
	}
	if mockArgs.Get(2) != nil {
		stderr = mockArgs.Get(2).(io.ReadCloser)
	}

	return stdin, stdout, stderr, mockArgs.Error(3)
}
