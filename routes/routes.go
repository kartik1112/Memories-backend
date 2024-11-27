package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kartik1112/Memories-backend/controllers"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/status", controllers.StatusCheck)
	server.POST("/signup", controllers.Signup)
	server.POST("/login", controllers.Login)
}
