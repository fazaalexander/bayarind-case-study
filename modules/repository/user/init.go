package user

import (
	ue "github.com/fazaalexander/bayarind-case-study.git/modules/entity/user"
	"gorm.io/gorm"
)

type UserRepo interface {
	CreateUser(user *ue.RegisterRequest) error
	Login(username string) (*ue.LoginResponse, string, error)
}

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserRepo {
	return &userRepo{
		db,
	}
}
