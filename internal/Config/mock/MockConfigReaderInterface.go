package MockConfigApi

import (
	"TranslateServer/internal/Config/api"
	"github.com/stretchr/testify/mock"
)

type MockConfigReaderInterface struct {
	mock.Mock
}

func (m *MockConfigReaderInterface) Read() (ConfigApi.ConfigData, error) {
	args := m.Called()
	return args.Get(0).(ConfigApi.ConfigData), args.Error(1)
}
