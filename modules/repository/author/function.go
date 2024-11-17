package author

import (
	ae "github.com/fazaalexander/bayarind-case-study.git/modules/entity/author"
	"github.com/labstack/echo/v4"
)

func (ar *authorRepo) GetAllAuthor(authors *[]ae.Author, offset, pageSize int) ([]ae.Author, int64, error) {
	var count int64
	if err := ar.db.Model(&ae.Author{}).Count(&count).Error; err != nil {
		return nil, 0, echo.NewHTTPError(500, err)
	}

	if err := ar.db.Limit(pageSize).Order("created_at ASC").Find(&authors).Error; err != nil {
		return nil, 0, echo.NewHTTPError(404, err)
	}

	return *authors, count, nil
}

func (ar *authorRepo) GetAuthorById(authorId uint64, author *ae.Author) (*ae.Author, error) {
	if err := ar.db.Where("id = ?", authorId).First(&author).Error; err != nil {
		return nil, echo.NewHTTPError(404, err)
	}

	return author, nil
}

func (ar *authorRepo) CreateAuthor(author *ae.Author) error {
	if err := ar.db.Create(&author).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (ar *authorRepo) UpdateAuthor(authorId uint64, author *ae.Author) error {
	if err := ar.db.Model(&ae.Author{}).Where("id = ?", authorId).Updates(ae.Author{Name: author.Name, Birthdate: author.Birthdate}).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}

func (ar *authorRepo) DeleteAuthor(authorId uint64, author *ae.Author) error {
	if err := ar.db.Where("id = ?", authorId).Delete(&author).Error; err != nil {
		return echo.NewHTTPError(500, err)
	}

	return nil
}
