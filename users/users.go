package users

type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"username"`
	PasswordHash string `json:"passwordhash"`
	Email        string `json:"email"`
	Role         string `json:"role"`
}

type Users map[string]User

type UserService interface {
	GetUser(name string) (User, error)
}
