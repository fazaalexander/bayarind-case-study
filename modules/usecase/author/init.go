package author

import (
	ae "github.com/fazaalexander/bayarind-case-study.git/modules/entity/author"
	ar "github.com/fazaalexander/bayarind-case-study.git/modules/repository/author"
)

type AuthorUseCase interface {
	GetAllAuthor(authors *[]ae.Author, offset, pageSize int) ([]ae.Author, int64, error)
	GetAuthorById(authorId uint64, author *ae.Author) (*ae.Author, error)
	CreateAuthor(authorRequest *ae.AuthorRequest) error
	UpdateAuthor(authorId uint64, author *ae.AuthorRequest) error
	DeleteAuthor(authorId uint64, author *ae.Author) error
}

type authorUseCase struct {
	authorRepo ar.AuthorRepo
}

func New(authorRepo ar.AuthorRepo) *authorUseCase {
	return &authorUseCase{
		authorRepo,
	}
}
