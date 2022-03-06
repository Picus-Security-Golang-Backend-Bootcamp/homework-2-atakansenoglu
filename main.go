// TAMAMLANMADI.

package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-2-atakansenoglu/helper"
)

func init() {
}

var pl = fmt.Println
var pf = fmt.Printf

var Args []string = os.Args[1:]

func main() {

	switch checkConsoleInput() {
	case "search":
		// search - to do
	case "list":
		// list - to do
	case "":
		fmt.Println("Kitap ismi yazmadınız!")
	case "err":
		fmt.Println("Yanlış veya eksik komut girdiniz!")
	}

	// Store created books of type 'Books'
	var createdBooks [5]Books

	// bookNames
	var bookNames []string
	bookNames = append(bookNames,
		"Angels and Demons",
		"The Da Vinci Code",
		"Kayıp Tanrılar Ülkesi",
		"Denemeler",
		"Hayvan Çiftliği")

	// authors of type 'author'
	var authors []author
	authors = append(authors,
		*NewAuthor("Dan", "Brown"),
		*NewAuthor("Ahmet", "Ümit"),
		*NewAuthor("Montaigne", ""),
		*NewAuthor("George", "Orwell"))

	// Create new books
	/* for i, _ := range authors {
		createdBooks = append(createdBooks, *NewBook(i, bookNames[i], "1111", authors[i], false))
	} */

	createdBooks[0] = *NewBook(1, bookNames[0], "1111", authors[0], false)
	createdBooks[1] = *NewBook(2, bookNames[1], "2222", authors[0], false)
	createdBooks[2] = *NewBook(3, bookNames[2], "3333", authors[1], false)
	createdBooks[3] = *NewBook(4, bookNames[3], "4444", authors[2], false)
	createdBooks[4] = *NewBook(5, bookNames[4], "5555", authors[3], false)

	pl(createdBooks[0])
	pl(createdBooks[1])
	pl(createdBooks[2])
	pl(createdBooks[3])
	pl(createdBooks[4])

}

// ---------- STRUCTS ----------

// Struct of 'author' type
type author struct {
	firstName string
	lastName  string
}

// Struct of 'Books' type
type Books struct {
	ID        int
	Name      string
	StockCode string
	ISBN      string
	Pages     int
	Price     float64
	Stock     int
	author
	IsDeleted bool
	Deletable
}

// ---------- CONSTRUCTORS ----------

// Constructor for Books struct
func NewBook(id int, name, stockCode string, author author, isDeleted bool) *Books {

	// Create new seed value to use for random
	rand.Seed(time.Now().UnixNano())

	book := &Books{
		ID:        id,
		StockCode: stockCode,
		ISBN:      helper.RandomString(13),
		Pages:     rand.Intn(1001-200) + 200,
		Stock:     rand.Intn(10000) + 1,
		Name:      name,
		author:    author,
		Price:     math.Floor(((rand.Float64()*(40-10))+10)*100) / 100,
		IsDeleted: isDeleted,
	}

	return book
}

// Constructor for authors struct
func NewAuthor(firstName, lastName string) *author {
	author := &author{
		firstName: firstName,
		lastName:  lastName,
	}

	return author
}

// ---------- INTERFACE ----------

type Deletable interface {
	Delete()
}

func (book Books) Delete() {
	pl("Delete fonksiyonu Books struct için çalıştı.")
}

// ------------------------------
func checkConsoleInput() string {

	if len(Args) <= 0 {
		fmt.Printf("search => arama işlemi için \n list => listeleme işlemi için\n")
		return ""

	} else if len(Args) > 0 && Args[0] == "search" {
		if len(Args) == 1 {
			return ""
		}
		return "search"

	} else if len(Args) > 0 && Args[0] == "list" {
		return "list"

	} else {
		return "err"
	}

}

// List all books
func listBooks(arr []string) {

	fmt.Println("Listed all books:")

	for i, v := range arr {
		fmt.Printf("%d: %s\n", i, v)
	}

}

// Contains
func contains(arr []string, str string) bool {
	for i, v := range arr {
		if v == str {
			fmt.Printf("Kitap: %s\nSıra: %d\n", v, i+1)
			return true
		}
	}
	return false
}
