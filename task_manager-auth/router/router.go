package router

import (
	"github.com/gin-gonic/gin"
	"task_manager/controllers"
	"task_manager/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.Login)

	// require login routes
	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())

	auth.GET("/tasks", controllers.GetTasks)
	auth.GET("/tasks/:id", controllers.GetTaskByID)

	// Admin only routes
	admin := auth.Group("/")
	admin.Use(middleware.AdminOnlyMiddleware())

	admin.POST("/tasks", controllers.CreateTask)
	admin.PUT("/tasks/:id", controllers.UpdateTask)
	admin.DELETE("/tasks/:id", controllers.DeleteTask) 

	return router
}

