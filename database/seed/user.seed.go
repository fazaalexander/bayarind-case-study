package seed

import (
	"github.com/fazaalexander/bayarind-case-study.git/helper/password"
	ue "github.com/fazaalexander/bayarind-case-study.git/modules/entity/user"
)

func CreateUser() []*ue.User {
	hashPasswordUser1, _ := password.HashPassword("user1")
	hashPasswordUser2, _ := password.HashPassword("user2")
	hashPasswordUser3, _ := password.HashPassword("user3")

	users := []*ue.User{
		{
			ID:       1,
			Username: "username1",
			Password: string(hashPasswordUser1),
		},
		{
			ID:       2,
			Username: "username2",
			Password: string(hashPasswordUser2),
		},
		{
			ID:       3,
			Username: "username3",
			Password: string(hashPasswordUser3),
		},
	}

	return users
}
