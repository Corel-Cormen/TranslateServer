package OsPlatformImpl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileFacade_Read_Success(t *testing.T) {
	content := []byte("Hello, world!")
	tmpFile, err := os.CreateTemp("", "filefacade_test")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write(content)
	assert.NoError(t, err)
	tmpFile.Close()

	file, err := os.Open(tmpFile.Name())
	assert.NoError(t, err)
	defer file.Close()

	facade := NewFileFacade(file)
	readData, err := facade.Read()

	assert.NoError(t, err)
	assert.Equal(t, content, readData)
}

func TestFileFacade_Close_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "filefacade_test")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	file, err := os.Open(tmpFile.Name())
	assert.NoError(t, err)

	facade := NewFileFacade(file)
	err = facade.Close()

	assert.NoError(t, err)
}

func TestFileFacade_Read_NilFile(t *testing.T) {
	facade := NewFileFacade(nil)
	data, err := facade.Read()

	assert.ErrorIs(t, err, os.ErrInvalid)
	assert.Equal(t, []byte{}, data)
}

func TestFileFacade_Close_NilFile(t *testing.T) {
	facade := NewFileFacade(nil)
	err := facade.Close()

	assert.NoError(t, err)
}
