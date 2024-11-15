package book

import (
	ae "github.com/fazaalexander/bayarind-case-study.git/modules/entity/author"
	"gorm.io/gorm"
)

type Book struct {
	*gorm.Model `json:"-"`
	ID          uint64    `gorm:"primaryKey" json:"BookId"`
	Title       string    `validate:"required"`
	Isbn        string    `gorm:"unique" validate:"required"`
	AuthorId    uint64    `json:"-" validate:"required"`
	Author      ae.Author `gorm:"foreignKey:AuthorId" json:"-"`
}
