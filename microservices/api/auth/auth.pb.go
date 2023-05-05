// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/auth/auth.proto

package auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa33bda803119f75, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa33bda803119f75, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LoginResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type GetUserRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa33bda803119f75, []int{2}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetUserResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa33bda803119f75, []int{3}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetUserResponse) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GetUserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type CreateUserRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa33bda803119f75, []int{4}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa33bda803119f75, []int{5}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateUserResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa33bda803119f75, []int{6}
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

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "auth.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "auth.LoginResponse")
	proto.RegisterType((*GetUserRequest)(nil), "auth.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "auth.GetUserResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "auth.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "auth.CreateUserResponse")
	proto.RegisterType((*User)(nil), "auth.User")
}

func init() { proto.RegisterFile("api/auth/auth.proto", fileDescriptor_aa33bda803119f75) }

var fileDescriptor_aa33bda803119f75 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x0d, 0x05, 0x44, 0x47, 0xc5, 0xb8, 0x60, 0x6c, 0x7a, 0x22, 0x3d, 0x79, 0x6a, 0x89, 0x26,
	0x1e, 0x25, 0x6a, 0xa2, 0x17, 0x4f, 0x10, 0x2f, 0xde, 0x16, 0x3a, 0xd2, 0x8d, 0xb6, 0x5b, 0x77,
	0xb6, 0xfa, 0x6b, 0x7e, 0x9e, 0xa1, 0xbb, 0x85, 0xb6, 0x04, 0x12, 0x2f, 0x4d, 0x67, 0x67, 0xe6,
	0xcd, 0x7b, 0x6f, 0x76, 0x61, 0xc0, 0x33, 0x11, 0xf2, 0x5c, 0xc7, 0xc5, 0x27, 0xc8, 0x94, 0xd4,
	0x92, 0x75, 0x56, 0xff, 0xfe, 0x13, 0x9c, 0xbc, 0xc8, 0xa5, 0x48, 0xa7, 0xf8, 0x95, 0x23, 0x69,
	0xe6, 0xc1, 0x61, 0x4e, 0xa8, 0x52, 0x9e, 0xa0, 0xdb, 0x1a, 0xb5, 0xae, 0x8e, 0xa6, 0xeb, 0x78,
	0x95, 0xcb, 0x38, 0xd1, 0x8f, 0x54, 0x91, 0xeb, 0x98, 0x5c, 0x19, 0xfb, 0x13, 0x38, 0xb5, 0x38,
	0x94, 0xc9, 0x94, 0x90, 0x0d, 0xa1, 0xab, 0xe5, 0x07, 0xa6, 0x16, 0xc5, 0x04, 0xcc, 0x85, 0x5e,
	0x82, 0x44, 0x7c, 0x89, 0x16, 0xa1, 0x0c, 0xfd, 0x11, 0xf4, 0x9f, 0x51, 0xbf, 0x12, 0xaa, 0x92,
	0x4a, 0x1f, 0x1c, 0x11, 0xd9, 0x76, 0x47, 0x44, 0xfe, 0x0c, 0xce, 0xd6, 0x15, 0x76, 0x48, 0xa3,
	0xa4, 0xc6, 0xde, 0x69, 0xb0, 0x1f, 0x42, 0x17, 0x13, 0x2e, 0x3e, 0xdd, 0xb6, 0x21, 0x54, 0x04,
	0x3e, 0x87, 0xf3, 0x47, 0x85, 0x5c, 0x63, 0x75, 0xf2, 0x3e, 0x13, 0xd6, 0x30, 0x4e, 0x05, 0xa6,
	0x66, 0x4d, 0xbb, 0x61, 0xcd, 0x1d, 0xb0, 0xea, 0x88, 0x1d, 0xd4, 0x77, 0x3b, 0x13, 0x41, 0x67,
	0xd5, 0xf9, 0x2f, 0xb1, 0x7b, 0xf8, 0x6c, 0x14, 0x74, 0x2a, 0x0a, 0xae, 0x7f, 0x5b, 0x70, 0x7c,
	0x9f, 0xeb, 0x78, 0x86, 0xea, 0x5b, 0x2c, 0x90, 0x8d, 0xa1, 0x5b, 0x2c, 0x94, 0xb1, 0xa0, 0xb8,
	0x34, 0xd5, 0x5b, 0xe2, 0x0d, 0x6a, 0x67, 0x56, 0xd1, 0x2d, 0xf4, 0xec, 0x7e, 0xd8, 0xd0, 0xe4,
	0xeb, 0x0b, 0xf5, 0x2e, 0x1a, 0xa7, 0xb6, 0x6f, 0x02, 0xb0, 0xf1, 0x87, 0x5d, 0x9a, 0xa2, 0xad,
	0xa5, 0x78, 0xee, 0x76, 0xc2, 0x00, 0x3c, 0x8c, 0xdf, 0x82, 0xa5, 0xd0, 0x71, 0x3e, 0x0f, 0x16,
	0x32, 0x09, 0xdf, 0x73, 0x12, 0x32, 0x8d, 0xb9, 0x9a, 0x4b, 0x15, 0x26, 0x62, 0xa1, 0x24, 0x19,
	0x49, 0x14, 0x96, 0xcf, 0x60, 0x7e, 0x50, 0x3c, 0x81, 0x9b, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xe5, 0x86, 0xa4, 0x55, 0x19, 0x03, 0x00, 0x00,
}
