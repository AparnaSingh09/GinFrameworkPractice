package service

import (
	"GinFrameworkPractice/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	bookService := New()
	assert.Equal(t, repository.DefaultBooks(), bookService.books)
}
