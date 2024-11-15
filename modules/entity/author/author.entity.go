package author

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	*gorm.Model `json:"-"`
	ID          uint64    `gorm:"primarykey" json:"AuthorId"`
	Name        string    `validate:"required"`
	Birthdate   time.Time `json:"-"`
}
