package ConfigCore

import (
	"errors"
	"testing"

	"TranslateServer/internal/OsPlatform/mock"
	"github.com/stretchr/testify/assert"
)

func TestEnvReader_LoadFileEnv_Success(t *testing.T) {
	mockOS := new(MockOsPlatformApi.MockOsInterface)
	mockFile := new(MockOsPlatformApi.MockFileInterface)
	envPath := ".env"

	content := []byte("VAR1=value1\nVAR2=value2\n")

	mockOS.On("FileExist", envPath).Return(true)
	mockOS.On("OpenFile", envPath).Return(mockFile, nil)
	mockFile.On("Close").Return(nil)
	mockFile.On("Read").Return(content, nil)
	mockOS.On("SetEnv", "VAR1", "value1").Return(nil)
	mockOS.On("SetEnv", "VAR2", "value2").Return(nil)

	reader := NewEnvReader(envPath, mockOS)
	err := reader.LoadFileEnv()

	assert.NoError(t, err)

	mockOS.AssertExpectations(t)
	mockFile.AssertExpectations(t)
}

func TestEnvReader_LoadFileEnv_EnvFileNotDetect(t *testing.T) {
	mockOS := new(MockOsPlatformApi.MockOsInterface)
	envPath := ".env"

	mockOS.On("FileExist", envPath).Return(false)

	reader := NewEnvReader(envPath, mockOS)
	err := reader.LoadFileEnv()

	assert.NoError(t, err)

	mockOS.AssertExpectations(t)
}

func TestEnvReader_LoadFileEnv_InvalidLine(t *testing.T) {
	mockOS := new(MockOsPlatformApi.MockOsInterface)
	mockFile := new(MockOsPlatformApi.MockFileInterface)
	envPath := ".env"

	content := []byte("INVALID_LINE\n")

	mockOS.On("FileExist", envPath).Return(true)
	mockOS.On("OpenFile", envPath).Return(mockFile, nil)
	mockFile.On("Close").Return(nil)
	mockFile.On("Read").Return(content, nil)

	reader := NewEnvReader(envPath, mockOS)
	err := reader.LoadFileEnv()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid env line")

	mockOS.AssertExpectations(t)
	mockFile.AssertExpectations(t)
}

func TestEnvReader_LoadFileEnv_ReadError(t *testing.T) {
	mockOS := new(MockOsPlatformApi.MockOsInterface)
	mockFile := new(MockOsPlatformApi.MockFileInterface)
	envPath := ".env"

	content := []byte("VAR1=value1\n")

	mockOS.On("FileExist", envPath).Return(true)
	mockOS.On("OpenFile", envPath).Return(mockFile, nil)
	mockFile.On("Close").Return(nil)
	mockFile.On("Read").Return(content, errors.New("read error"))

	reader := NewEnvReader(envPath, mockOS)
	err := reader.LoadFileEnv()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read env file")

	mockOS.AssertExpectations(t)
	mockFile.AssertExpectations(t)
}

func TestEnvReader_LoadFileEnv_SetEnvError(t *testing.T) {
	mockOS := new(MockOsPlatformApi.MockOsInterface)
	mockFile := new(MockOsPlatformApi.MockFileInterface)
	envPath := ".env"

	content := []byte("VAR1=value1\n")

	mockOS.On("FileExist", envPath).Return(true)
	mockOS.On("OpenFile", envPath).Return(mockFile, nil)
	mockFile.On("Close").Return(nil)
	mockFile.On("Read").Return(content, nil)
	mockOS.On("SetEnv", "VAR1", "value1").Return(errors.New("setenv failed"))

	reader := NewEnvReader(envPath, mockOS)
	err := reader.LoadFileEnv()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to set env variable VAR1")

	mockOS.AssertExpectations(t)
	mockFile.AssertExpectations(t)
}

func TestEnvReader_LoadFileEnv_OpenFileError(t *testing.T) {
	mockOS := new(MockOsPlatformApi.MockOsInterface)
	mockFile := new(MockOsPlatformApi.MockFileInterface)
	envPath := ".env"

	mockOS.On("FileExist", envPath).Return(true)
	mockOS.On("OpenFile", envPath).Return(mockFile, errors.New("open error"))
	mockFile.On("Close").Return(nil)

	reader := NewEnvReader(envPath, mockOS)
	err := reader.LoadFileEnv()

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to open env file")

	mockOS.AssertExpectations(t)
}

func TestEnvReader_Read_Success(t *testing.T) {
	mockOS := new(MockOsPlatformApi.MockOsInterface)
	envPath := ".env"

	mockOS.On("LookupEnv", "VAR1").Return("value1", true)

	reader := NewEnvReader(envPath, mockOS)
	val, err := reader.Read("VAR1")

	assert.NoError(t, err)
	assert.Equal(t, "value1", val)
}

func TestEnvReader_Read_EnvNotFound(t *testing.T) {
	mockOS := new(MockOsPlatformApi.MockOsInterface)
	envPath := ".env"

	mockOS.On("LookupEnv", "MISSING_VAR").Return("", false)

	reader := NewEnvReader(envPath, mockOS)
	_, err := reader.Read("MISSING_VAR")

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "environment variable MISSING_VAR not found")
}
