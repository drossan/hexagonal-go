package model

import "gorm.io/gorm"

// LevelPrivileges Model
type LevelPrivileges struct {
	gorm.Model
	LevelID uint
	FormID  uint
	Form    Form
	Read    bool `json:"read,omitempty" gorm:"not null"`
	Write   bool `json:"write,omitempty" gorm:"not null"`
}
