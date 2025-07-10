package routes

import (
	"Event_Management/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// get all events
	server.GET("/events", getEvents)
	// create event
	server.POST("/events", middlewares.Authenticate, createEvent)
	// get event by id
	server.GET("/events/:id", getEventById)
	// update event by id
	server.PUT("/events/:id", updateEventById)
	// delete the event by id
	server.DELETE("/events/:id", deleteEventById)
	// signup
	server.POST("/signup", signUp)
	// login
	server.POST("/login", login)

}
