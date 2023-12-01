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
