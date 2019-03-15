package handler

import (
	"context"
	"errors"
	"time"

	"github.com/senonerk/sup/srv/auth/proto"

	"github.com/m1ome/randstr"

	"github.com/senonerk/sup/srv/profile/db"
	"github.com/senonerk/sup/srv/profile/models"

	proto "github.com/senonerk/sup/srv/profile/proto/profile"
)

// ProfileService srv
type ProfileService struct {
	Auth auth.AuthService
}

const emailPermissionTag = "EMAIL_VERIFIED"

// UpdateInfo updates First,Last -Name and Birthdate
func (ProfileService) UpdateInfo(ctx context.Context, req *proto.UpdateInfoRequest, res *proto.Response) error {
	p, err := getProfileByID(req.UserID)
	if err != nil {
		return err
	}

	p.FirstName = req.FirstName
	p.LastName = req.LastName
	p.BirthDate = time.Unix(req.Birth, 0)

	t := db.D().Save(p)

	return t.Error
}

// UpdateStatus updates status message
func (ProfileService) UpdateStatus(ctx context.Context, req *proto.UpdateStatusRequest, res *proto.Response) error {
	p, err := getProfileByID(req.UserID)
	if err != nil {
		return err
	}

	p.Status = req.NewStatus

	t := db.D().Save(p)

	return t.Error
}

// UpdateEmail sends confirmation email
func (s *ProfileService) UpdateEmail(ctx context.Context, req *proto.UpdateEmailRequest, res *proto.Response) error {
	p, err := getProfileByID(req.UserID)
	if err != nil {
		return err
	}

	p.Email = req.NewEmail
	p.EmailToken = randstr.GetString(60)
	p.EmailTokenExpires = time.Now().Add(time.Hour*0 + time.Minute*1 + time.Second*0)

	_, err = s.Auth.SetPermission(ctx, &auth.SetPermissionRequest{
		UserID:        req.UserID,
		PermissionTag: emailPermissionTag,
		Grant:         false,
	})

	if err != nil {
		return err
	}

	// TODO: send email

	t := db.D().Save(p)

	return t.Error
}

// ConfirmEmail checks if tokens match
func (s *ProfileService) ConfirmEmail(ctx context.Context, req *proto.ConfirmEmailRequest, res *proto.Response) error {
	p, err := getProfileByID(req.UserID)
	if err != nil {
		return err
	}

	if p.EmailToken != req.EmailToken {
		return errors.New("Invalid token")
	}

	if p.EmailTokenExpires.Unix() < time.Now().Unix() {
		return errors.New("Too late! Token expired")
	}

	_, err = s.Auth.SetPermission(ctx, &auth.SetPermissionRequest{
		UserID:        req.UserID,
		PermissionTag: emailPermissionTag,
		Grant:         true,
	})

	if err != nil {
		return err
	}

	return nil
}

func getProfileByID(id string) (*models.Profile, error) {
	var p models.Profile
	res := db.D().Where("user_id = ?", id).First(&p)
	return &p, res.Error
}
