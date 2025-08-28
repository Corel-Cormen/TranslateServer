package SentenceFormaterInstance

import (
	"sync"

	"TranslateServer/internal/SentenceFormatter/api"
	"TranslateServer/internal/SentenceFormatter/impl"
)

var (
	sentenceFormatter     SentenceFormatterApi.SentenceFormatterInterface
	onceSentenceFormatter sync.Once
)

func GetSentenceFormaterInstance() SentenceFormatterApi.SentenceFormatterInterface {
	onceSentenceFormatter.Do(func() {
		sentenceFormatter = &SentenceFormatterImpl.MarianSentenceFormatter{}
	})
	return sentenceFormatter
}
