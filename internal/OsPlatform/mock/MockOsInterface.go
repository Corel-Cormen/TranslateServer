package MockOsPlatformApi

import (
	"io"

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

func (m *MockOsInterface) ReadFile(path string) ([]byte, error) {
	args := m.Called(path)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockOsInterface) SetEnv(name string, value string) error {
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

func (m *MockOsInterface) AsyncCommand(name string, args ...string) (OsPlatformApi.ProcessProp, error) {
	callArgs := append([]interface{}{name}, stringSliceToInterfaceSlice(args)...)
	mockArgs := m.Called(callArgs...)

	processProp := OsPlatformApi.ProcessProp{}

	if mockArgs.Get(0) == nil {
		processProp.Pid = mockArgs.Get(0).(int)
	}
	if mockArgs.Get(1) != nil {
		processProp.In = mockArgs.Get(1).(io.WriteCloser)
	}
	if mockArgs.Get(2) != nil {
		processProp.Out = mockArgs.Get(2).(io.ReadCloser)
	}
	if mockArgs.Get(3) != nil {
		processProp.Err = mockArgs.Get(3).(io.ReadCloser)
	}

	return processProp, mockArgs.Error(4)

}

func (m *MockOsInterface) GetProcess(pid int) (OsPlatformApi.ProcessInterface, error) {
	args := m.Called(pid)
	return args.Get(0).(OsPlatformApi.ProcessInterface), args.Error(1)
}
