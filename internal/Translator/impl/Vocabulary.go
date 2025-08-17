package TranslatorImpl

import (
	"bufio"
	"fmt"
	"io"

	"TranslateServer/internal/Translator/api"
)

type Vocabulary struct {
	Id string

	Decoder string
	Model   string
	Vocab   string

	in      io.WriteCloser
	scanner *bufio.Scanner
}

func (v *Vocabulary) GetId() string {
	return v.Id
}

func (v *Vocabulary) GetProperties() TranslatorApi.VocabularyProperties {
	return TranslatorApi.VocabularyProperties{
		Decoder: v.Decoder,
		Model:   v.Model,
		Vocab:   v.Vocab,
	}
}

func (v *Vocabulary) Translate(text string) (string, error) {
	result := ""
	_, err := fmt.Fprintln(v.in, text)
	if err == nil {
		if v.scanner.Scan() {
			result = v.scanner.Text()
		}
	}
	return result, err
}

func (v *Vocabulary) RegisterInput(closer io.WriteCloser) error {
	v.in = closer
	return nil
}

func (v *Vocabulary) RegisterOutput(closer io.ReadCloser) error {
	v.scanner = bufio.NewScanner(closer)
	return nil
}

func (v *Vocabulary) Unregister() error {
	return v.in.Close()
}
