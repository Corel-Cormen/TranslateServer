package ServerCore

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGinContextHandler_TextCallbackResponseTest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ginHandler := &GinContextHandler{Context: c}

	expectHttpStatus := http.StatusOK
	expectCallbackResponse := "(｡◕‿◕｡)"
	expectCallbackHeader := "text/plain"
	ginHandler.TextCallback(expectHttpStatus, expectCallbackResponse)

	assert.Equal(t, expectHttpStatus, w.Code)
	assert.Equal(t, expectCallbackHeader, w.Header().Get("Content-Type"))
	assert.Equal(t, expectCallbackResponse, strings.TrimSpace(w.Body.String()))
}

func TestGinContextHandler_JSONCallbackResponseTest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	ginHandler := &GinContextHandler{Context: c}

	expectHttpStatus := http.StatusOK
	expectCallbackResponse := "(｡◕‿◕｡)"
	expectCallbackHeader := "application/json"
	ginHandler.JsonCallback(expectHttpStatus, expectCallbackResponse)

	assert.Equal(t, expectHttpStatus, w.Code)
	assert.Equal(t, expectCallbackHeader, w.Header().Get("Content-Type"))
	assert.Equal(t, "\""+expectCallbackResponse+"\"", strings.TrimSpace(w.Body.String()))
}

func TestGinContextHandler_BindJSONSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)

	jsonBody := `{"metadata": "data"}`
	req := httptest.NewRequest("POST", "/test", strings.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ginHandler := &GinContextHandler{Context: c}

	var payload struct {
		Metadata string `json:"metadata"`
	}

	err := ginHandler.BindJSON(&payload)
	assert.NoError(t, err)
	assert.Equal(t, payload.Metadata, "data")
}
