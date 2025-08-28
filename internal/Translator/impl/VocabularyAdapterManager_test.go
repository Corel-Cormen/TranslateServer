package TranslatorImpl

import (
	"bytes"
	"errors"
	"github.com/stretchr/testify/mock"
	"io"
	"testing"

	"TranslateServer/internal/OsPlatform/mock"
	"TranslateServer/internal/SentenceFormatter/mock"
	"TranslateServer/internal/Supervisor/api"
	"TranslateServer/internal/Supervisor/mock"
	"TranslateServer/internal/Translator/api"
	"TranslateServer/internal/Translator/mock"
	"github.com/stretchr/testify/require"
)

func TestVocabularyAdapterManager_Subscribe_Success(t *testing.T) {
	mockSupervisor := new(MockSupervisorApi.MockSupervisor)
	mockSentenceFormatter := new(MockSentenceFormatter.MockSentenceFormatterInterface)
	manager := NewVocabularyAdapterManager(mockSupervisor, mockSentenceFormatter)

	mockVocab1 := new(MockTranslator.MockVocabularyInterface)
	mockVocab1.On("GetId").Return("v1")

	err := manager.Subscribe(mockVocab1)
	require.NoError(t, err)

	mockVocab2 := new(MockTranslator.MockVocabularyInterface)
	mockVocab2.On("GetId").Return("v2")

	err = manager.Subscribe(mockVocab2)
	require.NoError(t, err)

	require.Equal(t, len(manager.(*VocabularyAdapterManager).vocabularyList), 2)

	mockVocab1.AssertExpectations(t)
	mockVocab2.AssertExpectations(t)
}

func TestVocabularyAdapterManager_Subscribe_VocabIsSubscribed(t *testing.T) {
	mockSupervisor := new(MockSupervisorApi.MockSupervisor)
	mockSentenceFormatter := new(MockSentenceFormatter.MockSentenceFormatterInterface)
	manager := NewVocabularyAdapterManager(mockSupervisor, mockSentenceFormatter)

	mockVocab1 := new(MockTranslator.MockVocabularyInterface)
	mockVocab1.On("GetId").Return("v1")

	err := manager.Subscribe(mockVocab1)
	require.NoError(t, err)

	mockVocab2 := new(MockTranslator.MockVocabularyInterface)
	mockVocab2.On("GetId").Return("v1")

	err = manager.Subscribe(mockVocab2)
	require.Error(t, err)

	require.Equal(t, len(manager.(*VocabularyAdapterManager).vocabularyList), 1)

	mockVocab1.AssertExpectations(t)
	mockVocab2.AssertExpectations(t)
}

func TestVocabularyAdapterManager_Init_Success(t *testing.T) {
	mockSupervisor := new(MockSupervisorApi.MockSupervisor)
	mockSentenceFormatter := new(MockSentenceFormatter.MockSentenceFormatterInterface)
	manager := NewVocabularyAdapterManager(mockSupervisor, mockSentenceFormatter)

	mockVocab := new(MockTranslator.MockVocabularyInterface)
	mockVocab.On("GetId").Return("v1")
	props := TranslatorApi.VocabularyProperties{Decoder: "dec", Model: "mod", Vocab: "voc"}
	mockVocab.On("GetProperties").Return(props)

	in := &MockOsPlatformApi.NopWriteCloser{}
	out := io.NopCloser(bytes.NewBufferString("out"))
	mockVocab.On("RegisterInput", in).Return(nil)
	mockVocab.On("RegisterOutput", out).Return(nil)

	mockSupervisor.On("InitVocabTaskChannel", "v1", "dec", "mod", "voc").Return(
		SupervisorApi.Channel{In: in, Out: out}, nil)

	_ = manager.Subscribe(mockVocab)
	err := manager.Init()
	require.NoError(t, err)

	mockVocab.AssertExpectations(t)
	mockSupervisor.AssertExpectations(t)
}

func TestVocabularyAdapterManager_Init_SupervisorError(t *testing.T) {
	mockSupervisor := new(MockSupervisorApi.MockSupervisor)
	mockSentenceFormatter := new(MockSentenceFormatter.MockSentenceFormatterInterface)
	manager := NewVocabularyAdapterManager(mockSupervisor, mockSentenceFormatter)

	mockVocab := new(MockTranslator.MockVocabularyInterface)
	mockVocab.On("GetId").Return("v1")
	props := TranslatorApi.VocabularyProperties{Decoder: "dec", Model: "mod", Vocab: "voc"}
	mockVocab.On("GetProperties").Return(props)

	mockSupervisor.On("InitVocabTaskChannel", "v1", "dec", "mod", "voc").Return(
		SupervisorApi.Channel{}, errors.New("init error"))

	_ = manager.Subscribe(mockVocab)
	err := manager.Init()
	require.Error(t, err)
	require.EqualError(t, err, "init error")

	mockVocab.AssertExpectations(t)
	mockSupervisor.AssertExpectations(t)
}

