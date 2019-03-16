// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: notify.proto

/*
Package notify is a generated protocol buffer package.

It is generated from these files:
	notify.proto

It has these top-level messages:
	SendEmailRequest
	Response
*/
package notify

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

// Client API for Notificator service

type NotificatorService interface {
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...client.CallOption) (*Response, error)
}

type notificatorService struct {
	c    client.Client
	name string
}

func NewNotificatorService(name string, c client.Client) NotificatorService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "notificator"
	}
	return &notificatorService{
		c:    c,
		name: name,
	}
}

func (c *notificatorService) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Notificator.SendEmail", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Notificator service

type NotificatorHandler interface {
	SendEmail(context.Context, *SendEmailRequest, *Response) error
}

func RegisterNotificatorHandler(s server.Server, hdlr NotificatorHandler, opts ...server.HandlerOption) error {
	type notificator interface {
		SendEmail(ctx context.Context, in *SendEmailRequest, out *Response) error
	}
	type Notificator struct {
		notificator
	}
	h := &notificatorHandler{hdlr}
	return s.Handle(s.NewHandler(&Notificator{h}, opts...))
}

type notificatorHandler struct {
	NotificatorHandler
}

func (h *notificatorHandler) SendEmail(ctx context.Context, in *SendEmailRequest, out *Response) error {
	return h.NotificatorHandler.SendEmail(ctx, in, out)
}