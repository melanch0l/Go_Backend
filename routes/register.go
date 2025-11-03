package routes

import (
	"example/restapi/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerEvent(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Request Valid Event ID",
		})
		return
	}
	event, err := models.GetEventByID(eventID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not retrive event data",
		})
		return
	}
	err = event.Register(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not register to event",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Registration success",
	})
}
func cancelRegister(context *gin.Context) {
	userID := context.GetInt64("userID")
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Request Valid Event ID",
		})
		return
	}
	var event models.Event
	event.ID = eventID
	err = event.Cancel(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not cancel registration",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Cancellation success",
	})
}
