package TranslatorApi

type VocabularyAdapterManagerInterface interface {
	Subscribe(vocabularyInterface VocabularyInterface) error
	Init() error
	Deinit() error
	Translate(id string, text string) (string, error)
}
