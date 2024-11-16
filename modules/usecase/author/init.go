package author

import ar "github.com/fazaalexander/bayarind-case-study.git/modules/repository/author"

type AuthorUseCase interface {
}

type authorUseCase struct {
	authorRepo ar.AuthorRepo
}

func New(authorRepo ar.AuthorRepo) *authorUseCase {
	return &authorUseCase{
		authorRepo,
	}
}
