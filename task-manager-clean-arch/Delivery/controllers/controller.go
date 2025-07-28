package controllers

import (
	"net/http"

	domain "task-manager-clean-arch/Domain"
	usecases "task-manager-clean-arch/Usecases"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUseCase *usecases.TaskUseCase
}

func NewTaskController(taskUseCase *usecases.TaskUseCase) *TaskController {
	return &TaskController{
		taskUseCase: taskUseCase,
	}
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.taskUseCase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")

	task, err := tc.taskUseCase.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var newTask domain.Task

	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	err := tc.taskUseCase.CreateTask(&newTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")

	var updatedTask domain.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON input"})
		return
	}

	task, err := tc.taskUseCase.UpdateTask(id, &updatedTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	err := tc.taskUseCase.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

type UserController struct {
	userUseCase *usecases.UserUseCase
}

func NewUserController(userUseCase *usecases.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := uc.userUseCase.RegisterUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role,
		},
	})
}

func (uc *UserController) Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := uc.userUseCase.LoginUser(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
