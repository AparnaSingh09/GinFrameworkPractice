package repository

import "GinFrameworkPractice/entity"

func DefaultBooks() []entity.Book {
	return []entity.Book{
		{
			Id:     1,
			Title:  "The Bee Sting",
			Author: "Paul Murray",
			Price:  "700",
		},
		{
			Id:     2,
			Title:  "North Woods",
			Author: "Daniel Mason",
			Price:  "900",
		},
	}

}
