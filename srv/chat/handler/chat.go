package handler

import (
	"context"
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/senonerk/sup/shared/tags"

	"github.com/senonerk/sup/srv/auth/proto"
	"github.com/senonerk/sup/srv/chat/db"
	"github.com/senonerk/sup/srv/chat/models"

	proto "github.com/senonerk/sup/srv/chat/proto/chat"
)

// ChatService struct
type ChatService struct {
	Auth auth.AuthService
}

// Send sends
func (c *ChatService) Send(ctx context.Context, req *proto.SendRequest, res *proto.Response) error {

	_, err := c.Auth.CheckPermissions(ctx, &auth.CheckPermissionsRequest{
		UserID:         req.FromUserID,
		PermissionTags: []string{tags.PERMISSION_EMAIL},
	})

	if err != nil {
		return errors.New("Email not confirmed")
	}

	chat := &models.Chat{
		FromID:     req.FromUserID,
		ToID:       req.ToUserID,
		Message:    req.Message,
		ReceivedAt: time.Unix(0, 0),
		ReadAt:     time.Unix(0, 0),
	}
	err, _ = db.D().New(chat)
	chat.Save()

	return err
}

// Receive receives
func (ChatService) Receive(ctx context.Context, req *proto.ReceiveRequest, res *proto.ReceiveResponse) error {

	var r []*models.Chat
	err := db.D().Find(bson.M{
		"toID":      req.UserID,
		"createdAt": bson.M{"$lt": time.Unix(req.From, 0)},
	}).Sort("-createdAt").Skip(int(req.Skip)).Limit(int(req.Amount)).Exec(&r)

	if err != nil {
		return err
	}

	res.Chats = convertChatsToProto(r)

	return nil
}

// ReceiveNew rn
func (ChatService) ReceiveNew(ctx context.Context, req *proto.ReceiveNewRequest, res *proto.ReceiveResponse) error {

	var r []*models.Chat
	err := db.D().Find(bson.M{
		"toID":       req.UserID,
		"receivedAt": time.Unix(0, 0),
		"readAt":     time.Unix(0, 0),
	}).Exec(&r)

	if err != nil {
		return err
	}

	res.Chats = convertChatsToProto(r)

	return nil
}

func (ChatService) Read(ctx context.Context, req *proto.UpdateRequest, res *proto.Response) error {

	if !bson.IsObjectIdHex(req.ChatID) {
		return errors.New("Invalid chat ID")
	}

	var r models.Chat
	err := db.D().FindOne(bson.M{
		"toID": req.UserID,
		"_id":  bson.ObjectIdHex(req.ChatID),
	}).Exec(&r)

	if err != nil {
		return err
	}

	r.ReadAt = time.Unix(req.When, 0)

	return r.Save()
}

func (ChatService) Received(ctx context.Context, req *proto.UpdateRequest, res *proto.Response) error {

	if !bson.IsObjectIdHex(req.ChatID) {
		return errors.New("Invalid chat ID")
	}

	var r models.Chat
	err := db.D().FindOne(bson.M{
		"toID": req.UserID,
		"_id":  bson.ObjectIdHex(req.ChatID),
	}).Exec(&r)

	if err != nil {
		return err
	}

	r.ReceivedAt = time.Unix(req.When, 0)

	return r.Save()
}

func convertChatsToProto(chats []*models.Chat) (res []*proto.UserChat) {
	for _, c := range chats {
		res = append(res, &proto.UserChat{
			FromID:     c.FromID,
			Message:    c.Message,
			ReceivedAt: c.ReceivedAt.Unix(),
			ReadAt:     c.ReadAt.Unix(),
			Deleted:    c.Deleted,
			Id:         c.Id.Hex(),
			CreatedAt:  c.CreatedAt.Unix(),
		})
	}
	return res
}
