package models

import (
	"github.com/zebresel-com/mongodm"
)

// User ODM
type Chat struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`
	FromID               string `json:"from" bson:"fromID"`
	ToID                 string `json:"to" bson:"toID"`
	Message              string `json:"message" bson:"message"`
}
