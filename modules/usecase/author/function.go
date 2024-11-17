package author

import (
	"time"

	vld "github.com/fazaalexander/bayarind-case-study.git/helper/validator"
	ae "github.com/fazaalexander/bayarind-case-study.git/modules/entity/author"
)

func (ac *authorUseCase) GetAllAuthor(authors *[]ae.Author, offset, pageSize int) ([]ae.Author, int64, error) {
	return ac.authorRepo.GetAllAuthor(authors, offset, pageSize)
}

func (ac *authorUseCase) GetAuthorById(authorId uint64, author *ae.Author) (*ae.Author, error) {
	return ac.authorRepo.GetAuthorById(authorId, author)
}

func (ac *authorUseCase) CreateAuthor(authorRequest *ae.AuthorRequest) error {
	if err := vld.Validation(authorRequest); err != nil {
		return err
	}

	birthdate, err := time.Parse("02-01-2006", authorRequest.Birthdate)
	if err != nil {
		return err
	}

	author := &ae.Author{
		Name:      authorRequest.Name,
		Birthdate: birthdate,
	}

	return ac.authorRepo.CreateAuthor(author)
}

func (ac *authorUseCase) UpdateAuthor(authorId uint64, authorRequest *ae.AuthorRequest) error {
	if err := vld.Validation(authorRequest); err != nil {
		return err
	}

	birthdate, err := time.Parse("02-01-2006", authorRequest.Birthdate)
	if err != nil {
		return err
	}

	author := &ae.Author{
		Name:      authorRequest.Name,
		Birthdate: birthdate,
	}

	return ac.authorRepo.UpdateAuthor(authorId, author)
}

func (ac *authorUseCase) DeleteAuthor(authorId uint64, author *ae.Author) error {
	return ac.authorRepo.DeleteAuthor(authorId, author)
}
