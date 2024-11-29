package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kartik1112/Memories-backend/controllers"
	"github.com/kartik1112/Memories-backend/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/status", controllers.StatusCheck)
	server.POST("/signup", controllers.Signup)
	server.POST("/login", controllers.Login)

	protected := server.Group("/")
	protected.Use(middlewares.CheckAuthentication)
	protected.POST("/event/create", controllers.CreateEvent)
	protected.POST("/events/join", controllers.JoinEvent)
	protected.GET("/events/fetch", controllers.FetchEvents)
}
