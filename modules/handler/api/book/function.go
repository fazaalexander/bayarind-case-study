package book

import (
	"math"
	"net/http"
	"strconv"

	ae "github.com/fazaalexander/bayarind-case-study.git/modules/entity/author"
	be "github.com/fazaalexander/bayarind-case-study.git/modules/entity/book"
	"github.com/labstack/echo/v4"
)

func (bh *BookHandler) GetAllBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		var books []be.Book

		pageParam := c.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		books, total, err := bh.bookUseCase.GetAllBooks(&books, offset, pageSize)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusNotFound,
			})
		}

		if len(books) == 0 {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": err,
				"Status":  http.StatusNotFound,
			})
		}

		totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
		if page > totalPages {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": err,
				"Status":  http.StatusNotFound,
			})
		}

		var bookResponses []be.BookResponse
		for _, book := range books {
			bookResponse := be.BookResponse{
				BookId: book.ID,
				Title:  book.Title,
				Isbn:   book.Isbn,
				Author: ae.AuthorResponse{
					AuthorID:  book.AuthorId,
					Name:      book.Author.Name,
					Birthdate: book.Author.Birthdate.Format("02-01-2006"),
				},
			}

			bookResponses = append(bookResponses, bookResponse)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Books":     bookResponses,
			"Page":      page,
			"TotalPage": totalPages,
			"Status":    http.StatusOK,
		})
	}
}

func (bh *BookHandler) GetBookById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var book *be.Book
		bookID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

		book, err := bh.bookUseCase.GetBookById(bookID, book)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusNotFound,
			})
		}

		bookResponse := be.BookResponse{
			BookId: book.ID,
			Title:  book.Title,
			Isbn:   book.Isbn,
			Author: ae.AuthorResponse{
				AuthorID:  book.AuthorId,
				Name:      book.Author.Name,
				Birthdate: book.Author.Birthdate.Format("02-01-2006"),
			},
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Author": bookResponse,
			"Status": http.StatusOK,
		})
	}
}

func (bh *BookHandler) CreateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var book be.BookRequest
		if err := c.Bind(&book); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		if err := bh.bookUseCase.CreateBook(&book); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Message": "New Book Created",
			"Status":  http.StatusOK,
		})
	}
}

func (bh *BookHandler) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var book be.BookRequest
		bookID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		if err := c.Bind(&book); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		err := bh.bookUseCase.UpdateBook(bookID, &book)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Message": "Book Updated Successfully",
			"Status":  http.StatusOK,
		})
	}
}

func (bh *BookHandler) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		var book *be.Book
		bookID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		if err := c.Bind(&book); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		err := bh.bookUseCase.DeleteBook(bookID, book)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Book successfully deleted",
			"Status":  http.StatusOK,
		})
	}
}
