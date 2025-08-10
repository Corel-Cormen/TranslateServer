package TranslatorInstance

import (
	"sync"

	"TranslateServer/internal/OsPlatform"
	"TranslateServer/internal/Translator/api"
	"TranslateServer/internal/Translator/impl"
)

var (
	translatorInstance        TranslatorApi.TranslatorInterface
	vocabularyManagerInstance TranslatorApi.VocabularyAdapterManagerInterface
	onceTranslatorInstance    sync.Once
)

func GetTranslatorInstance() TranslatorApi.TranslatorInterface {
	onceTranslatorInstance.Do(func() {
		vocabularyManagerInstance = TranslatorImpl.NewVocabularyAdapterManager(OsPlatform.GetOsInstance())
		translatorInstance = TranslatorImpl.NewTranslatorManager(vocabularyManagerInstance)
	})
	return translatorInstance
}
