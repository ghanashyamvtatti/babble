// Code generated by protoc-gen-go. DO NOT EDIT.
// source: models.proto

package models

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Post struct {
	Post                 string               `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
	Username             string               `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Post) Reset()         { *m = Post{} }
func (m *Post) String() string { return proto.CompactTextString(m) }
func (*Post) ProtoMessage()    {}
func (*Post) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{0}
}

func (m *Post) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Post.Unmarshal(m, b)
}
func (m *Post) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Post.Marshal(b, m, deterministic)
}
func (m *Post) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Post.Merge(m, src)
}
func (m *Post) XXX_Size() int {
	return xxx_messageInfo_Post.Size(m)
}
func (m *Post) XXX_DiscardUnknown() {
	xxx_messageInfo_Post.DiscardUnknown(m)
}

var xxx_messageInfo_Post proto.InternalMessageInfo

func (m *Post) GetPost() string {
	if m != nil {
		return m.Post
	}
	return ""
}

func (m *Post) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Post) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Post) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type User struct {
	FullName             string               `protobuf:"bytes,1,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Password             string               `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b5431a010549573, []int{1}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetFullName() string {
	if m != nil {
		return m.FullName
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Post)(nil), "models.Post")
	proto.RegisterType((*User)(nil), "models.User")
}

func init() {
	proto.RegisterFile("models.proto", fileDescriptor_0b5431a010549573)
}

var fileDescriptor_0b5431a010549573 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x91, 0x31, 0x4b, 0xc4, 0x40,
	0x10, 0x85, 0x59, 0x0d, 0x47, 0x32, 0x5a, 0x6d, 0x15, 0x22, 0x68, 0xb8, 0xea, 0x1a, 0x13, 0xd0,
	0xca, 0xf2, 0xfc, 0x01, 0x22, 0x87, 0x36, 0x36, 0xc7, 0x5e, 0x76, 0xee, 0x50, 0xb2, 0x37, 0xcb,
	0xce, 0x04, 0xff, 0x94, 0xad, 0xff, 0x4f, 0xb2, 0x7b, 0x6b, 0x6b, 0x69, 0x37, 0xef, 0xf1, 0x1e,
	0xef, 0x83, 0x81, 0x4b, 0x47, 0x16, 0x47, 0xee, 0x7c, 0x20, 0x21, 0xbd, 0x48, 0xaa, 0xb9, 0x39,
	0x10, 0x1d, 0x46, 0xec, 0xa3, 0xbb, 0x9b, 0xf6, 0xbd, 0xbc, 0x3b, 0x64, 0x31, 0xce, 0xa7, 0xe0,
	0xf2, 0x4b, 0x41, 0xf1, 0x4c, 0x2c, 0x5a, 0x43, 0xe1, 0x89, 0xa5, 0x56, 0xad, 0x5a, 0x55, 0x9b,
	0x78, 0xeb, 0x06, 0xca, 0x89, 0x31, 0x1c, 0x8d, 0xc3, 0xfa, 0x2c, 0xfa, 0xbf, 0x5a, 0x3f, 0x00,
	0x0c, 0x01, 0x8d, 0xa0, 0xdd, 0x1a, 0xa9, 0xcf, 0x5b, 0xb5, 0xba, 0xb8, 0x6b, 0xba, 0x34, 0xd7,
	0xe5, 0xb9, 0xee, 0x25, 0xcf, 0x6d, 0xaa, 0x53, 0x7a, 0x2d, 0x73, 0x75, 0xf2, 0x36, 0x57, 0x8b,
	0xbf, 0xab, 0xa7, 0xf4, 0x5a, 0x96, 0xdf, 0x0a, 0x8a, 0x57, 0xc6, 0xa0, 0xaf, 0xa0, 0xda, 0x4f,
	0xe3, 0xb8, 0x8d, 0x6c, 0x89, 0xb9, 0x9c, 0x8d, 0xa7, 0x99, 0xad, 0x81, 0xd2, 0x1b, 0xe6, 0x4f,
	0x0a, 0x36, 0x73, 0x67, 0xfd, 0x3f, 0xdc, 0x8f, 0xed, 0xdb, 0xb5, 0xe5, 0x5b, 0x1f, 0xe8, 0x03,
	0x07, 0xe9, 0x07, 0x72, 0x8e, 0x8e, 0xe9, 0x29, 0x7d, 0xfa, 0xd4, 0x6e, 0x11, 0xd5, 0xfd, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x07, 0x47, 0x5c, 0xc8, 0x01, 0x00, 0x00,
}
