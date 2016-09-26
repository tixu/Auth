package users

type User struct {
	ID       string `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type UserService interface {
	GetUser() User
}
