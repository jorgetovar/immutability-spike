package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Book struct {
	Author string
	Title  string
	Pages  int
}

func (book Book) GetPagesWithCode() string {
	var n = rand.Intn(200)
	book.Pages -= n
	return "The " + strings.ToUpper(book.Title) + " book has " + strconv.Itoa(book.Pages) + " pages"
}

func main() {

	var book = Book{Author: "Rich Hickey", Title: "Clojure for the Brave & True", Pages: 232}
	fmt.Printf("1) %+v\n", book)
	changeAuthor(&book, "Daniel Higginbotham")
	fmt.Printf("2) %+v\n", book)
	fmt.Println(book.GetPagesWithCode())
	fmt.Printf("3) %+v\n", book)

}

func changeAuthor(book *Book, author string) {
	book.Author = author
}
