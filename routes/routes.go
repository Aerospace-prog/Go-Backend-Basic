package routes

import (
	"backend_gin.com/gin/controller"
	"github.com/gin-gonic/gin"
	"backend_gin.com/gin/middleware"
)

func SetupRoutes(server *gin.Engine) {
	api := server.Group("/api")

	api.GET("/", controller.Greet)
	api.GET("/users/:name", controller.Name)
	api.GET("/search", controller.Search)


	auth := server.Group("/auth")
	
	auth.Use(middleware.AuthMiddleware()) // Apply authentication middleware to auth routes
	auth.POST("/login", controller.Login)
}