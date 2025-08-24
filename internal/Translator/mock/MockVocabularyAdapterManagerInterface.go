package MockTranslator

import (
	"TranslateServer/internal/Translator/api"
	"github.com/stretchr/testify/mock"
)

type MockVocabularyAdapterManagerInterface struct {
	mock.Mock
}

func (m *MockVocabularyAdapterManagerInterface) Subscribe(vocabularyInterface TranslatorApi.VocabularyInterface) error {
	return m.Called(vocabularyInterface).Error(0)
}

func (m *MockVocabularyAdapterManagerInterface) Init() error {
	return m.Called().Error(0)
}

func (m *MockVocabularyAdapterManagerInterface) Deinit() error {
	return m.Called().Error(0)
}

func (m *MockVocabularyAdapterManagerInterface) Translate(id string, text string) (string, error) {
	args := m.Called(id, text)
	return args.String(0), args.Error(1)
}
