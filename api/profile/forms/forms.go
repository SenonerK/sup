package forms

import "time"

type UpdateInfo struct {
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	BirthDate time.Time `json:"birthdate"`
}

type UpdateStatus struct {
	NewStatus string `json:"newstatus"`
}

type UpdateEmail struct {
	NewEmail string `json:"newemail"`
}

type ConfirmEmail struct {
	Token string `json:"token"`
}
