package controller

import (
	"GinFrameworkPractice/entity"
	"GinFrameworkPractice/service"
	"github.com/gin-gonic/gin"
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
	return bookController.bookService.FindAll()
}

func (bookController *BookController) FindBookById(id int) entity.Book {
	return bookController.bookService.FindBookById(id)
}

func (bookController *BookController) AddBook(context *gin.Context) entity.Book {
	var book entity.Book
	return bookController.bookService.AddBook(book)

}
