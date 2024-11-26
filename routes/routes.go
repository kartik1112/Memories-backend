package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/status", statusCheck)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
