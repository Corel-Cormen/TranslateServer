package OsPlatformImpl

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignal_Success(t *testing.T) {
	cmd := exec.Command("sleep", "10")
	assert.NoError(t, cmd.Start())
	defer cmd.Process.Kill()

	p := NewProcessFacade(cmd.Process)
	assert.NoError(t, p.Signal(0))
}
