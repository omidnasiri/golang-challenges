package library

import (
	"fmt"
	"strings"
)

type Library struct {
	Books []Book
	Cap   int
}

type Book struct {
	Name           string
	BorrowedPerson string
}

func NewLibrary(capacity int) *Library {
	return &Library{
		Books: make([]Book, 0),
		Cap:   capacity,
	}
}

func (library *Library) AddBook(name string) string {
	index := library.findBookIndexByName(name)
	if index != -1 {
		return "The book is already in the library"
	}

	if len(library.Books) == library.Cap {
		return "Not enough capacity"
	}

	library.Books = append(library.Books, Book{Name: strings.ToLower(name), BorrowedPerson: ""})

	return "OK"
}

func (library *Library) BorrowBook(bookName, personName string) string {
	index := library.findBookIndexByName(bookName)
	if index == -1 {
		return "The book is not defined in the library"
	}

	if library.Books[index].BorrowedPerson != "" {
		return fmt.Sprintf("The book is already borrowed by %s", library.Books[index].BorrowedPerson)
	}

	library.Books[index].BorrowedPerson = personName
	fmt.Println(library)

	return "OK"
}

func (library *Library) ReturnBook(bookName string) string {
	index := library.findBookIndexByName(bookName)
	if index == -1 {
		return "The book is not defined in the library"
	}

	if library.Books[index].BorrowedPerson == "" {
		return "The book has not been borrowed"
	}

	library.Books[index].BorrowedPerson = ""

	return "OK"
}

func (library *Library) findBookIndexByName(name string) int {
	for i, v := range library.Books {
		if strings.ToLower(name) == v.Name {
			return i
		}
	}
	return -1
}
