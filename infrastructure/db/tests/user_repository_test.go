package db_test

import (
	"fmt"
	"github.com/drossan/core-api/utils"
	"testing"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewUserRepository(database)

	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	err := repo.Create(user)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, user.ID)
}

func TestUserRepository_GetByID(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewUserRepository(database)

	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	err := repo.Create(user)
	assert.Nil(t, err)

	foundUser, err := repo.GetByID(user.ID)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, foundUser.ID)
	assert.Equal(t, user.Username, foundUser.Username)
	assert.Equal(t, user.Email, foundUser.Email)
}

func TestUserRepository_GetByEmail(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewUserRepository(database)

	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	err := repo.Create(user)
	assert.Nil(t, err)

	foundUser, err := repo.GetByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, foundUser.ID)
	assert.Equal(t, user.Username, foundUser.Username)
	assert.Equal(t, user.Email, foundUser.Email)
}

func TestUserRepository_GetAll(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewUserRepository(database)

	user1 := &model.User{
		Username: "testuser1",
		Email:    "test1@example.com",
		Password: "password",
	}
	user2 := &model.User{
		Username: "testuser2",
		Email:    "test2@example.com",
		Password: "password",
	}

	err := repo.Create(user1)
	assert.Nil(t, err)
	err = repo.Create(user2)
	assert.Nil(t, err)

	users, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Len(t, users, 2)
}

func TestUserRepository_Paginate(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewUserRepository(database)

	for i := 1; i <= 25; i++ {
		user := &model.User{
			Username: fmt.Sprintf("testuser%d", i),
			Email:    fmt.Sprintf("test%d@example.com", i),
			Password: "password",
		}
		err := repo.Create(user)
		assert.Nil(t, err)
	}

	users, total, err := repo.Paginate(1, 10)
	assert.Nil(t, err)
	assert.Len(t, users, 10)
	assert.Equal(t, 25, total)
}

func TestUserRepository_Delete(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewUserRepository(database)

	user := &model.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}

	err := repo.Create(user)
	assert.Nil(t, err)

	err = repo.Delete(user)
	assert.Nil(t, err)

	_, err = repo.GetByID(user.ID)
	assert.NotNil(t, err)
}
