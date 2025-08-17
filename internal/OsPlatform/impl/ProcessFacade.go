package OsPlatformImpl

import (
	"os"
	"syscall"
)

type ProcessFacade struct {
	process *os.Process
}

func NewProcessFacade(process *os.Process) *ProcessFacade {
	return &ProcessFacade{process: process}
}

func (p *ProcessFacade) Signal(signal int) error {
	return p.process.Signal(syscall.Signal(signal))
}
