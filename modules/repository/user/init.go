package user

import "gorm.io/gorm"

type UserRepo interface {
}

type userRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) UserRepo {
	return &userRepo{
		db,
	}
}
