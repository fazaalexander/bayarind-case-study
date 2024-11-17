package author

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	*gorm.Model `json:"-"`
	ID          uint64    `gorm:"primarykey" json:"AuthorId"`
	Name        string    `validate:"required"`
	Birthdate   time.Time `validate:"required"`
}

type AuthorRequest struct {
	Name      string `json:"Name" form:"Name" validate:"required"`
	Birthdate string `json:"Birthday" form:"Birthday" validate:"required"`
}

type AuthorResponse struct {
	AuthorID  uint64 `json:"AuthorId" form:"AuthorId"`
	Name      string `json:"Name" form:"Name"`
	Birthdate string `json:"Birthdate" form:"Birthdate"`
}
