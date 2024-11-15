package user

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model `json:"-"`
	ID          uint64 `gorm:"primarykey" json:"UserId"`
	Username    string `validate:"required"`
	Password    string `validate:"required,min=8"`
}
