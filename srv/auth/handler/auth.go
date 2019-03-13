package handler

import (
	"context"
	"errors"

	"github.com/zebresel-com/mongodm"

	"gopkg.in/mgo.v2/bson"

	"github.com/senonerk/sup/srv/auth/db"

	"github.com/senonerk/sup/srv/auth/models"
	proto "github.com/senonerk/sup/srv/auth/proto"
	"github.com/senonerk/sup/srv/auth/salter"
)

const (
	FQDN = "senonerk.sup.srv.auth"
)

type authService struct{}

// New returns a service implementation
func New() *authService {
	return new(authService)
}

func (a *authService) Login(ctx context.Context, req *proto.UserRequest, res *proto.LoginResponse) error {
	res.Token = "Fake"
	return nil
}

func (a *authService) Register(ctx context.Context, req *proto.UserRequest, res *proto.Response) error {
	err := db.D().FindOne(bson.M{
		"username": req.Username,
		"deleted":  false,
	}).Exec(&models.User{})

	if err != nil {
		if _, ok := err.(*mongodm.NotFoundError); !ok {
			return err
		}
	} else {
		return errors.New("User already exists")
	}

	salt, err := salter.GenerateHMAC(req.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		UserName: req.Username,
		Password: salt,
		Permissions: []models.Permission{
			models.Permission{
				Tag:   "LOGIN",
				Grant: true,
			},
		},
	}

	if err, _ = db.D().New(user); err != nil {
		return err
	}

	if err = user.Save(); err != nil {
		return err
	}

	return nil
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
