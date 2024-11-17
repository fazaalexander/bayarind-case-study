package author

import (
	"math"
	"net/http"
	"strconv"

	ae "github.com/fazaalexander/bayarind-case-study.git/modules/entity/author"
	"github.com/labstack/echo/v4"
)

func (ah *AuthorHandler) GetAllAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var authors []ae.Author

		pageParam := c.QueryParam("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil || page < 1 {
			page = 1
		}

		pageSize := 10
		offset := (page - 1) * pageSize

		authors, total, err := ah.authorUseCase.GetAllAuthor(&authors, offset, pageSize)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusNotFound,
			})
		}

		if len(authors) == 0 {
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

		var authorResponses []ae.AuthorResponse
		for _, author := range authors {
			authorResponse := ae.AuthorResponse{
				AuthorID:  author.ID,
				Name:      author.Name,
				Birthdate: author.Birthdate.Format("02-01-2006"),
			}

			authorResponses = append(authorResponses, authorResponse)
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Authors":   authorResponses,
			"Page":      page,
			"TotalPage": totalPages,
			"Status":    http.StatusOK,
		})
	}
}

func (ah *AuthorHandler) GetAuthorById() echo.HandlerFunc {
	return func(c echo.Context) error {
		var author *ae.Author
		authorID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

		author, err := ah.authorUseCase.GetAuthorById(authorID, author)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusNotFound,
			})
		}

		authorResponse := ae.AuthorResponse{
			AuthorID:  author.ID,
			Name:      author.Name,
			Birthdate: author.Birthdate.Format("02-01-2006"),
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Author": authorResponse,
			"Status": http.StatusOK,
		})
	}
}

func (ah *AuthorHandler) CreateAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var author ae.AuthorRequest
		if err := c.Bind(&author); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		err := ah.authorUseCase.CreateAuthor(&author)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Message": "New Author Created",
			"Status":  http.StatusOK,
		})
	}
}

func (ah *AuthorHandler) UpdateAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var author ae.AuthorRequest
		authorID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		if err := c.Bind(&author); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		err := ah.authorUseCase.UpdateAuthor(authorID, &author)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"Message": "Author Updated Successfully",
			"Status":  http.StatusOK,
		})
	}
}

func (ah *AuthorHandler) DeleteAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		var author *ae.Author
		authorID, _ := strconv.ParseUint(c.Param("id"), 10, 64)

		author, err := ah.authorUseCase.GetAuthorById(authorID, author)
		if err != nil {
			return c.JSON(http.StatusNotFound, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusNotFound,
			})
		}

		err = ah.authorUseCase.DeleteAuthor(authorID, author)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Author successfully deleted",
			"Status":  http.StatusOK,
		})
	}
}
