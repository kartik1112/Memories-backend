package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"Health" : "All Systems are up and Running",
	})
}