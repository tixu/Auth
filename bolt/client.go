package bolt

import (
	"time"

	"github.com/boltdb/bolt"
)

type Client struct {
	// Filename to the BoltDB database.
	Path string

	// Returns the current time.
	Now func() time.Time

	db *bolt.DB
}

func NewClient() *Client {
	return &Client{
		Now: time.Now,
	}
}

// Open opens and initializes the BoltDB database.
func (c *Client) Open() error {
	// Open database file.
	db, err := bolt.Open(c.Path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	c.db = db

	// Initialize top-level buckets.
	tx, err := c.db.Begin(true)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.CreateBucketIfNotExists([]byte(userBucket)); err != nil {
		return err
	}

	return tx.Commit()
}

// Close closes then underlying BoltDB database.
func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}

// Connect returns a new session to the BoltDB database.
func (c *Client) Connect() *Session {
	s := newSession(c.db)
	s.now = c.Now()
	return s
}
