package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbInstance *gorm.DB

func Initialize(databaseURL string) *gorm.DB {
	// Obtener el valor de la variable de entorno GORM_LOG_MODE
	logMode := os.Getenv("LOG_MODE")

	// Configurar el logger de GORM
	var gormConfig *gorm.Config
	if logMode == "true" {
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		}
	} else {
		gormConfig = &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}
	}

	db, err := gorm.Open(mysql.Open(databaseURL), gormConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	dbInstance = db
	return dbInstance
}

func GetConnection() *gorm.DB {
	if dbInstance == nil {
		log.Fatalf("Database connection is not initialized")
	}
	return dbInstance
}
