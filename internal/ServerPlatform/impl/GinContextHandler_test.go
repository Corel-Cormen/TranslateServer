package ServerCore

import (
	"net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestGinContextHandler_CallbackResponseTest(t *testing.T) {
    gin.SetMode(gin.TestMode)

    w := httptest.NewRecorder()
    c, _ := gin.CreateTestContext(w)

    ginHandler := &GinContextHandler{Context: c}

	expectHttpStatus := http.StatusOK
	expectCallbackResponse := "(｡◕‿◕｡)"
	expectCallbackHeader := "application/text"
    ginHandler.Callback(expectHttpStatus, expectCallbackResponse)

    assert.Equal(t, expectHttpStatus, w.Code)
    assert.Equal(t, expectCallbackHeader, w.Header().Get("Content-Type"))
    assert.Equal(t, expectCallbackResponse, strings.TrimSpace(w.Body.String()))
}
