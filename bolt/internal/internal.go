package internal

import (
	"github.com/gogo/protobuf/proto"
	"github.com/tixu/Auth/users"
)

//go:generate protoc --gogo_out=. internal.proto

// MarshalUser encodes a user to binary format.
func MarshalUser(d *users.User) ([]byte, error) {
	return proto.Marshal(&User{
		ID:           d.ID,
		Username:     d.Name,
		PasswordHash: d.PasswordHash,
		Role:         d.Role,
		Email:        d.Email,
	})
}

// UnmarshalUser decodes a user from a binary data.
func UnmarshalUser(data []byte, u *users.User) error {
	var pb User
	if err := proto.Unmarshal(data, &pb); err != nil {
		return err
	}

	u.ID = pb.ID
	u.Name = pb.Username
	u.Email = pb.Email
	u.Role = pb.Role
	u.PasswordHash = pb.PasswordHash

	return nil
}
