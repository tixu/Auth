package internal_test

import (
	"reflect"
	"testing"

	"github.com/tixu/Auth/bolt/internal"
	"github.com/tixu/Auth/users"
)

// Ensure dial can be marshaled and unmarshaled.
func TestMarshalUser(t *testing.T) {
	v := users.User{
		ID:           1,
		Name:         "Xavier",
		PasswordHash: "token",
		Email:        "xavier@email.com",
		Role:         "Role",
	}

	var other users.User
	if buf, err := internal.MarshalUser(&v); err != nil {
		t.Fatal(err)
	} else if err := internal.UnmarshalUser(buf, &other); err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(v, other) {
		t.Fatalf("unexpected copy: %#v", other)
	}
}
