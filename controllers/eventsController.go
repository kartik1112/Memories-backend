package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kartik1112/Memories-backend/models"
)

func CreateEvent(ctx *gin.Context) {

	var event models.Event
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

	event.CreatedBy = userId.(uint)

	err = event.CreateEvent()

	if err != nil {
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

func FetchEvents(ctx *gin.Context) {

	userId, ok := ctx.Get("userId")

	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized Access",
		})
		return
	}

	eventsCreated, err := models.GetEventsCreatedByUserId(userId.(uint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could Not Fetch Events",
		})
		return
	}

	eventsJoined, err := models.GetEventsJoinedByUserId(userId.(uint))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could Not Fetch Events",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"created_events": eventsCreated,
		"joined_events":  eventsJoined,
	})

}
