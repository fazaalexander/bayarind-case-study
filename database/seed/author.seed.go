package seed

import (
	"time"

	ae "github.com/fazaalexander/bayarind-case-study.git/modules/entity/author"
)

func CreateAuthor() []*ae.Author {
	birthdate1, _ := time.Parse("2006-01-02", "1975-07-03")
	birthdate2, _ := time.Parse("2006-01-02", "1970-01-01")

	authors := []*ae.Author{
		{
			ID:        1,
			Name:      "Matt Haig",
			Birthdate: birthdate1,
		},
		{
			ID:        2,
			Name:      "Henry Manampiring",
			Birthdate: birthdate2,
		},
	}

	return authors
}
