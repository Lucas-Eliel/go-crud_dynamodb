package main

import (
	"crud_dynamodb/internal/controllers"
	"crud_dynamodb/internal/repositories"
	"crud_dynamodb/internal/services"
	"crud_dynamodb/pkg/db"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	clientdb := db.CreateLocalClient()

	log.Printf("Conectando com o DynamoDB")

	controllers.RegisterHandlers(router, services.NewService(repositories.NewRepository(clientdb)))
	log.Println("Configurando Routers")

	router.Run("localhost:8090")
}
