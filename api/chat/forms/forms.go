package forms

type Send struct {
	ToUserID string `json:"to"`
	Message  string `json:"message"`
}
