package OsPlatformImpl

import (
	"os"
	"io"
)

type FileFasade struct{
	file *os.File
}

func NewFileFasade(file *os.File) *FileFasade {
	return &FileFasade{
		file: file,
	}
}

func (f *FileFasade) Close() error {
	if f.file != nil {
		return f.file.Close()
	}
	return nil
}

func (f *FileFasade) Read() ([]byte, error) {
	if f.file != nil {
		data, err := io.ReadAll(f.file)
		if err != nil {
			return []byte{}, err
		}
		return data, nil
	}
	return []byte{}, os.ErrInvalid
}
