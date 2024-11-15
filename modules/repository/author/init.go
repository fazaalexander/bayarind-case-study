package author

import "gorm.io/gorm"

type AuthorRepo interface {
}

type authorRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) AuthorRepo {
	return &authorRepo{
		db,
	}
}
