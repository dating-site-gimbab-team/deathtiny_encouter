package routes

import (
	"deathtiny_encounters/controllers"
	"deathtiny_encounters/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.POST("/register", controllers.RegisterUser)
	router.GET("/login/google", controllers.GoogleLogin)
	router.GET("/auth/google/callback", controllers.GoogleAuthCallback)

	// Protected routes
	auth := router.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/profile", controllers.GetProfile)
		auth.PUT("/profile", controllers.UpdateProfile)
		auth.POST("/match", controllers.CreateMatch)
		auth.GET("/matches", controllers.GetMatches)
		auth.POST("/message", controllers.SendMessage)
		auth.GET("/messages", controllers.GetMessages)
	}
}
