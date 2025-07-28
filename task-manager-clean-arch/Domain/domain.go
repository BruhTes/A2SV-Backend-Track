package domain

import (
	"time"
)

type TaskStatus string

const (
	StatusPending   TaskStatus = "Pending"
	StatusCompleted TaskStatus = "Completed"
	StatusOverdue   TaskStatus = "Overdue"
)

type Task struct {
	ID          string     `json:"id" bson:"_id,omitempty"`
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"`
	DueDate     time.Time  `json:"due_date" bson:"due_date"`
	Status      TaskStatus `json:"status" bson:"status"`
	CreatedAt   time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" bson:"updated_at"`
}

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

type TaskRepository interface {
	Create(task *Task) error
	GetByID(id string) (*Task, error)
	GetAll() ([]*Task, error)
	Update(id string, task *Task) error
	Delete(id string) error
}

type UserRepository interface {
	Create(user *User) error
	GetByUsername(username string) (*User, error)
	GetByID(id string) (*User, error)
	ExistsByUsername(username string) (bool, error)
	GetUserCount() (int64, error)
}

type JWTService interface {
	GenerateToken(userID, username, role string) (string, error)
	ValidateToken(tokenString string) (*JWTClaims, error)
}

type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hash string) bool
}

type JWTClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
} 