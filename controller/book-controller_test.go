package controller

import (
	"GinFrameworkPractice/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookController_FindAll(t *testing.T) {
	bookController := New()
	assert.Equal(t, repository.DefaultBooks(), bookController.FindAll())
}
