package OsPlatformImpl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileExist_True(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	osFacade := &OsFacade{}
	exists := osFacade.FileExist(tmpFile.Name())
	assert.True(t, exists)
}

func TestFileExist_False(t *testing.T) {
	osFacade := &OsFacade{}
	exists := osFacade.FileExist("/nonexistent/file/path")
	assert.False(t, exists)
}

func TestOpenFile_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	tmpFile.Write([]byte("test content"))
	tmpFile.Close()

	osFacade := &OsFacade{}
	fileInterface, err := osFacade.OpenFile(tmpFile.Name())

	assert.NoError(t, err)
	assert.NotNil(t, fileInterface)

	content, readErr := fileInterface.Read()
	assert.NoError(t, readErr)
	assert.Equal(t, "test content", string(content))

	closeErr := fileInterface.Close()
	assert.NoError(t, closeErr)
}

func TestOpenFile_Fail(t *testing.T) {
	osFacade := &OsFacade{}
	fileInterface, err := osFacade.OpenFile("/nonexistent/file/path")

	assert.Error(t, err)
	assert.Nil(t, fileInterface)
}

func TestSetEnvAndLookupEnv(t *testing.T) {
	osFacade := &OsFacade{}

	key := "UNIT_TEST_ENV_VAR"
	value := "TestValue"

	err := osFacade.SetEnv(key, value)
	assert.NoError(t, err)

	gotValue, exists := osFacade.LookupEnv(key)
	assert.True(t, exists)
	assert.Equal(t, value, gotValue)
}
