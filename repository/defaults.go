package repository

import "GinFrameworkPractice/entity"

func DefaultBooks() []entity.Book {
	return []entity.Book{
		{
			Title:  "The Bee Sting",
			Author: "Paul Murray",
			Price:  "700",
		},
		{
			Title:  "North Woods",
			Author: "Daniel Mason",
			Price:  "900",
		},
	}

}
