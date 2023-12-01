package controller

import (
	"GinFrameworkPractice/entity"
	"GinFrameworkPractice/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var bookController *BookController = New()

func TestBookController_FindAll(t *testing.T) {
	assert.Equal(t, repository.DefaultBooks(), bookController.FindAll())
}

func TestBookController_FindBookById(t *testing.T) {
	book := entity.Book{
		Id:     1,
		Title:  "The Bee Sting",
		Author: "Paul Murray",
		Price:  "700",
	}
	assert.Equal(t, book, bookController.FindBookById(1))
}