func TestVocabularyAdapterManager_Init_RegisterInputError(t *testing.T) {
	mockSupervisor := new(MockSupervisorApi.MockSupervisor)
	mockSentenceFormatter := new(MockSentenceFormatter.MockSentenceFormatterInterface)
	manager := NewVocabularyAdapterManager(mockSupervisor, mockSentenceFormatter)

	mockVocab := new(MockTranslator.MockVocabularyInterface)
	mockVocab.On("GetId").Return("v1")
	props := TranslatorApi.VocabularyProperties{Decoder: "dec", Model: "mod", Vocab: "voc"}
	mockVocab.On("GetProperties").Return(props)

	in := &MockOsPlatformApi.NopWriteCloser{}
	out := io.NopCloser(bytes.NewBufferString("out"))
	mockVocab.On("RegisterInput", in).Return(errors.New("register input error"))

	mockSupervisor.On("InitVocabTaskChannel", "v1", "dec", "mod", "voc").Return(
		SupervisorApi.Channel{In: in, Out: out}, nil)

	_ = manager.Subscribe(mockVocab)
	err := manager.Init()
	require.Error(t, err)
	require.EqualError(t, err, "register input error")

	mockVocab.AssertExpectations(t)
	mockSupervisor.AssertExpectations(t)
}

func TestVocabularyAdapterManager_Init_RegisterOutputError(t *testing.T) {
	mockSupervisor := new(MockSupervisorApi.MockSupervisor)
	mockSentenceFormatter := new(MockSentenceFormatter.MockSentenceFormatterInterface)
	manager := NewVocabularyAdapterManager(mockSupervisor, mockSentenceFormatter)

	mockVocab := new(MockTranslator.MockVocabularyInterface)
	mockVocab.On("GetId").Return("v1")
	props := TranslatorApi.VocabularyProperties{Decoder: "dec", Model: "mod", Vocab: "voc"}
	mockVocab.On("GetProperties").Return(props)

	in := &MockOsPlatformApi.NopWriteCloser{}
	out := io.NopCloser(bytes.NewBufferString("out"))
	mockVocab.On("RegisterInput", in).Return(nil)
	mockVocab.On("RegisterOutput", out).Return(errors.New("register output error"))

	mockSupervisor.On("InitVocabTaskChannel", "v1", "dec", "mod", "voc").Return(
		SupervisorApi.Channel{In: in, Out: out}, nil)

	_ = manager.Subscribe(mockVocab)
	err := manager.Init()
	require.Error(t, err)
	require.EqualError(t, err, "register output error")

	mockVocab.AssertExpectations(t)
	mockSupervisor.AssertExpectations(t)
}

func TestVocabularyAdapterManager_Deinit(t *testing.T) {
	mockSupervisor := new(MockSupervisorApi.MockSupervisor)

	mockVocab1 := new(MockTranslator.MockVocabularyInterface)
	mockVocab1.On("Unregister").Return(nil)
	mockVocab2 := new(MockTranslator.MockVocabularyInterface)
	mockVocab2.On("Unregister").Return(errors.New("fail"))

	manager := VocabularyAdapterManager{
		supervisorInterface: mockSupervisor,
		vocabularyList: []VocabularyObject{
			{vocabulary: mockVocab1, isInit: true},
			{vocabulary: mockVocab2, isInit: true},
		},
	}

	err := manager.Deinit()
	require.Error(t, err)
	require.EqualError(t, err, "fail")

	require.False(t, manager.vocabularyList[0].isInit)
	require.True(t, manager.vocabularyList[1].isInit)

	mockVocab1.AssertExpectations(t)
	mockVocab2.AssertExpectations(t)
}

func TestVocabularyAdapterManager_Translate(t *testing.T) {
	mockSupervisor := new(MockSupervisorApi.MockSupervisor)
	mockSentenceFormatter := new(MockSentenceFormatter.MockSentenceFormatterInterface)
	manager := NewVocabularyAdapterManager(mockSupervisor, mockSentenceFormatter)

	mockVocab := new(MockTranslator.MockVocabularyInterface)
	mockVocab.On("GetId").Return("v1")

	_ = manager.Subscribe(mockVocab)

	result, err := manager.Translate("v1", "hello")
	require.Error(t, err)
	require.EqualError(t, err, "translator is not init")
	require.Equal(t, "", result)

	manager.(*VocabularyAdapterManager).vocabularyList[0].isInit = true
	mockSentenceFormatter.On("PrepareInput", mock.Anything).Return("hello")
	mockVocab.On("Translate", "hello").Return("translated", nil)
	mockSentenceFormatter.On("CleanOutput", mock.Anything).Return("translated")

	result, err = manager.Translate("v1", "hello")
	require.NoError(t, err)
	require.Equal(t, "translated", result)

	result, err = manager.Translate("v2", "hello")
	require.Error(t, err)
	require.EqualError(t, err, "not found translator")
	require.Equal(t, "", result)

	mockVocab.AssertExpectations(t)
	mockSentenceFormatter.AssertExpectations(t)
}
