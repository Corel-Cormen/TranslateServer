package ConfigCore

import (
	"errors"
	"testing"

	"TranslateServer/internal/Config/api"
	"TranslateServer/internal/Config/mock"
	"github.com/stretchr/testify/assert"
)

func TestConfigSysEnv_Init_Success(t *testing.T) {
	mockReader := new(MockConfigApi.MockConfigReaderInterface)
	expectedConfig := ConfigApi.ConfigData{
		MarianInstallPath: "/test/marian",
		VocabBtPath:       "/test/vocab_bt",
		VocabPath:         "/test/vocab",
	}

	mockReader.On("Read").Return(expectedConfig, nil)

	sysEnv := NewConfigSysEnv(mockReader)
	err := sysEnv.Init()

	assert.NoError(t, err)

	cfg, err := sysEnv.Get()
	assert.NoError(t, err)
	assert.Equal(t, expectedConfig, cfg)

	mockReader.AssertExpectations(t)
}

func TestConfigSysEnv_Init_Failure(t *testing.T) {
	mockReader := new(MockConfigApi.MockConfigReaderInterface)
	mockReader.On("Read").Return(ConfigApi.ConfigData{}, errors.New("read error"))

	sysEnv := NewConfigSysEnv(mockReader)
	err := sysEnv.Init()

	assert.Error(t, err)
	assert.EqualError(t, err, "read error")

	_, getErr := sysEnv.Get()
	assert.Error(t, getErr)
	assert.EqualError(t, getErr, "ConfigSysEnv not initialized")

	mockReader.AssertExpectations(t)
}

func TestConfigSysEnv_Get_BeforeInit(t *testing.T) {
	mockReader := new(MockConfigApi.MockConfigReaderInterface)
	sysEnv := NewConfigSysEnv(mockReader)

	_, err := sysEnv.Get()
	assert.Error(t, err)
	assert.EqualError(t, err, "ConfigSysEnv not initialized")
}
