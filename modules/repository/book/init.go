package book

import (
	be "github.com/fazaalexander/bayarind-case-study.git/modules/entity/book"
	"gorm.io/gorm"
)

type BookRepo interface {
	GetAllBooks(books *[]be.Book, offset, pageSize int) ([]be.Book, int64, error)
	GetBookById(bookId uint64, book *be.Book) (*be.Book, error)
	CreateBook(book *be.Book) error
	UpdateBook(bookId uint64, book *be.Book) error
	DeleteBook(bookId uint64, book *be.Book) error
}

type bookRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) BookRepo {
	return &bookRepo{
		db,
	}
}
