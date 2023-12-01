package service

import (
	"GinFrameworkPractice/entity"
)

type BookService struct {
	books []entity.Book
}

func New() *BookService {
	return &BookService{}
}
