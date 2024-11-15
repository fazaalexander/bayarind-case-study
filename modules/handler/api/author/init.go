package author

import ac "github.com/fazaalexander/bayarind-case-study.git/modules/usecase/author"

type AuthorHandler struct {
	authorHandler ac.AuthorUseCase
}

func New(authorUseCase ac.AuthorUseCase) *AuthorHandler {
	return &AuthorHandler{
		authorUseCase,
	}
}
