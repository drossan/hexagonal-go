package model

import "gorm.io/gorm"

type MenuTree struct {
	gorm.Model
	Title string `json:"title,omitempty" gorm:"not null;"`
	Icon  string `json:"icon,omitempty" gorm:"not null;"`
	Color string `json:"color" gorm:"not null;"`
	Order int    `json:"order,omitempty" gorm:"not null;"`
}
