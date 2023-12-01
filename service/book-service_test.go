package service

import (
	"GinFrameworkPractice/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

var bookService = New()

func TestNew(t *testing.T) {
	assert.Equal(t, repository.DefaultBooks(), bookService.books)
}

func TestBookService_FindAll(t *testing.T) {
	assert.Equal(t, repository.DefaultBooks(), bookService.FindAll())
}
