package models

import (
	"time"

	"github.com/zebresel-com/mongodm"
)

// Chat ODM
type Chat struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`
	FromID               string    `json:"from" bson:"fromID"`
	ToID                 string    `json:"to" bson:"toID"`
	Message              string    `json:"message" bson:"message"`
	ReceivedAt           time.Time `json:"received" bson:"receivedAt"`
	ReadAt               time.Time `json:"read" bson:"readAt"`
}
