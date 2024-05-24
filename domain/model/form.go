package model

import "gorm.io/gorm"

// Form Model
type Form struct {
	gorm.Model
	Title            string `json:"title,omitempty" gorm:"not null;"`
	Icon             string `json:"icon,omitempty" gorm:"not null;"`
	Link             string `json:"link,omitempty" gorm:"not null;"`
	Color            string `json:"color" gorm:"not null;"`
	Count            string `json:"count,omitempty"`
	Order            int    `json:"order,omitempty" gorm:"not null;"`
	TotalCount       int    `gorm:"-"`
	Setting          bool   `json:"setting,omitempty" gorm:"not null;"`
	PublicToIntranet bool   `json:"public_to_intranet,omitempty" gorm:"not null;"`
	MenuTree         MenuTree
	MenuTreeID       *uint
	Condition        string `json:"condition,omitempty" gorm:"not null;"`
	PathAPI          string `json:"path_api" gorm:"not null;"`
}
