package TranslatorImpl

import (
	"errors"
	"fmt"

	"TranslateServer/internal/SentenceFormatter/api"
	"TranslateServer/internal/Supervisor/api"
	"TranslateServer/internal/Translator/api"
)

type VocabularyObject struct {
	vocabulary TranslatorApi.VocabularyInterface
	isInit     bool
}

type VocabularyAdapterManager struct {
	supervisorInterface        SupervisorApi.SupervisorInterface
	vocabularyList             []VocabularyObject
	sentenceFormatterInterface SentenceFormatterApi.SentenceFormatterInterface
}

func NewVocabularyAdapterManager(supervisorInterface SupervisorApi.SupervisorInterface,
	sentenceFormatterInterface SentenceFormatterApi.SentenceFormatterInterface) TranslatorApi.VocabularyAdapterManagerInterface {
	return &VocabularyAdapterManager{
		supervisorInterface:        supervisorInterface,
		vocabularyList:             []VocabularyObject{},
		sentenceFormatterInterface: sentenceFormatterInterface,
	}
}

func (m *VocabularyAdapterManager) checkIsSubscribe(vocabularyInterface TranslatorApi.VocabularyInterface) bool {
	result := false
	for idx := range m.vocabularyList {
		if m.vocabularyList[idx].vocabulary.GetId() == vocabularyInterface.GetId() {
			result = true
			break
		}
	}
	return result
}

func (m *VocabularyAdapterManager) Subscribe(vocabularyInterface TranslatorApi.VocabularyInterface) error {
	err := error(nil)
	if !m.checkIsSubscribe(vocabularyInterface) {
		m.vocabularyList = append(m.vocabularyList, VocabularyObject{vocabulary: vocabularyInterface, isInit: false})
	} else {
		err = errors.New(vocabularyInterface.GetId() + " already subscribed")
	}
	return err
}

func (m *VocabularyAdapterManager) Init() error {
	err := error(nil)
	for idx := range m.vocabularyList {
		if !m.vocabularyList[idx].isInit {
			vocabProp := m.vocabularyList[idx].vocabulary.GetProperties()
			channel, cmdErr := m.supervisorInterface.InitVocabTaskChannel(
				m.vocabularyList[idx].vocabulary.GetId(), vocabProp.Decoder, vocabProp.Model, vocabProp.Vocab)

			if cmdErr == nil {
				if cmdErr = m.vocabularyList[idx].vocabulary.RegisterInput(channel.In); cmdErr == nil {
					if cmdErr = m.vocabularyList[idx].vocabulary.RegisterOutput(channel.Out); cmdErr == nil {
						m.vocabularyList[idx].isInit = true
					}
				}
			}

			if cmdErr != nil {
				err = cmdErr
				break
			}
		}
	}
	return err
}

func (m *VocabularyAdapterManager) Deinit() error {
	err := error(nil)
	for idx := range m.vocabularyList {
		if m.vocabularyList[idx].isInit {
			err = m.vocabularyList[idx].vocabulary.Unregister()
			if err != nil {
				break
			}
			m.vocabularyList[idx].isInit = false
		}
	}
	return err
}

func (m *VocabularyAdapterManager) Translate(id string, text string) (string, error) {
	err := fmt.Errorf("not found translator")
	result := ""

	for idx := range m.vocabularyList {
		if m.vocabularyList[idx].vocabulary.GetId() == id {
			if m.vocabularyList[idx].isInit {
				text = m.sentenceFormatterInterface.PrepareInput(text)
				result, err = m.vocabularyList[idx].vocabulary.Translate(text)
				if err == nil {
					result = m.sentenceFormatterInterface.CleanOutput(result)
				}
			} else {
				err = fmt.Errorf("translator is not init")
			}
			break
		}
	}

	return result, err
}
