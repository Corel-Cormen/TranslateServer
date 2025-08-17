package TranslatorApi

import "io"

type VocabularyProperties struct {
	Decoder string
	Model   string
	Vocab   string
}

type VocabularyInterface interface {
	GetId() string
	GetProperties() VocabularyProperties
	Translate(text string) (string, error)
	RegisterInput(closer io.WriteCloser) error
	RegisterOutput(closer io.ReadCloser) error
	Unregister() error
}
