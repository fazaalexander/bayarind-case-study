package user

import (
	ue "github.com/fazaalexander/bayarind-case-study.git/modules/entity/user"
	ur "github.com/fazaalexander/bayarind-case-study.git/modules/repository/user"
)

type UserUseCase interface {
	Login(request *ue.LoginRequest) (interface{}, error)
	Register(user *ue.RegisterRequest) error
}

type userUseCase struct {
	userRepo ur.UserRepo
}

func New(userRepo ur.UserRepo) *userUseCase {
	return &userUseCase{
		userRepo,
	}
}
