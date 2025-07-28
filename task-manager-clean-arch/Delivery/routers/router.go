package routers

import (
	"task-manager-clean-arch/Delivery/controllers"
	"task-manager-clean-arch/Domain"
	infrastructure "task-manager-clean-arch/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(
	taskController *controllers.TaskController,
	userController *controllers.UserController,
	jwtService domain.JWTService,
) *gin.Engine {
	router := gin.Default()

	// Public routes
	router.POST("/register", userController.RegisterUser)
	router.POST("/login", userController.Login)

	// Protected routes
	auth := router.Group("/")
	auth.Use(infrastructure.AuthMiddleware(jwtService))

	// Task routes requiring authentication
	auth.GET("/tasks", taskController.GetTasks)
	auth.GET("/tasks/:id", taskController.GetTaskByID)

	// Admin-only routes
	admin := auth.Group("/")
	admin.Use(infrastructure.AdminOnlyMiddleware())

	admin.POST("/tasks", taskController.CreateTask)
	admin.PUT("/tasks/:id", taskController.UpdateTask)
	admin.DELETE("/tasks/:id", taskController.DeleteTask)

	return router
} 