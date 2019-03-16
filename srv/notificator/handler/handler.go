package handler

import (
	"context"

	"gopkg.in/gomail.v2"

	proto "github.com/senonerk/sup/srv/notificator/proto/notificator"
)

// NotifyService struct for event handler
type NotifyService struct {
	Mailer chan *gomail.Message
}

// SendEmail srv
func (n *NotifyService) SendEmail(ctx context.Context, req *proto.SendEmailRequest, res *proto.Response) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", "4FI.sup@gmail.com")
	mail.SetHeader("To", "supclient@mailinator.com") // TODO: change to req.Recipient for production
	mail.SetHeader("Subject", req.Subject)
	mail.SetBody("text/html", req.Body)

	n.Mailer <- mail

	return nil
}
