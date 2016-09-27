package mocks

import (
	"errors"

	"github.com/tixu/Auth/users"
)

var DB = users.Users{
	"user": users.User{
		Name: "user",
		// bcrypt has for "password"
		PasswordHash: "$2a$10$KgFhp4HAaBCRAYbFp5XYUOKrbO90yrpUQte4eyafk4Tu6mnZcNWiK",
		Email:        "user@example.com",
		Role:         "wtfd",
	},
}

type UserService struct {
	DB users.Users
}

func (*UserService) GetUser(name string) (user users.User, err error) {

	user, ok := DB[name]
	if ok == false {
		return user, errors.New("not found")
	}
	return user, nil
}

func GetUserMockService() *UserService {
	return &UserService{
		DB: users.Users{
			"user": users.User{
				Name: "user",
				// bcrypt has for "password"
				PasswordHash: "$2a$10$KgFhp4HAaBCRAYbFp5XYUOKrbO90yrpUQte4eyafk4Tu6mnZcNWiK",
				Email:        "user@example.com",
				Role:         "wtfd",
			},
		},
	}

}