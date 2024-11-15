package user

import (
	ur "github.com/fazaalexander/bayarind-case-study.git/modules/repository/user"
)

type UserUseCase interface {
}

type userUseCase struct {
	userRepo ur.UserRepo
}

func New(userRepo ur.UserRepo) *userUseCase {
	return &userUseCase{
		userRepo,
	}
}
