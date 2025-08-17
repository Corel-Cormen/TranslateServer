package OsPlatformImpl

import (
	"bufio"
	"os"
	"os/exec"
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

func TestReadFile_Success(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())
	tmpFile.Write([]byte("test content"))
	tmpFile.Close()

	osFacade := &OsFacade{}
	content, readErr := osFacade.ReadFile(tmpFile.Name())

	assert.NoError(t, readErr)
	assert.Equal(t, "test content", string(content))
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

func TestAsyncCommand_Echo(t *testing.T) {
	osFacade := &OsFacade{}

	prop, err := osFacade.AsyncCommand("echo", "hello")
	assert.NoError(t, err)
	assert.True(t, prop.Pid > 0)
	assert.NotNil(t, prop.In)
	assert.NotNil(t, prop.Out)
	assert.NotNil(t, prop.Err)

	scanner := bufio.NewScanner(prop.Out)
	scanner.Scan()
	output := scanner.Text()
	assert.Equal(t, "hello", output)
}

func TestAsyncCommand_InvalidCommand(t *testing.T) {
	osFacade := &OsFacade{}
	_, err := osFacade.AsyncCommand("not-existing-command")
	assert.Error(t, err)
}

func TestGetProcess_ValidPid(t *testing.T) {
	cmd := exec.Command("sleep", "10")
	assert.NoError(t, cmd.Start())
	defer cmd.Process.Kill()

	osFacade := &OsFacade{}
	proc, err := osFacade.GetProcess(cmd.Process.Pid)
	assert.NoError(t, err)
	assert.NotNil(t, proc)
}
