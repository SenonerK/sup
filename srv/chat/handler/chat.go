package handler

import (
	"context"

	"gopkg.in/mgo.v2/bson"

	"github.com/senonerk/sup/srv/chat/db"
	"github.com/senonerk/sup/srv/chat/models"

	proto "github.com/senonerk/sup/srv/chat/proto/chat"
)

type chatService struct{}

// New returns a service implementation
func New() *chatService {
	return &chatService{}
}

func (chatService) Send(ctx context.Context, req *proto.SendRequest, res *proto.Response) error {

	usr := &models.Chat{
		FromID:  req.FromUserID,
		ToID:    req.ToUserID,
		Message: req.Message,
	}
	err, _ := db.D().New(usr)
	usr.Save()

	return err
}

func (chatService) Receive(ctx context.Context, req *proto.ReceiveRequest, res *proto.ReceiveResponse) error {

	r := []*proto.UserChat{}
	err := db.D().Pipe([]bson.M{
		{"$match": bson.M{"toID": req.UserID}},
		{"$group": bson.M{"_id": "$fromID", "messages": bson.M{"$push": "$message"}}},
	}).All(&r)

	if err != nil {
		return err
	}

	res.Chats = r

	return nil
}
