package model

import "gorm.io/gorm"

// Level Model
type Level struct {
	gorm.Model
	Level           string `json:"level,omitempty" gorm:"not null;unique"`
	Description     string `json:"description,omitempty" gorm:"not null;unique"`
	LevelPrivileges []LevelPrivileges
}
