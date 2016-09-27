package bolt

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

// Session represents an authenticable connection to the database.
type Session struct {
	db  *bolt.DB
	now time.Time
	// Services
	userService UserService
}

// newSession returns a new instance of Session attached to db.
func newSession(db *bolt.DB) *Session {
	s := &Session{db: db}
	s.userService.session = s
	return s
}

// UserService returns a user service associated with this session.
func (s *Session) GetUserService() UserService { return s.userService }

// itob returns an 8-byte big-endian encoded byte slice of v.
//
// This function is typically used for encoding integer IDs to byte slices
// so that they can be used as BoltDB keys.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
