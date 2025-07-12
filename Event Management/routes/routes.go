package routes

import (
	"Event_Management/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	// authentication route group
	authRoutes := server.Group("/")
	// pass the middleware
	authRoutes.Use(middlewares.Authenticate, middlewares.GetUserIdFromToken)
	// add the routes
	authRoutes.GET("/events", getEvents)
	authRoutes.POST("/events", createEvent)
	authRoutes.GET("/events/:id", getEventById)
	authRoutes.PUT("/events/:id", updateEventById)
	authRoutes.DELETE("/events/:id", deleteEventById)
	// registering routes
	authRoutes.POST("/events/:id/register", registerForEvent)
	authRoutes.DELETE("/events/:id/register", cancelRegistation)

	// public routes
	// signup
	server.POST("/signup", signUp)
	// login
	server.POST("/login", login)

}
