package test

import (
	"bytes"
	"encoding/json"
	"keyValueStorage/api"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"keyValueStorage/store/mock"
)

func TestSetKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := new(mock.KeyValueService)
	mockService.On("Set", "username", "john_doe", time.Duration(20)*time.Second).Return(nil)

	h := api.NewHandlers(mockService)
	h.RegisterRoutes(router)

	body := api.SetValueRequest{
		Key:   "username",
		Value: "john_doe",
		TTL:   20,
	}
	bodyBytes, _ := json.Marshal(body)

	req, _ := http.NewRequest(http.MethodPost, "/set", bytes.NewBuffer(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Value set successfully")

	mockService.AssertExpectations(t)
}

func TestGetKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := new(mock.KeyValueService)
	expectedValue := "john_doe"
	mockService.On("Get", "username").Return(expectedValue, nil)

	h := api.NewHandlers(mockService)
	h.RegisterRoutes(router)

	req, _ := http.NewRequest(http.MethodGet, "/get/username", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), expectedValue)

	mockService.AssertCalled(t, "Get", "username")
}

func TestDeleteKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockService := new(mock.KeyValueService)
	mockService.On("Delete", "username").Return(nil)

	h := api.NewHandlers(mockService)
	h.RegisterRoutes(router)

	req, _ := http.NewRequest(http.MethodDelete, "/delete/username", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Key deleted successfully")

	mockService.AssertCalled(t, "Delete", "username")
}
