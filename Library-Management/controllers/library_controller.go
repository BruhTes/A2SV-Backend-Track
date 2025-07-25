package controllers

import (
	"bufio"
	"fmt"
	"Library-Management/models"
	"Library-Management/services"
	"os"
	"strconv"
	"strings"
)

type LibraryController struct {
	Library *services.Library
}

func NewLibraryController(library *services.Library) *LibraryController {
	return &LibraryController{Library: library}
}

func (lc *LibraryController) Start() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			lc.addBook(reader)
		case "2":
			lc.removeBook(reader)
		case "3":
			lc.borrowBook(reader)
		case "4":
			lc.returnBook(reader)
		case "5":
			lc.listAvailableBooks()
		case "6":
			lc.listBorrowedBooks(reader)
		case "7":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

func (lc *LibraryController) addBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	fmt.Print("Enter Title: ")
	title, _ := reader.ReadString('\n')

	fmt.Print("Enter Author: ")
	author, _ := reader.ReadString('\n')

	book := models.Book{
		ID:     id,
		Title:  strings.TrimSpace(title),
		Author: strings.TrimSpace(author),
		Status: models.Available,
	}

	lc.Library.AddBook(book)
	fmt.Println("Book added.")
}

func (lc *LibraryController) removeBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))
	lc.Library.RemoveBook(id)
	fmt.Println("Book removed.")
}

func (lc *LibraryController) borrowBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	// if member does not exist create them
	if _, exists := lc.Library.Members[memberID]; !exists {
		fmt.Print("Enter Member Name: ")
		name, _ := reader.ReadString('\n')
		member := models.Member{
			ID:            memberID,
			Name:          strings.TrimSpace(name),
			BorrowedBooks: []models.Book{},
		}
		lc.Library.Members[memberID] = member
	}

	if err := lc.Library.BorrowBook(bookID, memberID); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book borrowed.")
	}
}

func (lc *LibraryController) returnBook(reader *bufio.Reader) {
	fmt.Print("Enter Book ID: ")
	bookIDStr, _ := reader.ReadString('\n')
	bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	if err := lc.Library.ReturnBook(bookID, memberID); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Book returned.")
	}
}

func (lc *LibraryController) listAvailableBooks() {
	books := lc.Library.ListAvailableBooks()
	fmt.Println("Available Books:")
	for _, book := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", book.ID, book.Title, book.Author)
	}
}

func (lc *LibraryController) listBorrowedBooks(reader *bufio.Reader) {
	fmt.Print("Enter Member ID: ")
	memberIDStr, _ := reader.ReadString('\n')
	memberID, _ := strconv.Atoi(strings.TrimSpace(memberIDStr))

	books := lc.Library.ListBorrowedBooks(memberID)
	if len(books) == 0 {
		fmt.Println("No books borrowed or member not found.")
		return
	}

	fmt.Println("Borrowed Books:")
	for _, book := range books {
		fmt.Printf("ID: %d | Title: %s | Author: %s\n", book.ID, book.Title, book.Author)
	}
}
