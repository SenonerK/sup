package models

import (
	"github.com/zebresel-com/mongodm"
)

// User ODM
type User struct {
	mongodm.DocumentBase `json:",inline" bson:",inline"`
	UserName             string       `json:"username" bson:"username"`
	Password             string       `json:"-" bson:"password"`
	Permissions          []Permission `json:"permissions" bson:"permissions"`
}

// Permission ODM
type Permission struct {
	Tag   string `json:"tag" bson:"tag"`
	Grant bool   `json:"grant" bson:"grant"`
}
