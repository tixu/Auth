// Code generated by protoc-gen-go.
// source: internal.proto
// DO NOT EDIT!

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	internal.proto

It has these top-level messages:
	User
*/
package internal

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	ID           uint64 `protobuf:"varint,1,opt,name=ID,json=iD" json:"ID,omitempty"`
	Username     string `protobuf:"bytes,2,opt,name=Username,json=username" json:"Username,omitempty"`
	PasswordHash string `protobuf:"bytes,3,opt,name=PasswordHash,json=passwordHash" json:"PasswordHash,omitempty"`
	Email        string `protobuf:"bytes,4,opt,name=Email,json=email" json:"Email,omitempty"`
	Role         string `protobuf:"bytes,5,opt,name=Role,json=role" json:"Role,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*User)(nil), "internal.User")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x1a,
	0x18, 0xb9, 0x58, 0x42, 0x8b, 0x53, 0x8b, 0x84, 0xf8, 0xb8, 0x98, 0x3c, 0x5d, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x58, 0x82, 0x98, 0x32, 0x5d, 0x84, 0xa4, 0xb8, 0x38, 0x40, 0xe2, 0x79, 0x89, 0xb9,
	0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x1c, 0xa5, 0x50, 0xbe, 0x90, 0x12, 0x17, 0x4f,
	0x40, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a, 0x47, 0x62, 0x71, 0x86, 0x04, 0x33, 0x58, 0x9e,
	0xa7, 0x00, 0x49, 0x4c, 0x48, 0x84, 0x8b, 0xd5, 0x35, 0x37, 0x31, 0x33, 0x47, 0x82, 0x05, 0x2c,
	0xc9, 0x9a, 0x0a, 0xe2, 0x08, 0x09, 0x71, 0xb1, 0x04, 0xe5, 0xe7, 0xa4, 0x4a, 0xb0, 0x82, 0x05,
	0x59, 0x8a, 0xf2, 0x73, 0x52, 0x93, 0xd8, 0xc0, 0x6e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x7f, 0x0e, 0xb7, 0x7a, 0xa5, 0x00, 0x00, 0x00,
}
