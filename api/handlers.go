package api

import (
	"github.com/gin-gonic/gin"
	"keyValueStorage/store"
	"net/http"
	"time"
)

type Handlers struct {
	service store.KeyValueService
}

func NewHandlers(service store.KeyValueService) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) RegisterRoutes(router *gin.Engine) {
	router.POST("/set", h.setKey)
	router.GET("/get/:key", h.getKey)
	router.DELETE("/delete/:key", h.deleteKey)
}

type SetValueRequest struct {
	Key   string      `json:"key" binding:"required"`
	Value interface{} `json:"value" binding:"required"`
	TTL   int64       `json:"ttl"`
}

func (h *Handlers) setKey(c *gin.Context) {
	var request SetValueRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if request.TTL < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TTL is less than 0"})
		return
	}

	err := h.service.Set(request.Key, request.Value, time.Duration(request.TTL)*time.Second)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Value set successfully"})
}

func (h *Handlers) getKey(c *gin.Context) {
	key := c.Param("key")

	value, err := h.service.Get(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if value == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Key not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"value": value})
}

func (h *Handlers) deleteKey(c *gin.Context) {
	key := c.Param("key")

	err := h.service.Delete(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key deleted successfully"})
}
