package OsPlatformImpl

import (
	"io"
	"os"
)

type FileFacade struct {
	file *os.File
}

func NewFileFacade(file *os.File) *FileFacade {
	return &FileFacade{
		file: file,
	}
}

func (f *FileFacade) Close() error {
	if f.file != nil {
		return f.file.Close()
	}
	return nil
}

func (f *FileFacade) Read() ([]byte, error) {
	if f.file != nil {
		return io.ReadAll(f.file)
	}
	return []byte{}, os.ErrInvalid
}
