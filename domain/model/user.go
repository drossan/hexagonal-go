package model

import (
	"gorm.io/gorm"
	"time"
)

type RecoverPass struct {
	Token string
}

type Password struct {
	Password string
}

// User Model
type User struct {
	gorm.Model
	Username        string `json:"username,omitempty" gorm:"not null;unique"`
	Email           string `json:"email,omitempty" gorm:"not null;unique"`
	FullName        string `json:"fullname,omitempty" gorm:"not null"`
	Password        string `json:"password,omitempty" gorm:"not null;type:varchar(256)"`
	ConfirmPassword string `json:"confirmPassword,omitempty" gorm:"-"`
	Picture         string `json:"picture,omitempty"`
	LevelID         uint
	Level           Level
	Token           string    `json:"token,omitempty"`
	Failure         int       `json:"failure,omitempty" gorm:"default:0"`
	CreatedAt       time.Time `gorm:"type:datetime"`
}
