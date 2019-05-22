package handler

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/senonerk/sup/srv/notificator/proto/notificator"

	"github.com/senonerk/sup/shared/tags"

	"github.com/senonerk/sup/srv/auth/proto"

	"github.com/m1ome/randstr"

	"github.com/senonerk/sup/srv/profile/db"
	"github.com/senonerk/sup/srv/profile/models"

	proto "github.com/senonerk/sup/srv/profile/proto/profile"
)

// ProfileService srv
type ProfileService struct {
	Auth        auth.AuthService
	Notificator notify.NotificatorService
}

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
		PermissionTag: tags.PERMISSION_EMAIL,
		Grant:         false,
	})

	if err != nil {
		return err
	}

	go s.Notificator.SendEmail(ctx, &notify.SendEmailRequest{
		Recipient: req.NewEmail,
		Subject:   "New Email",
		Body:      "Hello " + p.EmailToken,
	})

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
		PermissionTag: tags.PERMISSION_EMAIL,
		Grant:         true,
	})

	if err != nil {
		return err
	}

	return nil
}

// GetInfo returns all info about user
func (ProfileService) GetInfo(ctx context.Context, req *proto.GetInfoRequest, res *proto.GetInfoResponse) error {
	p, err := getProfileByID(req.UserID)
	if err != nil {
		return err
	}

	res.FirstName = p.FirstName
	res.LastName = p.LastName
	res.Birth = p.BirthDate.Unix()
	res.Status = p.Status
	res.Email = p.Email

	return nil
}

func (ProfileService) Search(ctx context.Context, req *proto.SearchRequest, res *proto.SearchResponse) error {
	q := fmt.Sprintf("%%%v%%", req.Query)
	var profiles []models.Profile
	qres := db.D().Where("(first_name LIKE ? OR last_name LIKE ?) AND first_name != '' AND last_name != ''", q, q).Find(&profiles)

	var ress []*proto.SearchUser
	for _, p := range profiles {
		ress = append(ress, &proto.SearchUser{
			UserID: p.UserID,
			Name:   p.FirstName + " " + p.LastName,
		})
	}

	res.Users = ress

	return qres.Error
}

func getProfileByID(id string) (*models.Profile, error) {
	p := models.Profile{
		UserID: id,
	}
	res := db.D().Where("user_id = ?", id).FirstOrCreate(&p)
	return &p, res.Error
}
