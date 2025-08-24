package ServerCore

import (
	"errors"
	"net/http"
	"testing"

	"TranslateServer/internal/ServerPlatform/mock"
	"TranslateServer/internal/Translator/mock"
	"github.com/stretchr/testify/mock"
)

func TestTranslateHandlerHandler_HandleTranslateSuccess(t *testing.T) {
	mockTranslator := new(MockTranslator.MockTranslatorInterface)
	translateHandler := &TranslateHandler{mockTranslator}
	mockHandler := new(MockServerInterface.MockServerInterface)

	mockHandler.On("BindJSON", mock.Anything).Return(nil)
	mockTranslator.On("Translate", mock.Anything, mock.Anything).Return("translate_text", nil)
	mockHandler.On("JsonCallback", http.StatusOK, map[string]string{"result": "translate_text"}).Return()

	translateHandler.Handle(mockHandler)

	mockTranslator.AssertExpectations(t)
	mockHandler.AssertExpectations(t)
}

func TestTranslateHandlerHandler_HandleTranslateFail(t *testing.T) {
	mockTranslator := new(MockTranslator.MockTranslatorInterface)
	translateHandler := &TranslateHandler{mockTranslator}
	mockHandler := new(MockServerInterface.MockServerInterface)

	mockHandler.On("BindJSON", mock.Anything).Return(nil)
	mockTranslator.On("Translate", mock.Anything, mock.Anything).Return("", errors.New("translate_error"))
	mockHandler.On("JsonCallback", http.StatusBadRequest, map[string]string{"result": "translate_error"}).Return()

	translateHandler.Handle(mockHandler)

	mockTranslator.AssertExpectations(t)
	mockHandler.AssertExpectations(t)
}

func TestTranslateHandlerHandler_HandleTranslateFail_JsonCallback(t *testing.T) {
	mockTranslator := new(MockTranslator.MockTranslatorInterface)
	translateHandler := &TranslateHandler{mockTranslator}
	mockHandler := new(MockServerInterface.MockServerInterface)

	mockHandler.On("BindJSON", mock.Anything).Return(errors.New("bind_error"))
	mockHandler.On("JsonCallback", http.StatusBadRequest, map[string]string{"result": "invalid JSON"}).Return()

	translateHandler.Handle(mockHandler)

	mockTranslator.AssertExpectations(t)
	mockHandler.AssertExpectations(t)
}
