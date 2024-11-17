package seed

import (
	be "github.com/fazaalexander/bayarind-case-study.git/modules/entity/book"
)

func CreateBook() []*be.Book {
	books := []*be.Book{
		{
			ID:       1,
			Title:    "The Midnight Library",
			Isbn:     "9786020649320",
			AuthorId: 1,
		},
		{
			ID:       2,
			Title:    "How To Stop Time",
			Isbn:     "9786020618791",
			AuthorId: 1,
		},
		{
			ID:       3,
			Title:    "Filosofi Teras",
			Isbn:     "9786233463034",
			AuthorId: 2,
		},
	}

	return books
}
