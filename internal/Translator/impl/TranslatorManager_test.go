package TranslatorImpl

import (
	"fmt"
	"testing"

	"TranslateServer/internal/Config/api"
	"TranslateServer/internal/Translator/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestTranslatorManager_Configure_Success(t *testing.T) {
	mockMgr := new(MockTranslator.MockVocabularyAdapterManagerInterface)
	manager := NewTranslatorManager(mockMgr).(*TranslatorManager)

	cfg := ConfigApi.ConfigData{
		MarianInstallPath: "/usr/bin",
		VocabBtPath:       "/data/bt",
		VocabPath:         "/data/main",
	}

	mockMgr.On("Subscribe", mock.Anything).Return(nil).Twice()

	err := manager.Configure(cfg)
	require.NoError(t, err)

	require.Equal(t, "en-pl-BT", manager.en2plBT.Id)
	require.Contains(t, manager.en2plBT.Model, "/data/bt/")
	require.Contains(t, manager.en2plBT.Vocab, "/data/bt/")

	require.Equal(t, "en-pl", manager.en2pl.Id)
	require.Contains(t, manager.en2pl.Model, "/data/main/")
	require.Contains(t, manager.en2pl.Vocab, "/data/main/")

	mockMgr.AssertExpectations(t)
}

func TestTranslatorManager_Configure_SubscribeError(t *testing.T) {
	mockMgr := new(MockTranslator.MockVocabularyAdapterManagerInterface)
	manager := NewTranslatorManager(mockMgr).(*TranslatorManager)

	cfg := ConfigApi.ConfigData{
		MarianInstallPath: "/usr/bin",
		VocabBtPath:       "/data/bt",
		VocabPath:         "/data/main",
	}

	mockMgr.On("Subscribe", mock.Anything).Return(fmt.Errorf("subscribe error")).Once()

	err := manager.Configure(cfg)
	require.Error(t, err)
	require.EqualError(t, err, "subscribe error")

	mockMgr.AssertExpectations(t)
}

func TestTranslatorManager_Run(t *testing.T) {
	mockMgr := new(MockTranslator.MockVocabularyAdapterManagerInterface)
	manager := NewTranslatorManager(mockMgr).(*TranslatorManager)

	mockMgr.On("Init").Return(nil).Once()

	err := manager.Run()
	require.NoError(t, err)

	mockMgr.AssertExpectations(t)
}

func TestTranslatorManager_Translate(t *testing.T) {
	mockMgr := new(MockTranslator.MockVocabularyAdapterManagerInterface)
	manager := NewTranslatorManager(mockMgr).(*TranslatorManager)

	mockMgr.On("Translate", "en-pl", "hello").Return("cześć", nil).Once()

	result, err := manager.Translate("en-pl", "hello")
	require.NoError(t, err)
	require.Equal(t, "cześć", result)

	mockMgr.AssertExpectations(t)
}
