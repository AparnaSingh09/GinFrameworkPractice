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

func (service *BookService) RemoveBook(id int) entity.Book {
	book := service.FindBookById(id)
	return removeBook(service, book)
	//TODO if id not found
}

func removeBook(service *BookService, book entity.Book) entity.Book {
	var newBooks []entity.Book
	for i, j := range service.books {
		if j.Id == book.Id {
			newBooks = service.books[0:i]
			newBooks = append(service.books[i+1 : len(service.books)])
		}
	}
	service.books = newBooks
	return book
}
