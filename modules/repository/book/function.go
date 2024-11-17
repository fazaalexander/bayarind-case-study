package book

import (
	ae "github.com/fazaalexander/bayarind-case-study.git/modules/entity/author"
	be "github.com/fazaalexander/bayarind-case-study.git/modules/entity/book"
	"github.com/labstack/echo/v4"
)

func (br *bookRepo) GetAllBooks(books *[]be.Book, offset, pageSize int) ([]be.Book, int64, error) {
	var count int64
	if err := br.db.Model(&be.Book{}).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := br.db.Preload("Author").
		Offset(offset).Limit(pageSize).Order("created_at ASC").
		Find(&books).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return *books, count, nil
}

func (br *bookRepo) GetBookById(bookId uint64, book *be.Book) (*be.Book, error) {
	if err := br.db.Preload("Author").Where("id = ?", bookId).First(&book).Error; err != nil {
		return nil, echo.NewHTTPError(404, err)
	}

	return book, nil
}

func (br *bookRepo) CreateBook(book *be.Book) error {
	if err := br.db.Where("id = ?", book.AuthorId).First(&ae.Author{}).Error; err != nil {
		return echo.NewHTTPError(404, err)
	}

	if err := br.db.Create(&book).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (br *bookRepo) UpdateBook(bookId uint64, book *be.Book) error {
	if err := br.db.Where("id = ?", book.AuthorId).First(&ae.Author{}).Error; err != nil {
		return echo.NewHTTPError(404, err)
	}

	if err := br.db.Model(&be.Book{}).Where("id = ?", bookId).Updates(be.Book{Title: book.Title, Isbn: book.Isbn, AuthorId: book.AuthorId}).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (br *bookRepo) DeleteBook(bookId uint64, book *be.Book) error {
	if err := br.db.Where("id = ?", bookId).Delete(&book).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}
