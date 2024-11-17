package book

import (
	be "github.com/fazaalexander/bayarind-case-study.git/modules/entity/book"
	br "github.com/fazaalexander/bayarind-case-study.git/modules/repository/book"
)

type BookUseCase interface {
	GetAllBooks(books *[]be.Book, offset, pageSize int) ([]be.Book, int64, error)
	GetBookById(bookId uint64, book *be.Book) (*be.Book, error)
	CreateBook(book *be.BookRequest) error
	UpdateBook(bookId uint64, bookRequest *be.BookRequest) error
	DeleteBook(bookId uint64, book *be.Book) error
}

type bookUseCase struct {
	bookRepo br.BookRepo
}

func New(bookRepo br.BookRepo) *bookUseCase {
	return &bookUseCase{
		bookRepo,
	}
}
