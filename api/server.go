package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "keyValueStorage/docs"
	"keyValueStorage/store"
	"log"
)

func StartServer(service *store.KeyValueStore) {
	router := gin.Default()
	handlers := NewHandlers(service)
	handlers.RegisterRoutes(router)

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	log.Println("Swagger UI is available at http://localhost:8080/swagger/index.html")

	err := router.Run(":8080")
	if err != nil {
		log.Println("Error starting server: ", err)
		return
	}
}
