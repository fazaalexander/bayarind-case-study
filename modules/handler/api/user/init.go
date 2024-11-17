package user

import (
	uc "github.com/fazaalexander/bayarind-case-study.git/modules/usecase/user"
)

type UserHandler struct {
	userUseCase uc.UserUseCase
}

func New(userUseCase uc.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase,
	}
}
