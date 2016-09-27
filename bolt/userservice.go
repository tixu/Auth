package bolt

import (
	"github.com/pkg/errors"
	"github.com/tixu/Auth/bolt/internal"
	"github.com/tixu/Auth/users"
)

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
		err := errors.New("No Data Found")
		return u, err
	} else if err := internal.UnmarshalUser(v, &u); err != nil {

		return u, errors.Wrap(err, "unable to unmarshall retrieve information")
	}
	return u, nil
}
