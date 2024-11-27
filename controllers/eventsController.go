package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/Memories-backend/initializers"
	"github.com/kartik1112/Memories-backend/models"
	"github.com/kartik1112/Memories-backend/utils"
)

func CreateEvent(ctx *gin.Context) {

	var event models.Events
	err := ctx.ShouldBindJSON(&event)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid JSON Provided",
		})
		return
	}

	userId, exists := ctx.Get("userId")

	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "User ID does not exist",
		})
	}

	event.Code = utils.GenerateEventCode(userId.(uint), event.EventName)
	event.CreatedBy = userId.(uint)

	result := initializers.DB.Create(&event)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Could Not Create Event",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Event Created Successfully",
		"eventCode": event.Code,
	})
}
