package forms

import "time"

type Send struct {
	ToUserID string `json:"to"`
	Message  string `json:"message"`
}

type Update struct {
	ChatID    string    `json:"chatid"`
	Timestamp time.Time `json:"time"`
}
