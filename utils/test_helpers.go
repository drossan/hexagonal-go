package utils

import (
	"testing"

	"github.com/drossan/core-api/domain/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupTestDB(t *testing.T) *gorm.DB {
	testDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	err = testDB.AutoMigrate(
		&model.User{},
		&model.Form{},
		&model.MenuTree{},
		&model.Level{},
		&model.LevelPrivileges{},
	)
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}
	return testDB
}

func ResetTestDB(db *gorm.DB, t *testing.T) {
	err := db.Migrator().DropTable(
		&model.User{},
		&model.Form{},
		&model.MenuTree{},
		&model.Level{},
		&model.LevelPrivileges{},
	)
	if err != nil {
		t.Fatalf("Failed to drop tables: %v", err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Form{},
		&model.MenuTree{},
		&model.Level{},
		&model.LevelPrivileges{},
	)
	if err != nil {
		t.Fatalf("Failed to migrate database: %v", err)
	}
}
