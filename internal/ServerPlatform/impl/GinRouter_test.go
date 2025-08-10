package ServerCore

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"TranslateServer/internal/ServerPlatform/api"
	"TranslateServer/internal/ServerPlatform/mock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGinRouter_GETRequestShouldCallHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockHandler := new(MockServerInterface.MockServerInterface)
	mockHandler.On("Callback", http.StatusOK, "ok").Return()

	router := NewGinRouter()
	router.GET("/test", func(c ServerCoreApi.HandlerInterface) {
		mockHandler.Callback(http.StatusOK, "ok")
	})

	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	router.engine.ServeHTTP(w, req)
}

func TestGinRouter_RunTest(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockHandler := new(MockServerInterface.MockServerInterface)
	mockHandler.On("Callback", http.StatusOK, "ok").Return()

	router := NewGinRouter()
	router.GET("/test", func(c ServerCoreApi.HandlerInterface) {
		mockHandler.Callback(http.StatusOK, "ok")
	})

	go func() {
		err := router.Run("0.0.0.0:5000")
		if err != nil {
			assert.Fail(t, "Server failed to start")
		}
	}()
	time.Sleep(250 * time.Millisecond)

	resp, err := http.Get("http://0.0.0.0:5000/test")
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()
}
