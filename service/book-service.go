package service

import (
	"GinFrameworkPractice/entity"
	"GinFrameworkPractice/repository"
)

type BookService struct {
	books []entity.Book
}

func New() *BookService {
	return &BookService{repository.DefaultBooks()}
}

func (service *BookService) FindAll() []entity.Book {
	return service.books
}

func (service *BookService) FindBookById(id int) entity.Book {
	for _, book := range service.books {
		if book.Id == id {
			return book
		} else {
			//TODO
		}
	}
	return entity.Book{}
}

func (service *BookService) AddBook(book entity.Book) entity.Book {
	service.books = append(service.books, book)
	return book
}
