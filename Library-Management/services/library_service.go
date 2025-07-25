package services

import (
	"errors"
	"Library-Management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books   map[int]models.Book
	Members map[int]models.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}
	if book.Status == models.Borrowed {
		return errors.New("book already borrowed")
	}

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	// Update book and member
	book.Status = models.Borrowed
	l.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)
	l.Members[memberID] = member

	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	book, exists := l.Books[bookID]
	if !exists {
		return errors.New("book not found")
	}

	member, exists := l.Members[memberID]
	if !exists {
		return errors.New("member not found")
	}

	// Remove book from members borrowed list
	found := false
	newList := []models.Book{}
	for _, b := range member.BorrowedBooks {
		if b.ID == bookID {
			found = true
			continue
		}
		newList = append(newList, b)
	}
	if !found {
		return errors.New("book not borrowed by this member")
	}
	member.BorrowedBooks = newList
	l.Members[memberID] = member

	// Mark a book as available
	book.Status = models.Available
	l.Books[bookID] = book

	return nil
}

func (l *Library) ListAvailableBooks() []models.Book {
	available := []models.Book{}
	for _, book := range l.Books {
		if book.Status == models.Available {
			available = append(available, book)
		}
	}
	return available
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, exists := l.Members[memberID]
	if !exists {
		return []models.Book{}
	}
	return member.BorrowedBooks
}
