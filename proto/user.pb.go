// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto/user.proto

It has these top-level messages:
	UserReply
	CreateUserRequest
	ChangeNicknameRequest
	ChangePasswordRequest
	CheckPasswordRequest
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type UserReply struct {
	Status int32  `protobuf:"varint,10,opt,name=status" json:"status,omitempty"`
	Msgid  string `protobuf:"bytes,20,opt,name=msgid" json:"msgid,omitempty"`
}

func (m *UserReply) Reset()                    { *m = UserReply{} }
func (m *UserReply) String() string            { return proto1.CompactTextString(m) }
func (*UserReply) ProtoMessage()               {}
func (*UserReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserReply) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *UserReply) GetMsgid() string {
	if m != nil {
		return m.Msgid
	}
	return ""
}

type CreateUserRequest struct {
	Username string `protobuf:"bytes,10,opt,name=username" json:"username,omitempty"`
	Nickname string `protobuf:"bytes,15,opt,name=nickname" json:"nickname,omitempty"`
	Password string `protobuf:"bytes,20,opt,name=password" json:"password,omitempty"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto1.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CreateUserRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *CreateUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type ChangeNicknameRequest struct {
	Id       int64  `protobuf:"varint,10,opt,name=id" json:"id,omitempty"`
	Nickname string `protobuf:"bytes,20,opt,name=nickname" json:"nickname,omitempty"`
}

func (m *ChangeNicknameRequest) Reset()                    { *m = ChangeNicknameRequest{} }
func (m *ChangeNicknameRequest) String() string            { return proto1.CompactTextString(m) }
func (*ChangeNicknameRequest) ProtoMessage()               {}
func (*ChangeNicknameRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ChangeNicknameRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ChangeNicknameRequest) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

type ChangePasswordRequest struct {
	Username    string `protobuf:"bytes,10,opt,name=username" json:"username,omitempty"`
	Oldpassowrd string `protobuf:"bytes,15,opt,name=oldpassowrd" json:"oldpassowrd,omitempty"`
	Password    string `protobuf:"bytes,20,opt,name=password" json:"password,omitempty"`
}

func (m *ChangePasswordRequest) Reset()                    { *m = ChangePasswordRequest{} }
func (m *ChangePasswordRequest) String() string            { return proto1.CompactTextString(m) }
func (*ChangePasswordRequest) ProtoMessage()               {}
func (*ChangePasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ChangePasswordRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ChangePasswordRequest) GetOldpassowrd() string {
	if m != nil {
		return m.Oldpassowrd
	}
	return ""
}

func (m *ChangePasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CheckPasswordRequest struct {
	Username string `protobuf:"bytes,10,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,20,opt,name=password" json:"password,omitempty"`
}

func (m *CheckPasswordRequest) Reset()                    { *m = CheckPasswordRequest{} }
func (m *CheckPasswordRequest) String() string            { return proto1.CompactTextString(m) }
func (*CheckPasswordRequest) ProtoMessage()               {}
func (*CheckPasswordRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CheckPasswordRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *CheckPasswordRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto1.RegisterType((*UserReply)(nil), "proto.UserReply")
	proto1.RegisterType((*CreateUserRequest)(nil), "proto.CreateUserRequest")
	proto1.RegisterType((*ChangeNicknameRequest)(nil), "proto.ChangeNicknameRequest")
	proto1.RegisterType((*ChangePasswordRequest)(nil), "proto.ChangePasswordRequest")
	proto1.RegisterType((*CheckPasswordRequest)(nil), "proto.CheckPasswordRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserReply, error)
	ChangeNickname(ctx context.Context, in *ChangeNicknameRequest, opts ...grpc.CallOption) (*UserReply, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*UserReply, error)
	CheckPassword(ctx context.Context, in *CheckPasswordRequest, opts ...grpc.CallOption) (*UserReply, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := grpc.Invoke(ctx, "/proto.UserService/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangeNickname(ctx context.Context, in *ChangeNicknameRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := grpc.Invoke(ctx, "/proto.UserService/ChangeNickname", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := grpc.Invoke(ctx, "/proto.UserService/ChangePassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) CheckPassword(ctx context.Context, in *CheckPasswordRequest, opts ...grpc.CallOption) (*UserReply, error) {
	out := new(UserReply)
	err := grpc.Invoke(ctx, "/proto.UserService/CheckPassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*UserReply, error)
	ChangeNickname(context.Context, *ChangeNicknameRequest) (*UserReply, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*UserReply, error)
	CheckPassword(context.Context, *CheckPasswordRequest) (*UserReply, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangeNickname_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeNicknameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangeNickname(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/ChangeNickname",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangeNickname(ctx, req.(*ChangeNicknameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_CheckPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CheckPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserService/CheckPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CheckPassword(ctx, req.(*CheckPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "ChangeNickname",
			Handler:    _UserService_ChangeNickname_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserService_ChangePassword_Handler,
		},
		{
			MethodName: "CheckPassword",
			Handler:    _UserService_CheckPassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}

func init() { proto1.RegisterFile("proto/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x51, 0x4f, 0x83, 0x30,
	0x14, 0x85, 0x1d, 0x86, 0x45, 0xee, 0xe2, 0x9c, 0x0d, 0x1a, 0x32, 0x7d, 0x20, 0x3c, 0xed, 0x69,
	0x26, 0xfa, 0xa4, 0x4f, 0x46, 0xde, 0x17, 0x83, 0xf1, 0x07, 0x20, 0xdc, 0x30, 0xb2, 0x8d, 0xb2,
	0xb6, 0xb8, 0xf8, 0x3b, 0xfc, 0xc3, 0xa6, 0xa5, 0x45, 0xd9, 0x08, 0x71, 0x4f, 0x70, 0x7a, 0xe0,
	0x3b, 0xb7, 0xe7, 0xc2, 0xa4, 0x64, 0x54, 0xd0, 0xbb, 0x8a, 0x23, 0x9b, 0xab, 0x57, 0x62, 0xab,
	0x47, 0xf0, 0x08, 0xce, 0x3b, 0x47, 0x16, 0x61, 0xb9, 0xfe, 0x22, 0xd7, 0x30, 0xe4, 0x22, 0x16,
	0x15, 0xf7, 0xc0, 0x1f, 0xcc, 0xec, 0x48, 0x2b, 0xe2, 0x82, 0xbd, 0xe1, 0x59, 0x9e, 0x7a, 0xae,
	0x3f, 0x98, 0x39, 0x51, 0x2d, 0x82, 0x0c, 0x2e, 0x43, 0x86, 0xb1, 0xc0, 0x1a, 0xb0, 0xad, 0x90,
	0x0b, 0x32, 0x85, 0x33, 0x19, 0x52, 0xc4, 0x1b, 0x54, 0x10, 0x27, 0x6a, 0xb4, 0xf4, 0x8a, 0x3c,
	0x59, 0x29, 0xef, 0xa2, 0xf6, 0x8c, 0x96, 0x5e, 0x19, 0x73, 0xbe, 0xa3, 0xcc, 0xa4, 0x34, 0x3a,
	0x08, 0xe1, 0x2a, 0x5c, 0xc6, 0x45, 0x86, 0x0b, 0xfd, 0xb5, 0x09, 0x1b, 0x83, 0x95, 0xa7, 0x2a,
	0xe6, 0x34, 0xb2, 0xf2, 0xb4, 0x15, 0xe0, 0xb6, 0x03, 0x82, 0xad, 0x81, 0xbc, 0x6a, 0xec, 0x7f,
	0x26, 0xf6, 0x61, 0x44, 0xd7, 0xa9, 0x1c, 0x84, 0xee, 0x58, 0xaa, 0x87, 0xfe, 0x7b, 0xd4, 0x3b,
	0xf7, 0x02, 0xdc, 0x70, 0x89, 0xc9, 0xea, 0x98, 0xc4, 0x1e, 0xde, 0xfd, 0xb7, 0x05, 0x23, 0xd9,
	0xf5, 0x1b, 0xb2, 0xcf, 0x3c, 0x41, 0xf2, 0x04, 0xf0, 0xbb, 0x00, 0xe2, 0xd5, 0x8b, 0x9d, 0x1f,
	0xec, 0x64, 0x3a, 0xd1, 0x4e, 0xb3, 0xe8, 0xe0, 0x84, 0xbc, 0xc0, 0xb8, 0xdd, 0x29, 0xb9, 0x35,
	0xff, 0x77, 0x55, 0xdd, 0xcf, 0x30, 0x17, 0xdc, 0x63, 0xec, 0xdd, 0xbb, 0x93, 0xf1, 0x0c, 0xe7,
	0xad, 0x8e, 0xc8, 0x4d, 0x83, 0x38, 0x6c, 0xae, 0x8b, 0xf0, 0x31, 0x54, 0x47, 0x0f, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x24, 0xd0, 0xb5, 0x94, 0xe3, 0x02, 0x00, 0x00,
}
