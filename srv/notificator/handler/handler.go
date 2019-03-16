package handler

import (
	"context"

	"github.com/micro/go-log"

	proto "github.com/senonerk/sup/srv/notificator/proto/notificator"
)

// NotifyService struct for event handler
type NotifyService struct{}

// SendEmail srv
func (n *NotifyService) SendEmail(ctx context.Context, req *proto.SendEmailRequest, res *proto.Response) error {
	log.Logf("SEND EMAIL!!!: %v", req)
	return nil
}
