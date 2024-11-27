package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/kartik1112/Memories-backend/initializers"
	"github.com/kartik1112/Memories-backend/models"
	"github.com/kartik1112/Memories-backend/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	err := initializers.DB.AutoMigrate(models.User{}, models.Events{})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run()
}
