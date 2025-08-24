package MockTranslator

import (
	"io"

	"TranslateServer/internal/Translator/api"
	"github.com/stretchr/testify/mock"
)

type MockVocabularyInterface struct {
	mock.Mock
}

func (m *MockVocabularyInterface) GetId() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockVocabularyInterface) GetProperties() TranslatorApi.VocabularyProperties {
	args := m.Called()
	return args.Get(0).(TranslatorApi.VocabularyProperties)
}

func (m *MockVocabularyInterface) Translate(text string) (string, error) {
	args := m.Called(text)
	return args.String(0), args.Error(1)
}

func (m *MockVocabularyInterface) RegisterInput(closer io.WriteCloser) error {
	args := m.Called(closer)
	return args.Error(0)
}

func (m *MockVocabularyInterface) RegisterOutput(closer io.ReadCloser) error {
	args := m.Called(closer)
	return args.Error(0)
}

func (m *MockVocabularyInterface) Unregister() error {
	args := m.Called()
	return args.Error(0)
}
