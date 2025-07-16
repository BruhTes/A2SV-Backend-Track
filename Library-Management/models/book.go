package models

type BookStatus string

const (
	Available = BookStatus("Available")
	Borrowed  = BookStatus("Borrowed")
)

type Book struct {
	ID     int
	Title  string
	Author string
	Status BookStatus
}
