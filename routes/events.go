package routes

import (
	"example/restapi/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve data"})
		return
	}
	context.JSON(http.StatusOK, events)
}

// get Event by ID
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "request valid event id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}
	context.JSON(http.StatusOK, event)

}
func createEvent(context *gin.Context) {

	var event models.Event

	//store post data to event variable
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Data Required",
		})
		return //skip later func block
	}
	userId := context.GetInt64("userID")
	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not create event"})
		return

	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created",
		"Event":   event,
	})

}
func updateEvent(context *gin.Context) {
	var updatedEvent models.Event
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "request valid event id"})
		return
	}
	userId := context.GetInt64("userID")

	event, err := models.GetEventByID(eventId)
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update event"})

		return
	}
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the event"})
		return
	}

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Data Required",
		})
		return //skip later func block
	}
	updatedEvent.ID = eventId
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update the event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "update sucessed",
	})
}
func deleteEvent(context *gin.Context) {
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
			"message": "Could not fetch the event",
		})
		return
	}
	userId := context.GetInt64("userID")
	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete event"})

		return
	}

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not Delete the event",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "event deleted",
	})

}
