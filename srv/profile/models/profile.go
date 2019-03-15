package models

import "time"

// Profile model for profiles
type Profile struct {
	UserID            string `gorm:"type:char(24);primary_key"`
	FirstName         string `gorm:"type:varchar(30)"`
	LastName          string `gorm:"type:varchar(50)"`
	BirthDate         time.Time
	Status            string `gorm:"type:varchar(100)"`
	Email             string `gorm:"type:varchar(100);unique_index"`
	EmailToken        string `gorm:"type:varchar"`
	EmailTokenExpires time.Time
}
