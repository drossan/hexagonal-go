package seeder

import (
	"crypto/sha256"
	"fmt"
	"log"
	"time"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/infrastructure/db"
	"gorm.io/gorm"
)

func Seed() {
	dbConn := db.GetConnection()

	// Seed forms
	if isTableEmpty(dbConn, &model.Form{}) {
		seedForms(dbConn)
	}

	// Seed levels
	if isTableEmpty(dbConn, &model.Level{}) {
		seedLevels(dbConn)
	}

	// Seed level_privileges
	if isTableEmpty(dbConn, &model.LevelPrivileges{}) {
		seedLevelPrivileges(dbConn)
	}

	// Seed users
	if isTableEmpty(dbConn, &model.User{}) {
		seedUsers(dbConn)
	}
}

func isTableEmpty(dbConn *gorm.DB, model interface{}) bool {
	var count int64
	dbConn.Model(model).Count(&count)
	return count == 0
}

func seedForms(dbConn *gorm.DB) {
	forms := []model.Form{
		{
			Title:   "Usuarios",
			Icon:    "mdi-account-check-outline",
			Link:    "usuarios",
			Setting: true,
			PathAPI: "user|users",
			Order:   1,
		},
		{
			Title:   "Identidades",
			Icon:    "mdi-account-check-outline",
			Link:    "roles",
			Setting: true,
			PathAPI: "level|levels",
			Order:   2,
		},
		{
			Title:   "Formularios",
			Icon:    "mdi-account-check-outline",
			Link:    "formularios",
			Setting: true,
			PathAPI: "form|forms",
			Order:   3,
		},
		{
			Title:   "Menú Expansibles",
			Icon:    "mdi-file-tree-outline",
			Link:    "menu-expansible",
			Setting: true,
			PathAPI: "expanses-menu|expanses-menus",
			Count:   "menu_trees",
			Order:   4,
		},
		{
			Title:   "SMTP Config",
			Icon:    "mdi-email-outline",
			Link:    "smtp-config",
			Setting: true,
			PathAPI: "smtp-config",
			Count:   "smtp_config",
			Order:   5,
		},
		{
			Title:   "Email notificaciones",
			Icon:    "mdi-email",
			Link:    "notificaciones-email",
			Setting: true,
			PathAPI: "email-notifications",
			Count:   "email_notifications",
			Order:   6,
		},
		{
			Title:   "Tipo email notificaciones",
			Icon:    "mdi-email",
			Link:    "notificaciones-email-tipo",
			Setting: true,
			PathAPI: "email-notifications",
			Count:   "email_notifications_types",
			Order:   7,
		},
		{
			Title:   "Notificaciones Push automáticas",
			Icon:    "mdi-firebase",
			Link:    "notificaciones-automaticas",
			Setting: true,
			PathAPI: "push-notification",
			Count:   "automatic_notification_pushes",
			Order:   8,
		},
	}

	for _, form := range forms {
		if err := dbConn.FirstOrCreate(&form, model.Form{Title: form.Title}).Error; err != nil {
			log.Printf("Failed to seed form %s: %v", form.Title, err)
		}
	}
}

func seedLevels(dbConn *gorm.DB) {
	levels := []model.Level{
		{
			Model: gorm.Model{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			Level:       "Desarrollo",
			Description: "Rol exclusivo para los desarrolladores",
		},
		{
			Model:       gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Level:       "Administrador",
			Description: "Lo puede hacer todo y más",
		},
		{
			Model:       gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()},
			Level:       "Invitado",
			Description: "Rol para visualizar datos",
		},
	}

	for _, level := range levels {
		if err := dbConn.FirstOrCreate(&level, model.Level{Level: level.Level}).Error; err != nil {
			log.Printf("Failed to seed level %s: %v", level.Level, err)
		}
	}
}

func seedLevelPrivileges(dbConn *gorm.DB) {
	levelPrivileges := []model.LevelPrivileges{
		{Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}, LevelID: 1, FormID: 1, Read: true, Write: true},
		{Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}, LevelID: 1, FormID: 2, Read: true, Write: true},
		{Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}, LevelID: 1, FormID: 3, Read: true, Write: true},
		{Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}, LevelID: 1, FormID: 4, Read: true, Write: true},
		{Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}, LevelID: 1, FormID: 5, Read: true, Write: true},
		{Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}, LevelID: 1, FormID: 6, Read: true, Write: true},
		{Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}, LevelID: 1, FormID: 7, Read: true, Write: true},
		{Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}, LevelID: 1, FormID: 8, Read: true, Write: true},
		{Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}, LevelID: 2, FormID: 1, Read: true, Write: true},
		{Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}, LevelID: 2, FormID: 2, Read: true, Write: true},
		{Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}, LevelID: 3, FormID: 1, Read: true, Write: false},
		{Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}, LevelID: 3, FormID: 2, Read: true, Write: false},
	}

	for _, levelPrivilege := range levelPrivileges {
		if err := dbConn.FirstOrCreate(&levelPrivilege, model.LevelPrivileges{LevelID: levelPrivilege.LevelID, FormID: levelPrivilege.FormID}).Error; err != nil {
			log.Printf("Failed to seed level privilege LevelID: %d, FormID: %d: %v", levelPrivilege.LevelID, levelPrivilege.FormID, err)
		}
	}
}

func seedUsers(dbConn *gorm.DB) {

	pw := sha256.Sum256([]byte("awesomepassword"))
	pwd := fmt.Sprintf("%x", pw)

	users := []model.User{
		{
			Username: "Admin",
			Email:    "admin@drossan.com",
			FullName: "Default user",
			LevelID:  1,
			Password: pwd,
		},
	}

	for _, user := range users {
		if err := dbConn.FirstOrCreate(&user, model.User{Email: user.Email}).Error; err != nil {
			log.Printf("Failed to seed user %s: %v", user.Email, err)
		}
	}
}
