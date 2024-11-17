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

type RegisterRequest struct {
	Username string `json:"Username" form:"Username" validate:"required"`
	Password string `json:"Password" form:"Password" validate:"required,min=8"`
}

type LoginRequest struct {
	Username string `json:"Username" form:"Username" validate:"required"`
	Password string `json:"Password" form:"Password" validate:"required"`
}

type LoginResponse struct {
	ID       uint64 `json:"Id" form:"Id"`
	Username string `json:"Username" form:"Username"`
	Token    string `json:"Token" form:"Token"`
}
