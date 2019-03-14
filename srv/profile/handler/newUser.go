package handler

import (
	"context"
	"errors"

	"github.com/senonerk/sup/srv/profile/db"
	"github.com/senonerk/sup/srv/profile/models"

	e "github.com/senonerk/sup/srv/auth/proto/events"
)

// NewUserSubscriber handler
type NewUserSubscriber struct{}

// ProcessEvent do stuff
func (s *NewUserSubscriber) ProcessEvent(ctx context.Context, event *e.NewUserEvent) error {
	res := db.D().Create(&models.Profile{
		UserID: event.UserID,
	})

	if res.Error != nil {
		return errors.New("Error while creating profile")
	}

	return nil
}
