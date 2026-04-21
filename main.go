package main

import (
	"backend_gin.com/gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.SetupRoutes(server) // Set up routes using the routes package
	server.Run(":9090")
}



// Project Structuring
	// It is layered architecture
	// CSR = Controllers, Services, Repositories
	// routes , model
	// routes : api_routes , auth_routes
	// controllers : user_controller , auth_controller (http handlers)
	// services : user_service , auth_service (business logic)
	// repositories : user_repository , auth_repository (database interactions)