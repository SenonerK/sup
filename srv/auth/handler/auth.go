package handler

import (
	"context"
	"errors"

	mm "github.com/micro/go-micro/errors"

	proto "github.com/senonerk/sup/srv/auth/proto"
)

const (
	FQDN = "srv.auth."
)

type authService struct{}

func New() *authService {
	return new(authService)
}

func (a *authService) Login(ctx context.Context, req *proto.UserRequest, res *proto.LoginResponse) error {
	res.Token = "Fake"
	return nil
}

func (a *authService) Register(ctx context.Context, req *proto.UserRequest, res *proto.Response) error {
	return mm.NotFound(FQDN+"register", "not implemented")
}

func (a *authService) CheckPermissions(ctx context.Context, req *proto.CheckPermissionsRequest, res *proto.Response) error {
	return nil
}

func (a *authService) ChangePassword(ctx context.Context, req *proto.ChangePasswordRequest, res *proto.Response) error {
	return nil
}

func (a *authService) SetPermission(ctx context.Context, req *proto.SetPermissionRequest, res *proto.Response) error {
	return nil
}

func (a *authService) VerifyToken(ctx context.Context, req *proto.VerifyTokenRequest, res *proto.VerifyTokenResponse) error {
	if req.Token == "VALID_TOKEN" {
		res.UserID = "123"
		return nil
	}
	return errors.New("test")
}

func (a *authService) CheckPassword(ctx context.Context, req *proto.CheckPasswordRequest, res *proto.Response) error {
	return nil
}
