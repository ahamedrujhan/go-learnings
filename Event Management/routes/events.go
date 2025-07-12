package routes

import (
	"Event_Management/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events.  Try again later. "})
		return
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	// parse the json from request body
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	// get user id from token
	userId := context.GetInt64("user_id")
	// bind the user id extracted from token to event
	event.UserID = int(userId)

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	event, err := models.GetEventByid(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func updateEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	userId := context.GetInt64("user_id")

	// validate the event id
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
		return
	}

	// verify the event is available for the eventID
	existingEvent, err := models.GetEventByid(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Event not found"})
		return
	}
	var updatedEvent models.Event
	// bind the event id to event
	updatedEvent.ID = eventId
	// bind the user id to event
	updatedEvent.UserID = int(userId)
	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	// validate the only event user id can update
	if int(userId) != existingEvent.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to update the event"})
		return
	}
	// update the event
	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated!", "event": updatedEvent})

}

func deleteEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	userId := int(context.GetInt64("user_id"))

	// validate the event id
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Colud not parse event id"})
		return
	}

	// verify the event is available for the event id
	event, err := models.GetEventByid(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Event not found"})
		return
	}

	// validate the token user and event user
	if userId != event.UserID {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized to delete the event"})
		return
	}

	// delete the event
	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})

}
