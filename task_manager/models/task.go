package models

import "time"

type TaskStatus string

const (
	StatusPending   TaskStatus = "Pending"
	StatusCompleted TaskStatus = "Completed"
	StatusOverdue   TaskStatus = "Overdue"
)

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     time.Time  `json:"due_date"`
	Status      TaskStatus `json:"status"`
}
