package TranslatorInstance

import (
	"sync"

	"TranslateServer/internal/SentenceFormatter/instance"
	"TranslateServer/internal/Supervisor/instance"
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
		vocabularyManagerInstance = TranslatorImpl.NewVocabularyAdapterManager(
			SupervisorInstance.GetSupervisorInstance(),
			SentenceFormaterInstance.GetSentenceFormaterInstance())
		translatorInstance = TranslatorImpl.NewTranslatorManager(vocabularyManagerInstance)
	})
	return translatorInstance
}
