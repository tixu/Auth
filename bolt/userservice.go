package bolt

import (
	"log"

	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
	"github.com/tixu/Auth/bolt/internal"
	"github.com/tixu/Auth/users"
)

// UserService represents a service for managing dials.

type UserService struct {
	session *Session
}

func (s *UserService) Ping() error {
	tx, err := s.session.db.Begin(false)
	if err != nil {
		return errors.Wrap(err, "unable to get a session while ping the database")
	}
	log.Println(tx.Size())
	defer s.session.db.Close()
	return nil
}

func (s *UserService) DeleteUser(name string) error {
	// Start writable transaction.
	tx, err := s.session.db.Begin(true)
	if err != nil {
		return errors.Wrap(err, "unable to get a session while deleting a user")
	}
	defer tx.Commit()

	b := tx.Bucket([]byte(userBucket))
	if b.Delete([]byte(name)) != nil {
		tx.Rollback()
		return errors.Wrap(err, "error while deleting message")
	}

	return nil

}

// GetUser returns a user by UserName.
func (s *UserService) GetUser(name string) (*users.User, error) {
	// Start read-only transaction.
	tx, err := s.session.db.Begin(false)
	if err != nil {
		return &users.User{}, err
	}
	defer tx.Rollback()

	// Find and unmarshal user.
	var u users.User

	if v := tx.Bucket([]byte(userBucket)).Get([]byte(name)); v == nil {
		err := errors.New("No Data Found")
		return &u, err
	} else if err := internal.UnmarshalUser(v, &u); err != nil {

		return &u, errors.Wrap(err, "unable to unmarshall retrieve information")
	}
	return &u, nil
}

//ListAll list all the users store in the bolt store
func (s *UserService) ListAll() (users.Users, error) {
	tx, err := s.session.db.Begin(false)
	defer tx.Commit()
	if err != nil {
		return nil, errors.Wrap(err, "unable to get a session while listing all the users")
	}
	// Find and unmarshal user.
	var u users.User
	var us users.Users = make(map[string]users.User)

	tx.DB().View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(userBucket))
		b.ForEach(func(k, v []byte) error {
			err := internal.UnmarshalUser(v, &u)
			if err != nil {
				return errors.Wrap(err, "error will unmarshalling db")
			}
			s := string(k)
			us[s] = u
			return nil
		})
		return nil
	})
	return us, nil
}

//AddUser adds an user to bolt store
func (s *UserService) AddUser(user *users.User) error {
	// Start writable transaction.
	tx, err := s.session.db.Begin(true)
	if err != nil {
		return errors.Wrap(err, "unable to get a session while creating a user")
	}
	defer tx.Commit()

	b := tx.Bucket([]byte(userBucket))
	seq, _ := b.NextSequence()
	//assigning a sequence
	user.ID = seq

	if um, err := internal.MarshalUser(user); err != nil {
		return errors.Wrap(err, "error while marshalling user before storage")
	} else if (b.Put([]byte(user.Name), um)) != nil {
		tx.Rollback()
		return errors.Wrap(err, "error while storing message")
	}

	return nil
}
