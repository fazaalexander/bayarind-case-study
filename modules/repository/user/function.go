package user

import (
	"errors"

	ue "github.com/fazaalexander/bayarind-case-study.git/modules/entity/user"
)

func (ur *userRepo) CreateUser(user *ue.RegisterRequest) error {
	existingUser := ue.User{}
	if err := ur.db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		return errors.New("username already exists")
	}

	userTable := ue.User{
		Username: user.Username,
		Password: user.Password,
	}

	if err := ur.db.Create(&userTable).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepo) Login(username string) (*ue.LoginResponse, string, error) {
	user := &ue.User{}
	if err := ur.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, "", errors.New("Record not found")
	}

	response := &ue.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
	}

	return response, user.Password, nil
}
