package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ricardofabila/arithmetic-calculator-backend/controllers"
	"github.com/ricardofabila/arithmetic-calculator-backend/middlewares"
)

func SetupRouter(operationController *controllers.OperationController) *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowAllOrigins = true
	config.AllowCredentials = true

	corsMiddleware := cors.New(config)
	router.Use(corsMiddleware)

	// Public Routes
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Protected Routes
	api := router.Group("/api/v1")
	api.Use(middlewares.JWTAuthMiddleware())

	api.POST("/operation", operationController.PerformOperation)
	api.GET("/records", controllers.GetRecords)
	api.DELETE("/records/:id", controllers.DeleteRecord)

	return router
}
