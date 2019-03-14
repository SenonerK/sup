package handler

import (
	"context"
	"time"

	"github.com/senonerk/sup/srv/profile/db"
	"github.com/senonerk/sup/srv/profile/models"

	proto "github.com/senonerk/sup/srv/profile/proto/profile"
)

// ProfileService srv
type ProfileService struct{}

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
func (ProfileService) UpdateEmail(ctx context.Context, req *proto.UpdateEmailRequest, res *proto.Response) error {
	return nil
}

// ConfirmEmail checks if tokens match
func (ProfileService) ConfirmEmail(ctx context.Context, req *proto.ConfirmEmailRequest, res *proto.Response) error {
	return nil
}

func getProfileByID(id string) (*models.Profile, error) {
	var p models.Profile
	res := db.D().Where("user_id = ?", id).First(&p)
	return &p, res.Error
}
