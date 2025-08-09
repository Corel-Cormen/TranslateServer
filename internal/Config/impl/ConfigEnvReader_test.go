package ConfigCore

import (
	"errors"
	"testing"

	"TranslateServer/internal/Config/mock"
	"github.com/stretchr/testify/assert"
)

func TestConfigEnvReader_Success(t *testing.T) {
	mockEnv := new(MockConfigApi.MockEnvReaderInterface)

	marianPath := "/path/to/marian"
	vocabBtPath := "/path/to/vocab_bt"
	vocabPath := "/path/to/vocab"

	mockEnv.On("LoadFileEnv").Return(nil)
	mockEnv.On("Read", "MARIAN_INSTALL_PATH").Return(marianPath, nil)
	mockEnv.On("Read", "VOCAB_BT_PATH").Return(vocabBtPath, nil)
	mockEnv.On("Read", "VOCAB_PATH").Return(vocabPath, nil)

	envReader := NewConfigEnvReader(mockEnv)
	config, err := envReader.Read()

	assert.NoError(t, err)
	assert.Equal(t, marianPath, config.MarianInstallPath)
	assert.Equal(t, vocabBtPath, config.VocabBtPath)
	assert.Equal(t, vocabPath, config.VocabPath)
	mockEnv.AssertExpectations(t)
}

func TestConfigEnvReader_MarianPathReturnEmpty(t *testing.T) {
	mockEnv := new(MockConfigApi.MockEnvReaderInterface)

	marianPath := ""
	vocabBtPath := "/path/to/vocab_bt"
	vocabPath := "/path/to/vocab"

	mockEnv.On("LoadFileEnv").Return(nil)
	mockEnv.On("Read", "MARIAN_INSTALL_PATH").Return(marianPath, nil)
	mockEnv.On("Read", "VOCAB_BT_PATH").Return(vocabBtPath, nil)
	mockEnv.On("Read", "VOCAB_PATH").Return(vocabPath, nil)

	envReader := NewConfigEnvReader(mockEnv)
	config, err := envReader.Read()

	assert.Error(t, err)
	assert.Equal(t, "", config.MarianInstallPath)
	assert.Equal(t, "", config.VocabBtPath)
	assert.Equal(t, "", config.VocabPath)
	mockEnv.AssertExpectations(t)
}

func TestConfigEnvReader_VocabBtPathReturnEmpty(t *testing.T) {
	mockEnv := new(MockConfigApi.MockEnvReaderInterface)

	marianPath := "/path/to/marian"
	vocabBtPath := ""
	vocabPath := "/path/to/vocab"

	mockEnv.On("LoadFileEnv").Return(nil)
	mockEnv.On("Read", "MARIAN_INSTALL_PATH").Return(marianPath, nil)
	mockEnv.On("Read", "VOCAB_BT_PATH").Return(vocabBtPath, nil)
	mockEnv.On("Read", "VOCAB_PATH").Return(vocabPath, nil)

	envReader := NewConfigEnvReader(mockEnv)
	config, err := envReader.Read()

	assert.Error(t, err)
	assert.Equal(t, "", config.MarianInstallPath)
	assert.Equal(t, "", config.VocabBtPath)
	assert.Equal(t, "", config.VocabPath)
	mockEnv.AssertExpectations(t)
}

func TestConfigEnvReader_VocabPathReturnEmpty(t *testing.T) {
	mockEnv := new(MockConfigApi.MockEnvReaderInterface)

	marianPath := "/path/to/marian"
	vocabBtPath := "/path/to/vocab_bt"
	vocabPath := ""

	mockEnv.On("LoadFileEnv").Return(nil)
	mockEnv.On("Read", "MARIAN_INSTALL_PATH").Return(marianPath, nil)
	mockEnv.On("Read", "VOCAB_BT_PATH").Return(vocabBtPath, nil)
	mockEnv.On("Read", "VOCAB_PATH").Return(vocabPath, nil)

	envReader := NewConfigEnvReader(mockEnv)
	config, err := envReader.Read()

	assert.Error(t, err)
	assert.Equal(t, "", config.MarianInstallPath)
	assert.Equal(t, "", config.VocabBtPath)
	assert.Equal(t, "", config.VocabPath)
	mockEnv.AssertExpectations(t)
}

func TestConfigEnvReader_LoadFileEnvError(t *testing.T) {
	mockEnv := new(MockConfigApi.MockEnvReaderInterface)
	mockEnv.On("LoadFileEnv").Return(errors.New("load error"))

	envReader := NewConfigEnvReader(mockEnv)
	config, err := envReader.Read()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load environment variables")
	assert.Equal(t, "", config.MarianInstallPath)
	assert.Equal(t, "", config.VocabBtPath)
	assert.Equal(t, "", config.VocabPath)
	mockEnv.AssertExpectations(t)
}

func TestConfigEnvReader_ReadMarianPathError(t *testing.T) {
	mockEnv := new(MockConfigApi.MockEnvReaderInterface)
	mockEnv.On("LoadFileEnv").Return(nil)
	mockEnv.On("Read", "MARIAN_INSTALL_PATH").Return("", errors.New("missing"))

	envReader := NewConfigEnvReader(mockEnv)
	config, err := envReader.Read()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read MARIAN_INSTALL_PATH")
	assert.Equal(t, "", config.MarianInstallPath)
	assert.Equal(t, "", config.VocabBtPath)
	assert.Equal(t, "", config.VocabPath)
	mockEnv.AssertExpectations(t)
}

func TestConfigEnvReader_ReadVocabBtPathError(t *testing.T) {
	mockEnv := new(MockConfigApi.MockEnvReaderInterface)
	mockEnv.On("LoadFileEnv").Return(nil)
	mockEnv.On("Read", "MARIAN_INSTALL_PATH").Return("/path", nil)
	mockEnv.On("Read", "VOCAB_BT_PATH").Return("", errors.New("missing"))

	envReader := NewConfigEnvReader(mockEnv)
	config, err := envReader.Read()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read VOCAB_BT_PATH")
	assert.Equal(t, "", config.MarianInstallPath)
	assert.Equal(t, "", config.VocabBtPath)
	assert.Equal(t, "", config.VocabPath)
	mockEnv.AssertExpectations(t)
}

func TestConfigEnvReader_ReadVocabPathError(t *testing.T) {
	mockEnv := new(MockConfigApi.MockEnvReaderInterface)
	mockEnv.On("LoadFileEnv").Return(nil)
	mockEnv.On("Read", "MARIAN_INSTALL_PATH").Return("/path", nil)
	mockEnv.On("Read", "VOCAB_BT_PATH").Return("/bt/path", nil)
	mockEnv.On("Read", "VOCAB_PATH").Return("", errors.New("missing"))

	envReader := NewConfigEnvReader(mockEnv)
	config, err := envReader.Read()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read VOCAB_PATH")
	assert.Equal(t, "", config.MarianInstallPath)
	assert.Equal(t, "", config.VocabBtPath)
	assert.Equal(t, "", config.VocabPath)
	mockEnv.AssertExpectations(t)
}
