package usecase_test

import (
	"gorm.io/gorm/logger"
	"testing"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/infrastructure/db"
	"github.com/drossan/core-api/usecase"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	// Usar SQLite en memoria para pruebas
	testDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrar esquemas
	_ = testDB.AutoMigrate(&model.User{})
	return testDB
}

func resetTestDB(testDB *gorm.DB) {
	testDB.Exec("DROP TABLE IF EXISTS users")
	_ = testDB.AutoMigrate(&model.User{})
}

func TestCreateUser(t *testing.T) {
	testDB := setupTestDB()
	defer resetTestDB(testDB)

	userRepo := db.NewUserRepository(testDB)

	userUseCase := usecase.NewUserUseCase(userRepo)

	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		FullName: "Test User",
		Password: "password",
	}

	err := userUseCase.CreateUser(user)

	assert.Nil(t, err)

	var createdUser model.User
	result := testDB.First(&createdUser, "email = ?", "test@example.com")
	assert.Nil(t, result.Error)
	assert.Equal(t, user.Username, createdUser.Username)
	assert.Equal(t, user.Email, createdUser.Email)
	assert.Equal(t, user.FullName, createdUser.FullName)
}

func TestGetUserByID(t *testing.T) {
	testDB := setupTestDB()
	defer resetTestDB(testDB)

	userRepo := db.NewUserRepository(testDB)
	userUseCase := usecase.NewUserUseCase(userRepo)

	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		FullName: "Test User",
		Password: "password",
	}

	// Crear el usuario antes de intentar obtenerlo
	err := userUseCase.CreateUser(user)
	assert.Nil(t, err)

	fetchedUser, err := userUseCase.GetUserByID(user.ID)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, fetchedUser.ID)
	assert.Equal(t, user.Username, fetchedUser.Username)
	assert.Equal(t, user.Email, fetchedUser.Email)
	assert.Equal(t, user.FullName, fetchedUser.FullName)
}

func TestUpdateUser(t *testing.T) {
	testDB := setupTestDB()
	defer resetTestDB(testDB)

	userRepo := db.NewUserRepository(testDB)
	userUseCase := usecase.NewUserUseCase(userRepo)

	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		FullName: "Test User",
		Password: "password",
	}

	err := userUseCase.CreateUser(user)
	assert.Nil(t, err)

	user.FullName = "Updated Test User"
	err = userUseCase.UpdateUser(user)
	assert.Nil(t, err)

	var updatedUser model.User
	testDB.First(&updatedUser, user.ID)

	assert.Equal(t, "Updated Test User", updatedUser.FullName)
}

func TestDeleteUser(t *testing.T) {
	testDB := setupTestDB()
	defer resetTestDB(testDB)

	userRepo := db.NewUserRepository(testDB)
	userUseCase := usecase.NewUserUseCase(userRepo)

	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		FullName: "Test User",
		Password: "password",
	}

	err := userUseCase.CreateUser(user)
	assert.Nil(t, err)

	err = userUseCase.DeleteUser(user)
	assert.Nil(t, err)

	var deletedUser model.User
	result := testDB.First(&deletedUser, user.ID)

	// Verificar que el usuario no se encuentra en la base de datos
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
}
