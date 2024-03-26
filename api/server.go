package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"keyValueStorage/config"
	_ "keyValueStorage/docs"
	"keyValueStorage/store"
	"log"
)

func StartServer(service *store.KeyValueStore) {
	cfg := config.LoadConfig()
	router := gin.Default()
	handlers := NewHandlers(service)
	handlers.RegisterRoutes(router)

	url := ginSwagger.URL(cfg.SwaggerUI)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	err := router.Run(cfg.Port)
	if err != nil {
		log.Println("Error starting server: ", err)
		return
	}
}
