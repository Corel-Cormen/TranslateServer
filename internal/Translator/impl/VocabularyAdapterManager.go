package TranslatorImpl

import (
	"errors"
	"fmt"

	"TranslateServer/internal/OsPlatform/api"
	"TranslateServer/internal/Translator/api"
)

type VocabularyObject struct {
	vocabulary TranslatorApi.VocabularyInterface
	isInit     bool
}

type VocabularyAdapterManager struct {
	osInterface    OsPlatformApi.OsInterface
	vocabularyList []VocabularyObject
}

func NewVocabularyAdapterManager(osPlatform OsPlatformApi.OsInterface) TranslatorApi.VocabularyAdapterManagerInterface {
	return &VocabularyAdapterManager{osInterface: osPlatform}
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

func (m *VocabularyAdapterManager) verifyInputFiles(vocabularyInterface TranslatorApi.VocabularyInterface) bool {
	vocabularyProperties := vocabularyInterface.GetProperties()
	return m.osInterface.FileExist(vocabularyProperties.Decoder) &&
		m.osInterface.FileExist(vocabularyProperties.Model) &&
		m.osInterface.FileExist(vocabularyProperties.Vocab)
}

func (m *VocabularyAdapterManager) Subscribe(vocabularyInterface TranslatorApi.VocabularyInterface) error {
	err := error(nil)
	if !m.checkIsSubscribe(vocabularyInterface) {
		if m.verifyInputFiles(vocabularyInterface) {
			m.vocabularyList = append(m.vocabularyList, VocabularyObject{vocabulary: vocabularyInterface, isInit: false})
		} else {
			err = errors.New(vocabularyInterface.GetId() + "configuration files not found")
		}
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
			stdIn, stdOut, stdErr, cmdErr := m.osInterface.AsyncCommand(
				vocabProp.Decoder,
				"-m", vocabProp.Model,
				"-v", vocabProp.Vocab, vocabProp.Vocab,
			)

			if cmdErr == nil {
				if cmdErr = m.vocabularyList[idx].vocabulary.RegisterInput(stdIn); cmdErr == nil {
					if cmdErr = m.vocabularyList[idx].vocabulary.RegisterOutput(stdOut); cmdErr == nil {
						if cmdErr = m.vocabularyList[idx].vocabulary.RegisterLog(stdErr); cmdErr == nil {
							m.vocabularyList[idx].isInit = true
						}
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
				result, err = m.vocabularyList[idx].vocabulary.Translate(text)
			} else {
				err = fmt.Errorf("translator is not init")
			}
			break
		}
	}

	return result, err
}
