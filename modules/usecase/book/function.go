package book

import (
	vld "github.com/fazaalexander/bayarind-case-study.git/helper/validator"
	be "github.com/fazaalexander/bayarind-case-study.git/modules/entity/book"
)

func (bc *bookUseCase) GetAllBooks(books *[]be.Book, offset, pageSize int) ([]be.Book, int64, error) {
	return bc.bookRepo.GetAllBooks(books, offset, pageSize)
}

func (bc *bookUseCase) GetBookById(bookId uint64, book *be.Book) (*be.Book, error) {
	return bc.bookRepo.GetBookById(bookId, book)
}

func (bc *bookUseCase) CreateBook(bookRequest *be.BookRequest) error {
	if err := vld.Validation(bookRequest); err != nil {
		return err
	}

	book := &be.Book{
		Title:    bookRequest.Title,
		Isbn:     bookRequest.Isbn,
		AuthorId: bookRequest.AuthorId,
	}
	return bc.bookRepo.CreateBook(book)
}

func (bc *bookUseCase) UpdateBook(bookId uint64, bookRequest *be.BookRequest) error {
	if err := vld.Validation(bookRequest); err != nil {
		return err
	}

	book := &be.Book{
		Title:    bookRequest.Title,
		Isbn:     bookRequest.Isbn,
		AuthorId: bookRequest.AuthorId,
	}
	return bc.bookRepo.UpdateBook(bookId, book)
}

func (bc *bookUseCase) DeleteBook(bookId uint64, book *be.Book) error {
	return bc.bookRepo.DeleteBook(bookId, book)
}
