package user

import (
	"errors"

	pw "github.com/fazaalexander/bayarind-case-study.git/helper/password"
	vld "github.com/fazaalexander/bayarind-case-study.git/helper/validator"
	"github.com/fazaalexander/bayarind-case-study.git/middleware/jwt"
	ue "github.com/fazaalexander/bayarind-case-study.git/modules/entity/user"
)

func (uc *userUseCase) Register(request *ue.RegisterRequest) error {
	if err := vld.Validation(request); err != nil {
		return err
	}

	hashedPassword, err := pw.HashPassword(request.Password)
	if err != nil {
		return err
	}

	request.Password = string(hashedPassword)

	err = uc.userRepo.CreateUser(request)
	if err != nil {
		return err
	}

	return nil
}

func (uc *userUseCase) Login(request *ue.LoginRequest) (interface{}, error) {
	if err := vld.Validation(request); err != nil {
		return nil, err
	}

	response, password, err := uc.userRepo.Login(request.Username)

	if err != nil {
		return nil, errors.New("wrong email or password")
	}

	if err = pw.VerifyPassword(password, request.Password); err != nil {
		return nil, errors.New("wrong Password")
	}

	token, err := jwt.CreateToken(response.ID, response.Username)
	if err != nil {
		return nil, err
	}

	response.Token = token

	return response, nil
}
