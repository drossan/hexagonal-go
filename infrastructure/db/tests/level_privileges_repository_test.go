package db_test

import (
	"github.com/drossan/core-api/utils"
	"testing"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestLevelPrivilegesRepository_CreateOrUpdate(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewLevelPrivilegesRepository(database)

	levelPrivileges := &model.LevelPrivileges{
		Read:  true,
		Write: true,
	}

	err := repo.CreateOrUpdate(levelPrivileges)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, levelPrivileges.ID)
}

func TestLevelPrivilegesRepository_GetAll(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewLevelPrivilegesRepository(database)

	levelPrivileges1 := &model.LevelPrivileges{
		Read:  true,
		Write: true,
	}
	levelPrivileges2 := &model.LevelPrivileges{
		Read:  true,
		Write: false,
	}

	err := repo.CreateOrUpdate(levelPrivileges1)
	assert.Nil(t, err)
	err = repo.CreateOrUpdate(levelPrivileges2)
	assert.Nil(t, err)

	levelPrivileges, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Len(t, levelPrivileges, 2)
}

func TestLevelPrivilegesRepository_Delete(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewLevelPrivilegesRepository(database)

	levelPrivileges := &model.LevelPrivileges{
		Read:  true,
		Write: true,
	}

	err := repo.CreateOrUpdate(levelPrivileges)
	assert.Nil(t, err)

	err = repo.Delete(levelPrivileges)
	assert.Nil(t, err)

	levelPrivilegesList, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Len(t, levelPrivilegesList, 0)
}
