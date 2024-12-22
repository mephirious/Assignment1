package Library

import (
	"fmt"
	"strconv"
)

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	CollectionOfBooks map[string]Book
}

func (l *Library) NewBook(title, author string) *Book {
	idCounter := strconv.Itoa(len(l.CollectionOfBooks) + 1)
	return &Book{idCounter, title, author, false}
}

func NewLibrary() *Library {
	return &Library{make(map[string]Book)}
}

func (l *Library) AddBook(book Book) {
	l.CollectionOfBooks[book.ID] = book
}

func (l *Library) BorrowBook(id string) {
	if len(l.CollectionOfBooks) == 0 {
		fmt.Println("Library is empty")
		return
	}
	book, isExist := l.CollectionOfBooks[id]
	if !isExist {
		fmt.Println("Library does not have book with this ID")
		return
	}
	if book.IsBorrowed {
		fmt.Println("Book is already borrowed")
	}
	book.IsBorrowed = true
	l.CollectionOfBooks[id] = book
}

func (l *Library) ReturnBook(id string) {
	if len(l.CollectionOfBooks) == 0 {
		fmt.Println("Library is empty")
		return
	}
	book, isExist := l.CollectionOfBooks[id]
	if !isExist {
		fmt.Println("Library does not have book with this ID")
		return
	}
	if !book.IsBorrowed {
		fmt.Println("Book is already returned")
		return
	}
	book.IsBorrowed = false
	l.CollectionOfBooks[id] = book
}

func (l *Library) ListBooks() {
	if len(l.CollectionOfBooks) == 0 {
		fmt.Println("Library is empty")
		return
	}
	for _, value := range l.CollectionOfBooks {
		if !value.IsBorrowed {
			fmt.Printf("ID: %s\nTitle: %s\nAuthor: %s\nIsBorrowed: %t\n", value.ID, value.Title, value.Author, value.IsBorrowed)
		}
	}
}
