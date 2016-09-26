package mocks

import "github.com/tixu/Auth/users"

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

func (*UserService) GetUser(name string) (user users.User, ok bool) {

	user, ok = DB[name]
	return user, ok
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
