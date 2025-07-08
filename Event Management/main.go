package main

import (
	"Event_Management/db"
	"Event_Management/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB() // db initialization

	server := gin.Default() // gin default engine initialized

	routes.RegisterRoutes(server) // register the routes

	server.Run(":8080") // localhost:8080

}
