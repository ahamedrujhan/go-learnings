package routes

import (
	"Event_Management/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// authentication route group
	authRoutes := server.Group("/")
	// pass the middleware
	authRoutes.Use(middlewares.Authenticate)
	// add the routes
	authRoutes.GET("/events", getEvents)
	authRoutes.POST("/events", middlewares.GetUserIdFromToken, createEvent)
	authRoutes.GET("/events/:id", getEventById)
	authRoutes.PUT("/events/:id", updateEventById)
	authRoutes.DELETE("/events/:id", deleteEventById)

	// public routes
	// signup
	server.POST("/signup", signUp)
	// login
	server.POST("/login", login)

}
