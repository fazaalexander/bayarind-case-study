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

type BookRequest struct {
	Title    string `json:"Title" form:"Title" validate:"required"`
	Isbn     string `json:"Isbn" form:"Title" validate:"required"`
	AuthorId uint64 `json:"AuthorId" form:"AuthorId" validate:"required"`
}

type BookResponse struct {
	BookId uint64            `json:"BookId"`
	Title  string            `json:"Title"`
	Isbn   string            `json:"Isbn"`
	Author ae.AuthorResponse `json:"Author"`
}
