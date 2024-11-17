package author

import (
	ae "github.com/fazaalexander/bayarind-case-study.git/modules/entity/author"
	"gorm.io/gorm"
)

type AuthorRepo interface {
	GetAllAuthor(authors *[]ae.Author, offset, pageSize int) ([]ae.Author, int64, error)
	GetAuthorById(authorId uint64, author *ae.Author) (*ae.Author, error)
	CreateAuthor(author *ae.Author) error
	UpdateAuthor(authorId uint64, author *ae.Author) error
	DeleteAuthor(authorId uint64, author *ae.Author) error
}

type authorRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) AuthorRepo {
	return &authorRepo{
		db,
	}
}
