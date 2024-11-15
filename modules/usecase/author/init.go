package author

import ap "github.com/fazaalexander/bayarind-case-study.git/modules/repository/author"

type AuthorUseCase interface {
}

type authorUseCase struct {
	authorRepo ap.AuthorRepo
}

func New(authorRepo ap.AuthorRepo) *authorUseCase {
	return &authorUseCase{
		authorRepo,
	}
}
