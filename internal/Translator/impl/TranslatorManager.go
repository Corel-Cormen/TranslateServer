package TranslatorImpl

import (
	"TranslateServer/internal/Config/api"
	"TranslateServer/internal/Translator/api"
	"fmt"
)

type TranslatorManager struct {
	vocabularyManagerInterface TranslatorApi.VocabularyAdapterManagerInterface

	en2plBT Vocabulary
	en2pl   Vocabulary
}

func NewTranslatorManager(vocabularyManagerInterface TranslatorApi.VocabularyAdapterManagerInterface) TranslatorApi.TranslatorInterface {
	return &TranslatorManager{vocabularyManagerInterface: vocabularyManagerInterface}
}

func (t *TranslatorManager) Configure(cfg ConfigApi.ConfigData) error {
	decoderProgram := cfg.MarianInstallPath + "/marian-decoder"

	t.en2plBT.Id = "en-pl-BT"
	t.en2plBT.Decoder = decoderProgram
	t.en2plBT.Model = cfg.VocabBtPath + "/opus+bt.spm32k-spm32k.transformer-align.model1.npz.best-perplexity.npz"
	t.en2plBT.Vocab = cfg.VocabBtPath + "/opus+bt.spm32k-spm32k.vocab.yml"

	t.en2pl.Id = "en-pl"
	t.en2pl.Decoder = decoderProgram
	t.en2pl.Model = cfg.VocabPath + "/opus.spm32k-spm32k.transformer.model1.npz.best-perplexity.npz"
	t.en2pl.Vocab = cfg.VocabPath + "/opus.spm32k-spm32k.vocab.yml"

	err := t.vocabularyManagerInterface.Subscribe(&t.en2plBT)
	if err == nil {
		err = t.vocabularyManagerInterface.Subscribe(&t.en2pl)
	}

	return err
}

func (t *TranslatorManager) Run() error {
	fmt.Println("Translator started")

	return t.vocabularyManagerInterface.Init()
}

func (t *TranslatorManager) Translate(language string, text string) (string, error) {
	return t.vocabularyManagerInterface.Translate(language, text)
}
