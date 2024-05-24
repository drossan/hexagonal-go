package db_test

import (
	"fmt"
	"github.com/drossan/core-api/utils"
	"testing"

	"github.com/drossan/core-api/domain/model"
	"github.com/drossan/core-api/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestLevelRepository_CreateOrUpdate(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewLevelRepository(database)

	level := &model.Level{
		Level:       "Test Level",
		Description: "Test Description",
	}

	err := repo.CreateOrUpdate(level)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, level.ID)
}

func TestLevelRepository_GetByID(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewLevelRepository(database)

	level := &model.Level{
		Level:       "Test Level",
		Description: "Test Description",
	}

	err := repo.CreateOrUpdate(level)
	assert.Nil(t, err)

	foundLevel, err := repo.GetByID(level.ID)
	assert.Nil(t, err)
	assert.Equal(t, level.ID, foundLevel.ID)
	assert.Equal(t, level.Level, foundLevel.Level)
	assert.Equal(t, level.Description, foundLevel.Description)
}

func TestLevelRepository_GetAll(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewLevelRepository(database)

	level1 := &model.Level{
		Level:       "Test Level 1",
		Description: "Test Description 1",
	}
	level2 := &model.Level{
		Level:       "Test Level 2",
		Description: "Test Description 2",
	}

	err := repo.CreateOrUpdate(level1)
	assert.Nil(t, err)
	err = repo.CreateOrUpdate(level2)
	assert.Nil(t, err)

	levels, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Len(t, levels, 2)
}

func TestLevelRepository_Paginate(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewLevelRepository(database)

	for i := 1; i <= 25; i++ {
		level := &model.Level{
			Level:       fmt.Sprintf("Test Level %d", i),
			Description: fmt.Sprintf("Test Description %d", i),
		}
		err := repo.CreateOrUpdate(level)
		assert.Nil(t, err)
	}

	levels, total, err := repo.Paginate(1, 10)
	assert.Nil(t, err)
	assert.Len(t, levels, 10)
	assert.Equal(t, 25, total)
}

func TestLevelRepository_Delete(t *testing.T) {
	database := utils.SetupTestDB(t)
	utils.ResetTestDB(database, t)
	repo := db.NewLevelRepository(database)

	level := &model.Level{
		Level:       "Test Level",
		Description: "Test Description",
	}

	err := repo.CreateOrUpdate(level)
	assert.Nil(t, err)

	err = repo.Delete(level)
	assert.Nil(t, err)

	_, err = repo.GetByID(level.ID)
	assert.NotNil(t, err)
}
