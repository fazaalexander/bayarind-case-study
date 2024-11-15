package book

import br "github.com/fazaalexander/bayarind-case-study.git/modules/repository/book"

type BookUseCase interface {
}

type bookUseCase struct {
	bookRepo br.BookRepo
}

func New(bookRepo br.BookRepo) *bookUseCase {
	return &bookUseCase{
		bookRepo,
	}
}
