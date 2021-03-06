package handler

import (
	"context"
	"errors"
	"time"

	"github.com/senonerk/sup/shared/tags"

	"github.com/satori/go.uuid"
	"github.com/senonerk/sup/srv/auth/proto/events"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"

	"github.com/senonerk/sup/srv/auth/jwt"

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

type authService struct {
	NewUserPub micro.Publisher
}

// New returns a service implementation
func New(newUserPub micro.Publisher) *authService {
	return &authService{
		NewUserPub: newUserPub,
	}
}

func (a *authService) Login(ctx context.Context, req *proto.UserRequest, res *proto.LoginResponse) error {
	user, err := getUser(req.Username)

	exerr := errors.New("Username/Password is incorrect")
	if _, ok := err.(*mongodm.NotFoundError); ok {
		return exerr
	} else if err != nil {
		return err
	}

	if err := salter.CompareHMAC(req.Password, user.Password); err != nil {
		return exerr
	}

	if err := checkPermissions(user, []string{tags.PERMISSION_LOGIN}); err != nil {
		return err
	}

	duration := time.Hour * 24 * 2

	token, err := jwt.GenerateToken(user.Id.Hex(), duration)
	if err != nil {
		return err
	}

	res.Token = token
	return nil
}

func (a *authService) Register(ctx context.Context, req *proto.UserRequest, res *proto.Response) error {
	_, err := getUser(req.Username)

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

	// Create user and add standard permission to allow login to the system
	user := &models.User{
		UserName: req.Username,
		Password: salt,
		Permissions: []models.Permission{
			models.Permission{
				Tag:   tags.PERMISSION_LOGIN,
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

	err = a.NewUserPub.Publish(ctx, &event.NewUserEvent{
		Id:        uuid.NewV4().String(),
		Timestamp: time.Now().Unix(),
		UserID:    user.Id.Hex(),
	})

	if err != nil {
		if err = user.Delete(); err != nil {
			log.Fatalf("User created, pub and emergency delete failed. UserID: %s", user.Id.Hex())
		}
		return err
	}

	return nil
}

func (a *authService) CheckPermissions(ctx context.Context, req *proto.CheckPermissionsRequest, res *proto.Response) error {
	user, err := getUserByID(req.UserID)
	if err != nil {
		return err
	}

	return checkPermissions(user, req.PermissionTags)
}

func (a *authService) ChangePassword(ctx context.Context, req *proto.ChangePasswordRequest, res *proto.Response) error {
	user, err := getUserByID(req.UserID)
	if err != nil {
		return err
	}

	pass, err := salter.GenerateHMAC(req.NewPassword)
	if err != nil {
		return err
	}

	user.Password = pass
	return user.Save()
}

func (a *authService) SetPermission(ctx context.Context, req *proto.SetPermissionRequest, res *proto.Response) error {
	user, err := getUserByID(req.UserID)
	if err != nil {
		return err
	}

	var found bool
	for i, p := range user.Permissions {
		if p.Tag == req.PermissionTag {
			user.Permissions[i].Grant = req.Grant
			found = true
		}
	}

	if !found {
		user.Permissions = append(user.Permissions, models.Permission{
			Tag:   req.PermissionTag,
			Grant: req.Grant,
		})
	}

	return user.Save()
}

func (a *authService) VerifyToken(ctx context.Context, req *proto.VerifyTokenRequest, res *proto.VerifyTokenResponse) error {
	userid, err := jwt.ValidateToken(req.Token)
	if err != nil {
		return err
	}

	res.UserID = userid

	return nil
}

func (a *authService) CheckPassword(ctx context.Context, req *proto.CheckPasswordRequest, res *proto.Response) error {
	user, err := getUserByID(req.UserID)
	if err != nil {
		return err
	}

	if err := salter.CompareHMAC(req.Password, user.Password); err != nil {
		return errors.New("Wrong password")
	}

	return nil
}

func (a *authService) NewToken(ctx context.Context, req *proto.NewTokenRequest, res *proto.LoginResponse) error {
	// Make db call to ensure that user still exists
	_, err := getUserByID(req.UserID)
	if err != nil {
		return err
	}

	token, err := jwt.GenerateToken(req.UserID, time.Hour*time.Duration(req.ExpiresInHours))
	if err != nil {
		return err
	}

	res.Token = token

	return nil
}

func getUserByID(userID string) (*models.User, error) {
	var user models.User
	if !bson.IsObjectIdHex(userID) {
		return nil, errors.New("Invalid UserID")
	}

	err := db.D().FindOne(bson.M{
		"_id":     bson.ObjectIdHex(userID),
		"deleted": false,
	}).Exec(&user)
	return &user, err
}

func getUser(username string) (*models.User, error) {
	var user models.User
	err := db.D().FindOne(bson.M{
		"username": username,
		"deleted":  false,
	}).Exec(&user)
	return &user, err
}

func checkPermissions(user *models.User, permissions []string) error {
	for _, p := range permissions {
		if pr, ok := user.GetPermission(p); !ok || !pr.Grant {
			return errors.New("Unauthorized")
		}
	}
	return nil
}
