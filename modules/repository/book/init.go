package book

import "gorm.io/gorm"

type BookRepo interface {
}

type bookRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) BookRepo {
	return &bookRepo{
		db,
	}
}
