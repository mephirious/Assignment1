package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/NursultanNurgaliyev/Assignment1/Library"
)

func main() {
	Library := NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nLibrary's book collection")
		fmt.Println("Options:")
		fmt.Println("1. Add – will add a book to a map")
		fmt.Println("2. Borrow – will change IsBorrowed state from false to true only.")
		fmt.Println("3. Return – will change IsBorrowed state from true to false only.")
		fmt.Println("4. List – will show all available books (not borrowed)")
		fmt.Println("5. Exit – finish program execution.")
		fmt.Print("\nInput number: ")

		scanner.Scan()
		input := scanner.Text()

		switch input {
		case "1":
			fmt.Print("Enter book title: ")
			scanner.Scan()
			title := scanner.Text()

			fmt.Print("Enter book author: ")
			scanner.Scan()
			author := scanner.Text()

			Library.AddBook(*Library.NewBook(title, author))
			fmt.Println("New book added")
		case "2":
			fmt.Print("Input id of the book to borrow: ")
			scanner.Scan()
			id := scanner.Text()
			Library.BorrowBook(id)
		case "3":
			fmt.Print("Input id of the book to return: ")
			scanner.Scan()
			id := scanner.Text()
			Library.ReturnBook(id)
		case "4":
			Library.ListBooks()
		case "5":
			fmt.Println("Exiting program. Goodbye!")
			return
		default:
			fmt.Println("Invalid input. Please try again.")
		}
	}
}
