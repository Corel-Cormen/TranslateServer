package MockTranslator

import (
	ConfigApi "TranslateServer/internal/Config/api"
	"github.com/stretchr/testify/mock"
)

type MockTranslatorInterface struct {
	mock.Mock
}

func (m *MockTranslatorInterface) Configure(cfg ConfigApi.ConfigData) error {
	args := m.Called(cfg)
	return args.Error(0)
}

func (m *MockTranslatorInterface) Run() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockTranslatorInterface) Translate(language string, text string) (string, error) {
	args := m.Called(language, text)
	return args.String(0), args.Error(1)
}
