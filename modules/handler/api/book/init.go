package book

import bc "github.com/fazaalexander/bayarind-case-study.git/modules/usecase/book"

type BookHandler struct {
	bookUseCase bc.BookUseCase
}

func New(bookUseCase bc.BookUseCase) *BookHandler {
	return &BookHandler{
		bookUseCase,
	}
}
