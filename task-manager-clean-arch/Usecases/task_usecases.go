package usecases

import (
	"errors"
	"time"

	"task-manager-clean-arch/Domain"
)

type TaskUseCase struct {
	taskRepo domain.TaskRepository
}

func NewTaskUseCase(taskRepo domain.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: taskRepo,
	}
}

func (uc *TaskUseCase) CreateTask(task *domain.Task) error {
	if task.Title == "" {
		return errors.New("task title is required")
	}

	if task.Description == "" {
		return errors.New("task description is required")
	}

	if task.Status == "" {
		task.Status = domain.StatusPending
	}

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	return uc.taskRepo.Create(task)
}

func (uc *TaskUseCase) GetTaskByID(id string) (*domain.Task, error) {
	if id == "" {
		return nil, errors.New("task ID is required")
	}

	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if task.Status == domain.StatusPending && !task.DueDate.IsZero() && task.DueDate.Before(time.Now()) {
		task.Status = domain.StatusOverdue
		uc.taskRepo.Update(id, task)
	}

	return task, nil
}

func (uc *TaskUseCase) GetAllTasks() ([]*domain.Task, error) {
	tasks, err := uc.taskRepo.GetAll()
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if task.Status == domain.StatusPending && !task.DueDate.IsZero() && task.DueDate.Before(time.Now()) {
			task.Status = domain.StatusOverdue
			uc.taskRepo.Update(task.ID, task)
		}
	}

	return tasks, nil
}

func (uc *TaskUseCase) UpdateTask(id string, updatedTask *domain.Task) (*domain.Task, error) {
	if id == "" {
		return nil, errors.New("task ID is required")
	}

	existingTask, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if updatedTask.Title == "" {
		return nil, errors.New("task title is required")
	}

	if updatedTask.Description == "" {
		return nil, errors.New("task description is required")
	}

	existingTask.Title = updatedTask.Title
	existingTask.Description = updatedTask.Description
	existingTask.DueDate = updatedTask.DueDate
	existingTask.Status = updatedTask.Status
	existingTask.UpdatedAt = time.Now()

	err = uc.taskRepo.Update(id, existingTask)
	if err != nil {
		return nil, err
	}

	return existingTask, nil
}

func (uc *TaskUseCase) DeleteTask(id string) error {
	if id == "" {
		return errors.New("task ID is required")
	}

	_, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return err
	}

	return uc.taskRepo.Delete(id)
} 