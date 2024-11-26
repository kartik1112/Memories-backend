package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func statusCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Health" : "All Systems are up and Running",
	})
}