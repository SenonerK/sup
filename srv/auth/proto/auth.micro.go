// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	VerifyTokenRequest
	VerifyTokenResponse
	UserRequest
	Response
	LoginResponse
	CheckPermissionsRequest
	ChangePasswordRequest
	SetPermissionRequest
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthService interface {
	Login(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*Response, error)
	CheckPermissions(ctx context.Context, in *CheckPermissionsRequest, opts ...client.CallOption) (*Response, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...client.CallOption) (*Response, error)
	SetPermission(ctx context.Context, in *SetPermissionRequest, opts ...client.CallOption) (*Response, error)
	VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...client.CallOption) (*VerifyTokenResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "auth"
	}
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Login(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*LoginResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.Login", in)
	out := new(LoginResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Register(ctx context.Context, in *UserRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.Register", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CheckPermissions(ctx context.Context, in *CheckPermissionsRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.CheckPermissions", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.ChangePassword", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) SetPermission(ctx context.Context, in *SetPermissionRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Auth.SetPermission", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) VerifyToken(ctx context.Context, in *VerifyTokenRequest, opts ...client.CallOption) (*VerifyTokenResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.VerifyToken", in)
	out := new(VerifyTokenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	Login(context.Context, *UserRequest, *LoginResponse) error
	Register(context.Context, *UserRequest, *Response) error
	CheckPermissions(context.Context, *CheckPermissionsRequest, *Response) error
	ChangePassword(context.Context, *ChangePasswordRequest, *Response) error
	SetPermission(context.Context, *SetPermissionRequest, *Response) error
	VerifyToken(context.Context, *VerifyTokenRequest, *VerifyTokenResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		Login(ctx context.Context, in *UserRequest, out *LoginResponse) error
		Register(ctx context.Context, in *UserRequest, out *Response) error
		CheckPermissions(ctx context.Context, in *CheckPermissionsRequest, out *Response) error
		ChangePassword(ctx context.Context, in *ChangePasswordRequest, out *Response) error
		SetPermission(ctx context.Context, in *SetPermissionRequest, out *Response) error
		VerifyToken(ctx context.Context, in *VerifyTokenRequest, out *VerifyTokenResponse) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) Login(ctx context.Context, in *UserRequest, out *LoginResponse) error {
	return h.AuthHandler.Login(ctx, in, out)
}

func (h *authHandler) Register(ctx context.Context, in *UserRequest, out *Response) error {
	return h.AuthHandler.Register(ctx, in, out)
}

func (h *authHandler) CheckPermissions(ctx context.Context, in *CheckPermissionsRequest, out *Response) error {
	return h.AuthHandler.CheckPermissions(ctx, in, out)
}

func (h *authHandler) ChangePassword(ctx context.Context, in *ChangePasswordRequest, out *Response) error {
	return h.AuthHandler.ChangePassword(ctx, in, out)
}

func (h *authHandler) SetPermission(ctx context.Context, in *SetPermissionRequest, out *Response) error {
	return h.AuthHandler.SetPermission(ctx, in, out)
}

func (h *authHandler) VerifyToken(ctx context.Context, in *VerifyTokenRequest, out *VerifyTokenResponse) error {
	return h.AuthHandler.VerifyToken(ctx, in, out)
}
