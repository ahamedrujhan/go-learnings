package routes

import (
	"Event_Management/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	// getting the event id
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	// getting the user id
	userId := context.GetInt64("user_id")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	var register models.Register
	//verify the event is available on that id
	_, err = models.GetEventByid(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "event is not found"})
		return
	}

	// binding the event id and user id to register
	register.Event_id = int(eventId)
	register.User_id = int(userId)

	// save the event
	err = register.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "event registered!!!"})

}

func cancelRegistation(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	userId := context.GetInt64("user_id")

	// validate the event
	_, err = models.GetEventByid(eventId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "event is not found"})
		return
	}

	// delete the registration
	var register models.Register
	register.Event_id = int(eventId)
	register.User_id = int(userId)
	err = register.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "registration cancelled"})
}
