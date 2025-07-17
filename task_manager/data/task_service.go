package data

import (
	"errors"
	"time"

	"task_manager/models"
)

var tasks = []models.Task{
	{
		ID:          1,
		Title:       "Sample Task 1",
		Description: "This is a sample task",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      models.StatusPending,
	},
	{
		ID:          2,
		Title:       "Sample Task 2",
		Description: "Another sample task",
		DueDate:     time.Now().Add(48 * time.Hour),
		Status:      models.StatusCompleted,
	},
}

var nextID = 3

func GetTasks() []models.Task {
	return tasks
}

func GetTaskByID(id int) (*models.Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func CreateTask(newTask models.Task) models.Task {
	newTask.ID = nextID
	nextID++

	if newTask.Status == "" {
		newTask.Status = models.StatusPending
	}

	tasks = append(tasks, newTask)
	return newTask
}

func UpdateTask(id int, updatedTask models.Task) (*models.Task, error) {
	for i := range tasks {
		if tasks[i].ID == id {
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if !updatedTask.DueDate.IsZero() {
				tasks[i].DueDate = updatedTask.DueDate
			}
			if updatedTask.Status != "" {
				tasks[i].Status = updatedTask.Status
			}
			return &tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func DeleteTask(id int) error {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
