package service

import (
	"GinFrameworkPractice/entity"
	"GinFrameworkPractice/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var bookService = New()
var book = entity.Book{
	Id:     1,
	Title:  "The Bee Sting",
	Author: "Paul Murray",
	Price:  "700",
}

func TestNew(t *testing.T) {
	assert.Equal(t, repository.DefaultBooks(), bookService.books)
}

func TestBookService_FindAll(t *testing.T) {
	assert.Equal(t, repository.DefaultBooks(), bookService.FindAll())
}

func TestBookService_FindBookById(t *testing.T) {
	assert.Equal(t, book, bookService.FindBookById(1))
}

func TestAddBook(t *testing.T) {
	assert.Equal(t, book, bookService.AddBook(book))
}

func TestBookService_RemoveBook(t *testing.T) {
	assert.Equal(t, book, bookService.RemoveBook(1))
}
