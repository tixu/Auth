package bolt_test

import (
	"reflect"
	"testing"

	"github.com/tixu/Auth/users"
)

// Ensure dial can be created and retrieved.
func TestDialService_CreateUser(t *testing.T) {
	c := MustOpenClient()
	defer c.Close()
	s := c.Connect().GetUserService()

	user := users.User{
		Name:         "User",
		PasswordHash: "pwd",
		Email:        "user@email.com",
		Role:         "Role",
	}

	// Create new dial.
	if err := s.AddUser(&user); err != nil {
		t.Fatal(err)
	} else if user.ID != 1 {
		t.Fatalf("unexpected id: %d", user.ID)
	}

	// Retrieve dial and compare.
	other, err := s.GetUser("User")
	if err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(user, other) {
		t.Fatalf("unexpected user: %#v expecting %#v", other, user)
	}

	if err = s.DeleteUser("User"); err != nil {
		t.Fatalf("unable to delete user: %s", "User")
	}

}
