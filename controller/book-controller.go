package controller

import (
	"GinFrameworkPractice/entity"
	"GinFrameworkPractice/service"
)

type BookController struct {
	bookService *service.BookService
}

func New() *BookController {
	return &BookController{
		service.New(),
	}
}

func (bookController *BookController) FindAll() []entity.Book {
	return []entity.Book{}
}
