package users

// User is the information about the user
type User struct {
	ID           uint64 `json:"id"`
	Name         string `json:"username"`
	PasswordHash string `json:"passwordhash"`
	Email        string `json:"email"`
	Role         string `json:"role"`
}

// Users contains all the users sorted by their name
type Users map[string]User

// UserService is the operations that will be used to respond to the client.
type UserService interface {
	GetUser(name string) (*User, error)
}

// UserAdmin contains all the operations neeeded to manage the users.
type UserAdmin interface {
	AddUser(user *User) error
	DeleteUser(name string) error
	ListAll() (Users, error)
}
