package ServerCore

import (
	"net/http"

	"TranslateServer/internal/ServerPlatform/api"
	"TranslateServer/internal/Translator/api"
)

type TranslateHandler struct {
	Translator TranslatorApi.TranslatorInterface
}

func (h *TranslateHandler) Handle(handler ServerCoreApi.HandlerInterface) {
	var req struct {
		Lang string `json:"lang"`
		Text string `json:"text"`
	}

	var code int
	var resp string

	if err := handler.BindJSON(&req); err == nil {
		var err error
		resp, err = h.Translator.Translate(req.Lang, req.Text)
		if err == nil {
			code = http.StatusOK
		} else {
			code = http.StatusBadRequest
			resp = err.Error()
		}
	} else {
		code = http.StatusBadRequest
		resp = "invalid JSON"
	}

	handler.JsonCallback(code, map[string]string{"result": resp})
}
