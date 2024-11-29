package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/Memories-backend/initializers"
	"github.com/kartik1112/Memories-backend/models"
)

func JoinEvent(ctx *gin.Context) {

	var userEvent models.UserEvent

	var eventCode models.EventCode
	var event models.Event
	userId := ctx.GetUint("userId")

	err := ctx.ShouldBindJSON(&eventCode)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Event Does Not Exist",
		})
		return
	}

	event.Code = eventCode.Code
	event.GetEventFromCode()

	userEvent.EventID = event.ID

	userEvent.UserID = userId

	initializers.DB.Create(&userEvent)

	// initializers.DB.Create()
	ctx.JSON(http.StatusOK, gin.H{
		"message": userEvent,
	})
	// initializers.DB.Create(&)
}
