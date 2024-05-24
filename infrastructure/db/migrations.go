package db

import (
	"github.com/drossan/core-api/domain/model"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.User{},
		&model.Level{},
		&model.MenuTree{},
		&model.Form{},
		&model.LevelPrivileges{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
