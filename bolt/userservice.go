package bolt

import "github.com/tixu/Auth/users"

// DialService represents a service for managing dials.

type UserService struct {
	session *Session
}

// GetUser returns a user by UserName.
func (s *UserService) GetUser(name string) (users.User, error) {
	// Start read-only transaction.
	tx, err := s.session.db.Begin(false)
	if err != nil {
		return users.User{}, err
	}
	defer tx.Rollback()

	// Find and unmarshal user.
	var u users.User

	if v := tx.Bucket([]byte(userBucket)).Get([]byte(name)); v == nil {

	}

	// else if err := internal.UnmarshalDial(v, &u); err != nil {
	//	return nil, err
	// }
	return u, nil
}
